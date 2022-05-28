import { Link } from 'react-router-dom'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import { applyTheme, selectTheme } from '../../features/global/global.slice'

export default function Home() {
  const theme = useAppSelector(selectTheme)

  const dispatch = useAppDispatch()

  const themeSwitch = () => {
    if (theme === 'light') {
      dispatch(applyTheme('dark'))
    } else {
      dispatch(applyTheme('light'))
    }
  }

  return (
    <div>
      <h1>notesplace</h1>
      <h1>Home page! {theme}</h1>
      <button onClick={themeSwitch}>
        turn the lights {theme === 'light' ? 'off' : 'on'}
      </button>
      <Link to="/nb/qwerty">Look at a notebook</Link>
    </div>
  )
}
