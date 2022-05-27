import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit'
import globalSlice from '../features/global/global.slice'

export const store = configureStore({
  reducer: {
    global: globalSlice,
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
