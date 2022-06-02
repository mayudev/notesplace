import { createEntityAdapter, createSlice } from '@reduxjs/toolkit'
import { RootState } from '../../app/store'
import {
  createNotebook,
  fetchNotebook,
  noteCreate,
  noteDelete,
} from './notebook.thunks'
import {
  Note,
  Notebook,
  NotebookState,
  ProtectionLevel,
} from './notebook.types'

const notebookAdapter = createEntityAdapter<Note>({
  sortComparer: (a, b) => b.order - a.order, // Sort by reversed order (highest = first)
})

const initialState = notebookAdapter.getInitialState<NotebookState>({
  id: '',
  name: '',
  protectionLevel: ProtectionLevel.None,
  createdAt: null,
  updatedAt: null,
  status: 'idle',
  error: undefined,
})

const notebookSlice = createSlice({
  name: 'notebook',
  initialState,
  reducers: {
    clearNotebook: state => {
      state.id = ''
    },
  },
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

        notebookAdapter.removeAll(state)

        if (action.payload.notes) {
          notebookAdapter.upsertMany(state, action.payload.notes)
        }
      })
      .addCase(fetchNotebook.rejected, (state, action) => {
        state.status = 'failed'
        state.error = action.error.message
      })

      // noteCreate
      .addCase(noteCreate.fulfilled, (state, action) => {
        notebookAdapter.upsertOne(state, action.payload)
      })

      // noteDelete
      .addCase(noteDelete.fulfilled, (state, action) => {
        const order = action.payload.note.order

        notebookAdapter.removeOne(state, action.payload.note.id)

        notebookAdapter
          .getSelectors()
          .selectAll(state)
          .filter(note => note.order > order)
          .forEach(note => {
            notebookAdapter.updateOne(state, {
              id: note.id,
              changes: {
                order: note.order - 1,
              },
            })
          })
      })
  },
})

export const selectNotebookId = (state: RootState): string => state.notebook.id

export const selectNotebookData = (state: RootState): Notebook => {
  const notebook = state.notebook
  const { ids, entities, status, error, ...rest } = notebook
  return rest as Notebook
}

export const { selectById: selectNoteById, selectIds: selectNoteIds } =
  notebookAdapter.getSelectors((state: RootState) => state.notebook)

export const { clearNotebook } = notebookSlice.actions

export default notebookSlice.reducer
