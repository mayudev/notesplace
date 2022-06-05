import { faPlus } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { EntityId } from '@reduxjs/toolkit'
import { useState } from 'react'
import { Outlet, useNavigate } from 'react-router-dom'
import { CSSTransition } from 'react-transition-group'
import { useAppDispatch, useAppSelector } from '../../../app/hooks'
import Editor from '../../../components/Editor/Editor'
import Note from '../../../components/Note/Note'
import { selectNoteIds } from '../../../features/notebook/notebook.slice'
import { noteCreate } from '../../../features/notebook/notebook.thunks'
import { Container, CreateNote } from './Notes.styles'

export default function Notes() {
  const dispatch = useAppDispatch()
  const noteIds = useAppSelector(selectNoteIds)
  const navigate = useNavigate()

  const [editingId, setEditingId] = useState('')

  const create = () => {
    dispatch(noteCreate())
  }

  const update = (id: EntityId) => {
    navigate('edit/' + id)
  }

  return (
    <Container>
      <CreateNote
        title="Create a new note"
        aria-label="Create a new note"
        onClick={() => create()}
      >
        <FontAwesomeIcon size="3x" icon={faPlus} />
      </CreateNote>
      {noteIds.map(noteId => (
        <Note key={noteId} noteId={noteId} onClick={() => update(noteId)} />
      ))}

      <Outlet />
    </Container>
  )
}
