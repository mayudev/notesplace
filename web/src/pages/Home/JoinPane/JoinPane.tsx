import { useState } from 'react'
import { useAppDispatch } from '../../../app/hooks'
import Button from '../../../components/Button/Button'
import {
  ButtonContainer,
  Container,
  Input,
  PaneHeading,
  PaneSubheading,
} from '../../../components/Panes/Panes'
import { fetchNotebook } from '../../../features/notebook/notebook.slice'

export default function JoinPane() {
  const dispatch = useAppDispatch()

  const [query, setQuery] = useState('')

  const load = () => {
    // TODO input validation, url checking
    dispatch(
      fetchNotebook({
        id: query,
        jwt: '',
      })
    )
  }

  return (
    <div>
      <PaneHeading>Open notebook</PaneHeading>
      <PaneSubheading>
        Just enter the URL or the ID you received!
      </PaneSubheading>
      <Container>
        <Input
          value={query}
          onChange={e => setQuery(e.target.value)}
          placeholder="Existing notebook URL or ID"
          type="text"
        />
      </Container>
      <ButtonContainer>
        <Button onClick={load}>Enter</Button>
      </ButtonContainer>
    </div>
  )
}
