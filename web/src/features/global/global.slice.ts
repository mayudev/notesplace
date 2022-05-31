import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit'
import { RootState } from '../../app/store'

type AppTheme = 'light' | 'dark'

const initialState: {
  theme: AppTheme
  token: string
} = {
  theme: 'light',
  token: '',
}

/**
 * Attemps to authenticate to a notebook using a password
 */
export const authenticate = createAsyncThunk(
  'global/authenticate',
  async ({ notebook, password }: { notebook: string; password: string }) => {
    const response = await fetch('/api/auth', {
      headers: {
        Notebook: notebook,
        Password: password,
      },
    })

    if (response.status === 200) {
      return {
        success: true,
        token: await response.text(),
      }
    } else {
      return {
        success: false,
        token: '',
      }
    }
  }
)

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
