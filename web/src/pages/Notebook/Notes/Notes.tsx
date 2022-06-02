import { useAppDispatch, useAppSelector } from '../../../app/hooks'
import Button from '../../../components/Button/Button'
import Note from '../../../components/Note/Note'
import {
  noteCreate,
  selectNoteIds,
} from '../../../features/notebook/notebook.slice'

export default function Notes() {
  const dispatch = useAppDispatch()
  const noteIds = useAppSelector(selectNoteIds)

  const create = () => {
    dispatch(noteCreate())
  }

  return (
    <div>
      <h1>notes</h1>
      <Button onClick={() => create()}>Create new note</Button>
      {noteIds.map(noteId => (
        <Note key={noteId} noteId={noteId} />
      ))}
    </div>
  )
}
