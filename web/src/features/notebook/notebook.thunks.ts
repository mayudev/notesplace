import { createAsyncThunk, SerializedError } from '@reduxjs/toolkit'
import { RootState } from '../../app/store'
import { authenticate } from '../global/global.slice'
import {
  Note,
  Notebook,
  NotebookCreateResponse,
  ProtectionLevel,
} from './notebook.types'
/**
 * Creates a notebook, logins to it if necessary and returns its details.
 */
export const createNotebook = createAsyncThunk(
  'notebook/createNotebook',
  async (
    {
      name,
      protectionLevel,
      password,
    }: {
      name: string
      protectionLevel: ProtectionLevel
      password: string | null
    },
    { dispatch, rejectWithValue }
  ) => {
    // Create the notebook
    const response = await fetch('/api/notebook', {
      method: 'POST',
      body: JSON.stringify({
        name,
        protection_level: protectionLevel,
        password,
      }),
    })

    const data = (await response.json()) as NotebookCreateResponse
    // Check for errors in notebook creation step (rare)
    if (!response.ok) {
      const error: SerializedError = {
        code: response.status.toString(),
        message: data.message,
      }

      throw error
    }

    // Attempt to authenticate
    let token = ''

    if (password && protectionLevel > 0) {
      const auth = await dispatch(
        authenticate({
          notebook: data.id!,
          password: password,
        })
      ).unwrap()

      if (!auth.success) {
        // TODO do it better
        const error: SerializedError = {
          code: '401',
          message: 'not authorized',
        }

        throw error
      }

      token = auth.token
    }

    // Fetch details about a notebook and return them
    await dispatch(fetchNotebook({ id: data.id!, jwt: token })).unwrap()
    return data
  }
)

/**
 * Returns details about a notebook.
 */
export const fetchNotebook = createAsyncThunk(
  'notebook/fetchNotebook',
  async (
    { id, jwt }: { id: string; jwt: string | null },
    { rejectWithValue }
  ) => {
    try {
      const response = await fetch('/api/notebook/' + id, {
        headers: {
          Authorization: jwt ? 'Bearer ' + jwt : '',
        },
      })

      const data = await response.json()

      if (!response.ok) {
        const error: SerializedError = {
          code: response.status.toString(),
          message: data.message,
        }

        throw error
      }

      data.protectionLevel = data.protection_level
      data.createdAt = data.created_at
      data.updatedAt = data.updated_at

      return data as Notebook
    } catch (e: unknown) {
      const err = e as SerializedError
      return rejectWithValue(err)
    }
  }
)

/**
 * Creates a note
 */
export const noteCreate = createAsyncThunk(
  'notebook/noteCreate',
  async (_, { getState, rejectWithValue }) => {
    try {
      const state = getState() as RootState
      const jwt = state.global.token

      const response = await fetch('/api/note/', {
        method: 'PUT',
        headers: {
          Authorization: jwt ? 'Bearer ' + jwt : '',
        },
        body: JSON.stringify({
          notebook_id: state.notebook.id,
        }),
      })

      const data = await response.json()

      if (!response.ok) {
        const error: SerializedError = {
          code: response.status.toString(),
          message: data.message,
        }

        throw error
      }

      data.notebookId = data.notebook_id
      return data as Note
    } catch (e) {
      const err = e as SerializedError
      return rejectWithValue(err)
    }
  }
)
