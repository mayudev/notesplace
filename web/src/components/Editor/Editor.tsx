import { useLayoutEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import { selectNoteById } from '../../features/notebook/notebook.slice'
import { noteDelete, noteUpdate } from '../../features/notebook/notebook.thunks'
import { Backdrop, Modal } from '../Modal'
import ConfirmationPrompt from './ConfirmationPrompt/ConfirmationPrompt'
import { Container, Contents, Textarea, TitleInput } from './Editor.styles'
import Header from './Header/Header'

type Params = {
  note: string
}

export default function Editor() {
  const params = useParams<Params>()
  const navigate = useNavigate()

  const dispatch = useAppDispatch()
  const note = useAppSelector(state => selectNoteById(state, params.note!))!

  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  const [deleteRequest, setDeleteRequest] = useState(false)
  const [exiting, setExiting] = useState(false)

  useLayoutEffect(() => {
    if (!note) return

    if (note.title) setTitle(note.title)
    if (note.content) setContent(note.content)
  }, [note])

  const exit = () => {
    setExiting(true)
    setTimeout(() => {
      navigate(-1)
    }, 200)
  }

  const close = () => {
    if (note.content !== content || note.title !== title) save()
    else exit()
  }

  const save = () => {
    dispatch(
      noteUpdate({
        ...note,
        title,
        content,
      })
    )

    exit()
  }

  const remove = () => {
    dispatch(noteDelete(note))
    exit()
  }

  return (
    <Modal className={`${exiting && 'editor-exit-active'}`}>
      <Backdrop onClick={close} />
      <Container>
        <Header onSave={close} onRemove={() => setDeleteRequest(true)} />
        <Contents>
          <TitleInput
            placeholder="Title"
            value={title}
            onChange={e => setTitle(e.target.value)}
          />
          <Textarea
            placeholder="Note"
            value={content}
            onChange={e => setContent(e.target.value)}
          ></Textarea>
        </Contents>
      </Container>
      {deleteRequest && (
        <ConfirmationPrompt
          message="Are you sure you want to delete this note?"
          confirmButton="Delete"
          onCancel={() => setDeleteRequest(false)}
          onConfirm={() => remove()}
        />
      )}
    </Modal>
  )
}
