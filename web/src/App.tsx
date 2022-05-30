import { useEffect } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { createGlobalStyle, ThemeProvider } from 'styled-components'
import { useAppDispatch, useAppSelector } from './app/hooks'
import { applyTheme, selectTheme } from './features/global/global.slice'
import { BackgroundColor, Error, ForegroundColor } from './lib/colors'
import Home from './pages/Home/Home'
import Notebook from './pages/Notebook/Notebook'

const GlobalStyle = createGlobalStyle`
  body {
    background: ${BackgroundColor};
    color: ${ForegroundColor};

    transition: background var(--transition-theme), color var(--transition-theme);
  }

  :root {
    --bg:  ${BackgroundColor};
    --fg: ${ForegroundColor};

    --color-error: ${Error};

    --transition-theme: 0.4s;
  }
`

function App() {
  const theme = useAppSelector(selectTheme)
  const dispatch = useAppDispatch()

  useEffect(() => {
    // Apply current theme
    const currentTheme = localStorage.getItem('theme')

    // Check if currentTheme is valid
    if (currentTheme === 'light' || currentTheme === 'dark') {
      dispatch(applyTheme(currentTheme))
    }
  }, [])

  return (
    <ThemeProvider theme={{ mode: theme }}>
      <GlobalStyle />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/nb/:id" element={<Notebook />} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  )
}

export default App
