import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit'
import globalSlice from '../features/global/global.slice'
import notebookSlice from '../features/notebook/notebook.slice'

export const store = configureStore({
  reducer: {
    global: globalSlice,
    notebook: notebookSlice,
  },
})

export type AppDispatch = typeof store.dispatch
export type RootState = ReturnType<typeof store.getState>
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>
