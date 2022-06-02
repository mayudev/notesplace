import { EntityId } from '@reduxjs/toolkit'
import { useAppSelector } from '../../app/hooks'
import { selectNoteById } from '../../features/notebook/notebook.slice'

type Params = {
  // Note ID
  noteId: EntityId
}

export default function Note({ noteId }: Params) {
  const note = useAppSelector(state => selectNoteById(state, noteId))!

  return (
    <div>
      {note.id} Order: {note.order}
    </div>
  )
}
