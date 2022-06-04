import { EntityId } from '@reduxjs/toolkit'
import { useAppSelector } from '../../app/hooks'
import { selectNoteById } from '../../features/notebook/notebook.slice'
import { Container, Content, Title } from './Note.styles'

type Params = {
  // Note ID
  noteId: EntityId
}

export default function Note({ noteId }: Params) {
  const note = useAppSelector(state => selectNoteById(state, noteId))!

  return (
    <Container>
      <Title>{note.title ? note.title : 'No title'}</Title>
      <Content>{note.content ? note.content : '<empty>'}</Content>
    </Container>
  )
}
