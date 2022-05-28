import { faLightbulb, faMoon } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import { applyTheme, selectTheme } from '../../features/global/global.slice'
import { Button, Container } from './ThemeSwitchButton.styles'

function Icon({ theme }: { theme: string }) {
  if (theme === 'light') {
    return <FontAwesomeIcon icon={faLightbulb} />
  } else {
    return <FontAwesomeIcon icon={faMoon} />
  }
}

export default function ThemeSwitchButton() {
  const theme = useAppSelector(selectTheme)

  const dispatch = useAppDispatch()

  const themeSwitch = () => {
    const newTheme = theme === 'light' ? 'dark' : 'light'

    // Update the store
    dispatch(applyTheme(newTheme))

    // Save in localStorage for later use
    localStorage.setItem('theme', newTheme)
  }

  return (
    <Container>
      <Button
        onClick={themeSwitch}
        title="Switch between dark and light mode"
        aria-label="Switch between dark and light mode"
      >
        <Icon theme={theme} />
      </Button>
    </Container>
  )
}
