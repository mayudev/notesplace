import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { createGlobalStyle, ThemeProvider } from 'styled-components'
import { useAppSelector } from './app/hooks'
import { selectTheme } from './features/global/global.slice'
import { BackgroundColor, ForegroundColor } from './lib/colors'
import Home from './pages/home/Home'
import Notebook from './pages/notebook/Notebook'

const GlobalStyle = createGlobalStyle`
  body {
    background: ${BackgroundColor};
    color: ${ForegroundColor};

    transition: background var(--transition-theme), color var(--transition-theme);
  }

  :root {
    --bg:  ${BackgroundColor};
    --fg: ${ForegroundColor};

    --transition-theme: 0.4s;
  }
`

function App() {
  const theme = useAppSelector(selectTheme)

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
