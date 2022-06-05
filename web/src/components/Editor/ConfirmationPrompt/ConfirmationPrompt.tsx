import Button from '../../Button/Button'
import { Backdrop, Modal } from '../../Modal'
import {
  Buttons,
  Contents,
  Message,
} from '../../PasswordPrompt/PasswordPrompt.styles'

type Props = {
  onConfirm: () => void
  onCancel: () => void
  message: string
  confirmButton: string
}

export default function ConfirmationPrompt(props: Props) {
  return (
    <Modal>
      <Backdrop />
      <Contents>
        <Message>{props.message}</Message>
        <Buttons>
          <Button onClick={props.onConfirm}>{props.confirmButton}</Button>
          <Button onClick={props.onCancel}>Cancel</Button>
        </Buttons>
      </Contents>
    </Modal>
  )
}
