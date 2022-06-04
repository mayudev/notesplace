import { useEffect, useRef, useState } from 'react'
import { CSSTransition } from 'react-transition-group'
import { useAppDispatch } from '../../app/hooks'
import { authenticate } from '../../features/global/global.slice'
import Button from '../Button/Button'
import { Backdrop, Modal } from '../Modal'
import { Input } from '../Panes/Panes'
import {
  Buttons,
  Contents,
  IncorrectMessage,
  Message,
} from './PasswordPrompt.styles'

type Props = {
  notebook: string
  onSubmit: (success: boolean, input: string) => void
}

export default function PasswordPrompt(props: Props) {
  const dispatch = useAppDispatch()
  const inputRef = useRef<HTMLInputElement>(null)

  const [password, setPassword] = useState('')
  const [incorrect, setIncorrect] = useState(false)

  useEffect(() => {
    inputRef.current?.focus()
  }, [])

  const close = () => {
    props.onSubmit(false, '')
  }

  const submit = async () => {
    try {
      const auth = await dispatch(
        authenticate({
          notebook: props.notebook,
          password: password,
        })
      ).unwrap()

      if (auth.success) {
        props.onSubmit(true, auth.token)
      }
    } catch (e) {
      incorrectEntered()
    }
  }

  const incorrectEntered = () => {
    setIncorrect(true)
    setPassword('')

    setTimeout(() => {
      setIncorrect(false)
    }, 1500)
  }

  const captureKey = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      // Submit on return key pressed
      submit()
    } else if (event.key === 'Escape') {
      // Cancel on escape pressed
      props.onSubmit(false, '')
    }
  }

  return (
    <Modal>
      <Backdrop onClick={close} />
      <Contents>
        <Message>Enter password to continue.</Message>
        <IncorrectMessage>
          <CSSTransition
            in={incorrect}
            timeout={400}
            classNames="incorrect"
            unmountOnExit
          >
            <span>Incorrect password</span>
          </CSSTransition>
        </IncorrectMessage>
        <Input
          type="password"
          ref={inputRef}
          value={password}
          onChange={e => setPassword(e.target.value)}
          onKeyDown={captureKey}
        />
        <Buttons>
          <Button onClick={() => submit()}>Continue</Button>
          <Button onClick={() => props.onSubmit(false, '')}>Cancel</Button>
        </Buttons>
      </Contents>
    </Modal>
  )
}
