import ThemeSwitchButton from '../../../components/ThemeSwitchButton/ThemeSwitchButton'
import { Fill, Nav, Title } from './TopBar.styles'

export default function TopBar() {
  return (
    <Nav>
      <Fill />
      <Title>notesplace</Title>
      <ThemeSwitchButton />
    </Nav>
  )
}
