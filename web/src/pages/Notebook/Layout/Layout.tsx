import { useAppSelector } from '../../../app/hooks'
import { selectNotebookData } from '../../../features/notebook/notebook.slice'
import { Nav, Title } from '../../Home/TopBar/TopBar.styles'

export default function Layout() {
  const notebook = useAppSelector(selectNotebookData)

  return (
    <Nav>
      <Title>Notebook {notebook.name}</Title>
    </Nav>
  )
}
