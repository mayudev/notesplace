import { Transition } from 'react-transition-group'
import { Blink } from './Panes'

export default function PaneError({
  visible,
  message,
}: {
  visible: boolean
  message: string
}) {
  return (
    <Transition in={visible} timeout={400} unmountOnExit>
      {state => <Blink state={state}>{message}</Blink>}
    </Transition>
  )
}
