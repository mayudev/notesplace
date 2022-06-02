import { faClose } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useAppSelector } from '../../../app/hooks'
import Button from '../../../components/Button/Button'
import { NavButton } from '../../../components/Button/NavButton'
import ThemeSwitchButton from '../../../components/ThemeSwitchButton/ThemeSwitchButton'
import { selectNotebookData } from '../../../features/notebook/notebook.slice'
import { Nav, Title } from './Layout.styles'

type Props = {
  onClose: () => void
  onUnlock: () => void
}

export default function Layout(props: Props) {
  const notebook = useAppSelector(selectNotebookData)

  return (
    <Nav>
      <div>
        <NavButton
          title="Close notebook"
          aria-label="Close notebook"
          onClick={props.onClose}
        >
          <FontAwesomeIcon icon={faClose} />
        </NavButton>
      </div>
      <Title>Notebook {notebook.name}</Title>
      <span style={{ flex: 1 }} />
      {!notebook.unlocked && <Button onClick={props.onUnlock}>Unlock</Button>}
      <ThemeSwitchButton />
    </Nav>
  )
}
