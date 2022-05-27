import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { RootState } from '../../app/store'

type AppTheme = 'light' | 'dark'

const initialState: {
  theme: AppTheme
} = {
  theme: 'light',
}

const globalSlice = createSlice({
  name: 'global',
  initialState,
  reducers: {
    applyTheme(state, action: PayloadAction<AppTheme>) {
      state.theme = action.payload
    },
  },
})

export const { applyTheme } = globalSlice.actions

export const selectTheme = (state: RootState) => state.global.theme

export default globalSlice.reducer
