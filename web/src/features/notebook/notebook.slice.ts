import {
  createAsyncThunk,
  createEntityAdapter,
  createSlice,
  SerializedError,
} from '@reduxjs/toolkit'
import { RootState } from '../../app/store'
import { authenticate } from '../global/global.slice'
import { Note, Notebook, NotebookCreateResponse } from './notebook.types'

export enum ProtectionLevel {
  None,
  ReadOnly,
  Protected,
}

export type Status = 'idle' | 'pending' | 'succeeded' | 'failed'

interface NotebookState extends Omit<Notebook, 'notes'> {
  status: Status
  error: string | undefined
}

const notebookAdapter = createEntityAdapter<Note>()

const initialState = notebookAdapter.getInitialState<NotebookState>({
  id: '',
  name: '',
  protectionLevel: ProtectionLevel.None,
  createdAt: null,
  updatedAt: null,
  status: 'idle',
  error: undefined,
})

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
    // TODO Authorization
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

const notebookSlice = createSlice({
  name: 'notebook',
  initialState,
  reducers: {},
  extraReducers(builder) {
    builder
      .addCase(createNotebook.rejected, (state, action) => {
        state.status = 'failed'
        state.error = action.error.message
      })
      // notebook/fetchNotebook
      .addCase(fetchNotebook.pending, (state, action) => {
        state.status = 'pending'
      })
      .addCase(fetchNotebook.fulfilled, (state, action) => {
        state.status = 'succeeded'
        state.id = action.payload.id
        state.name = action.payload.name!
        state.protectionLevel = action.payload.protectionLevel
        state.createdAt = action.payload.createdAt
        state.updatedAt = action.payload.updatedAt

        if (action.payload.notes) {
          notebookAdapter.upsertMany(state, action.payload.notes)
        }
      })
      .addCase(fetchNotebook.rejected, (state, action) => {
        state.status = 'failed'
        state.error = action.error.message
      })
  },
})

export const selectNotebookId = (state: RootState): string => state.notebook.id

export const selectNotebookData = (state: RootState): Notebook => {
  const notebook = state.notebook
  const { ids, entities, status, error, ...rest } = notebook
  return rest as Notebook
}
export default notebookSlice.reducer
