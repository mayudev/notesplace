import { EntityId } from '@reduxjs/toolkit'
import { useAppSelector } from '../../app/hooks'
import { selectNoteById } from '../../features/notebook/notebook.slice'
import { Note as INote } from '../../features/notebook/notebook.types'
import { Container, Content, Title } from './Note.styles'

type Props = {
  // Note ID
  noteId: EntityId

  onClick: () => void
}

export default function Note(props: Props) {
  const note = useAppSelector(state => selectNoteById(state, props.noteId))!

  return (
    <Container onClick={() => props.onClick()}>
      <Title>{note.title ? note.title : 'No title'}</Title>
      <Content>{note.content ? note.content : '<empty>'}</Content>
    </Container>
  )
}
