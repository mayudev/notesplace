import { SerializedError } from '@reduxjs/toolkit'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAppDispatch } from '../../../app/hooks'
import Button from '../../../components/Button/Button'
import {
  ButtonContainer,
  Container,
  Input,
  PaneHeading,
  PaneSubheading,
} from '../../../components/Panes/Panes'
import PasswordPrompt from '../../../components/PasswordPrompt/PasswordPrompt'
import { fetchNotebook } from '../../../features/notebook/notebook.slice'

export default function JoinPane() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const [query, setQuery] = useState('')
  const [showPasswordPrompt, setShowPasswordPrompt] = useState(false)

  const load = async () => {
    // TODO input validation, url checking
    try {
      const result = await dispatch(
        fetchNotebook({
          id: query,
          jwt: '', // TODO authentication
        })
      ).unwrap()

      navigate('/nb/' + result.id)
    } catch (e) {
      const err = e as SerializedError

      switch (err.code) {
        case '401':
          setShowPasswordPrompt(true)
          alert('unauthorized')
          break
        case '404':
          alert('not found')
          break
        default:
          alert('unknown error: ' + err.code)
      }
    }
  }

  const passwordEntered = (success: boolean, password: string) => {
    setShowPasswordPrompt(false)
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
      {showPasswordPrompt && <PasswordPrompt onSubmit={passwordEntered} />}
    </div>
  )
}
