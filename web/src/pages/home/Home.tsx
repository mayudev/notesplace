import { Link } from 'react-router-dom'
import { useAppSelector } from '../../app/hooks'
import { selectTheme } from '../../features/global/global.slice'
import TopBar from './topBar/TopBar'

export default function Home() {
  const theme = useAppSelector(selectTheme)

  return (
    <div>
      <TopBar />
      <h1>Home page! {theme}</h1>
      <Link to="/nb/qwerty">Look at a notebook</Link>
    </div>
  )
}
