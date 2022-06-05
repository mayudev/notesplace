import { faClose, faTrash } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { NavButton } from '../../Button/NavButton'
import { Container, Text } from './Header.styles'

type Props = {
  onSave: () => void
  onRemove: () => void
}

export default function Header(props: Props) {
  return (
    <Container>
      <Text>Edit note</Text>
      <span style={{ flex: 1 }} />
      <NavButton
        withText
        error
        title="Delete"
        aria-label="Delete"
        onClick={() => props.onRemove()}
      >
        <FontAwesomeIcon icon={faTrash} />
        <span>Delete</span>
      </NavButton>
      <NavButton
        primary
        title="Save and exit"
        aria-label="Save and exit"
        onClick={() => props.onSave()}
      >
        <FontAwesomeIcon icon={faClose} />
      </NavButton>
    </Container>
  )
}
