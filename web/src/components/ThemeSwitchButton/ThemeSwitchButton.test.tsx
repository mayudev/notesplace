import { configureStore } from '@reduxjs/toolkit'
import { fireEvent, render, screen } from '@testing-library/react'
import { Provider } from 'react-redux'
import globalSlice from '../../features/global/global.slice'
import ThemeSwitchButton from './ThemeSwitchButton'

test('it should switch the theme', () => {
  const store = configureStore({
    reducer: {
      global: globalSlice,
    },
  })
  render(
    <Provider store={store}>
      <ThemeSwitchButton />
    </Provider>
  )

  expect(store.getState().global.theme).toEqual('light')

  fireEvent.click(screen.getByRole('button'))

  expect(store.getState().global.theme).toEqual('dark')
})
