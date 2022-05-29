import {
  createAsyncThunk,
  createEntityAdapter,
  createSlice,
} from '@reduxjs/toolkit'
import { RootState } from '../../app/store'
import { Note, Notebook, NotebookCreateResponse } from './notebook.types'

export enum ProtectionLevel {
  None,
  ReadOnly,
  Protected,
}

interface NotebookState extends Omit<Notebook, 'notes'> {
  status: string
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
    { dispatch }
  ) => {
    const response = await fetch('/api/notebook', {
      method: 'POST',
      body: JSON.stringify({
        name,
        protection_level: protectionLevel,
        password,
      }),
    })

    if (!response.ok) {
      throw response.status
    }

    const data = (await response.json()) as NotebookCreateResponse
    dispatch(fetchNotebook({ id: data.id!, jwt: '' }))
  }
)

export const fetchNotebook = createAsyncThunk(
  'notebook/fetchNotebook',
  async ({ id, jwt }: { id: string; jwt: string | null }) => {
    // TODO Authorization
    const response = await fetch('/api/notebook/' + id)

    if (!response.ok) {
      throw response.status
    }

    const data = await response.json()

    data.protectionLevel = data.protection_level
    data.createdAt = data.created_at
    data.updatedAt = data.updated_at

    return data as Notebook
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
        state.status = 'loading'
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
