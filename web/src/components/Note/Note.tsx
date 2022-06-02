import { EntityId } from '@reduxjs/toolkit'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import { selectNoteById } from '../../features/notebook/notebook.slice'
import { noteDelete } from '../../features/notebook/notebook.thunks'
import Button from '../Button/Button'

type Params = {
  // Note ID
  noteId: EntityId
}

export default function Note({ noteId }: Params) {
  const dispatch = useAppDispatch()
  const note = useAppSelector(state => selectNoteById(state, noteId))!

  const deleteNote = () => {
    dispatch(noteDelete(note))
  }

  return (
    <div>
      {note.id} Order: {note.order}{' '}
      <Button onClick={() => deleteNote()}>delete</Button>
    </div>
  )
}
