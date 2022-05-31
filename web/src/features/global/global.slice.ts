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
  async (
    { notebook, password }: { notebook: string; password: string },
    { rejectWithValue }
  ) => {
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
      return rejectWithValue({
        success: false,
        token: '',
      })
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
    clearToken(state) {
      state.token = ''
    },
  },
  extraReducers(builder) {
    builder.addCase(authenticate.fulfilled, (state, action) => {
      state.token = action.payload.token
    })
  },
})

export const { applyTheme, clearToken } = globalSlice.actions

export const selectTheme = (state: RootState) => state.global.theme
export const selectToken = (state: RootState) => state.global.token

export default globalSlice.reducer
