import Button from '../../Button/Button'
import { Container } from './Header.styles'

export default function Header() {
  return (
    <Container>
      <div>Edit note</div>
      <span style={{ flex: 1 }} />
      <Button>Save</Button>
    </Container>
  )
}
