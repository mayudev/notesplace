import Button from '../../Button/Button'
import { Container } from './Header.styles'

type Props = {
  onSave: () => void
}

export default function Header(props: Props) {
  return (
    <Container>
      <div>Edit note</div>
      <span style={{ flex: 1 }} />
      <Button onClick={props.onSave}>Save</Button>
    </Container>
  )
}
