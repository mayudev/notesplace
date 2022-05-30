import { SerializedError } from '@reduxjs/toolkit'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAppDispatch } from '../../../app/hooks'
import Button from '../../../components/Button/Button'
import PaneError from '../../../components/Panes/PaneError'
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
  const navigate = useNavigate()

  const [query, setQuery] = useState('')
  const [errorMessage, setErrorMessage] = useState('')
  const [errorVisible, setErrorVisible] = useState(false)

  const [showPasswordPrompt, setShowPasswordPrompt] = useState(false)

  // Show an error and hide it after a few seconds
  const showError = (message: string) => {
    setErrorMessage(message)
    setErrorVisible(true)

    setTimeout(() => {
      setErrorVisible(false)
    }, 2000)

    // This is just as stupid as I am!
    setTimeout(() => {
      setErrorMessage('')
    }, 2400)
  }

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
          showError('Notebook not found')
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
      <PaneError visible={errorVisible} message={errorMessage} />
    </div>
  )
}
