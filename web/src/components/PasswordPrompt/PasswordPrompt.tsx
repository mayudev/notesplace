import { useState } from 'react'
import Button from '../Button/Button'
import { ButtonContainer, Input } from '../Panes/Panes'
import {
  Backdrop,
  Buttons,
  Container,
  Contents,
  Message,
} from './PasswordPrompt.styles'

type Props = {
  onSubmit: (success: boolean, input: string) => void
}

export default function PasswordPrompt(props: Props) {
  const [password, setPassword] = useState('')

  const close = () => {
    props.onSubmit(false, '')
  }

  const captureKey = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      // Submit on return key presset
      props.onSubmit(true, password)
    } else if (event.key === 'Escape') {
      // Cancel on escape pressed
      props.onSubmit(false, '')
    }
  }

  return (
    <Container>
      <Backdrop onClick={close} />
      <Contents>
        <Message>Enter password to continue.</Message>
        <Input
          type="password"
          value={password}
          onChange={e => setPassword(e.target.value)}
          onKeyDown={captureKey}
        />
        <Buttons>
          <Button onClick={() => props.onSubmit(true, password)}>
            Continue
          </Button>
          <Button onClick={() => props.onSubmit(false, '')}>Cancel</Button>
        </Buttons>
      </Contents>
    </Container>
  )
}
