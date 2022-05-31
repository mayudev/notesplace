import { faExclamation } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { PropsWithChildren } from 'react'
import { Link } from 'react-router-dom'
import Button from '../../../components/Button/Button'
import { Center, Message } from './NotebookError.styles'

export default function NotebookError({ children }: PropsWithChildren<{}>) {
  return (
    <Center>
      <FontAwesomeIcon icon={faExclamation} fixedWidth size="6x" />
      <Message>{children}</Message>
      <Link to="/">
        <Button as="span">Return to home page</Button>
      </Link>
    </Center>
  )
}
