import { useLayoutEffect, useState } from 'react'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import { selectNoteById } from '../../features/notebook/notebook.slice'
import { noteUpdate } from '../../features/notebook/notebook.thunks'
import { Backdrop, Modal } from '../Modal'
import { Container, Contents, Textarea, TitleInput } from './Editor.styles'
import Header from './Header/Header'

type Props = {
  noteId: string
  onClose: () => void
}

export default function Editor(props: Props) {
  const dispatch = useAppDispatch()
  const note = useAppSelector(state => selectNoteById(state, props.noteId))!

  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  useLayoutEffect(() => {
    if (!note) return

    if (note.title) setTitle(note.title)
    if (note.content) setContent(note.content)
  }, [note])

  const close = () => {
    props.onClose()
  }

  const save = () => {
    dispatch(
      noteUpdate({
        ...note,
        title,
        content,
      })
    )

    props.onClose()
  }

  return (
    <Modal>
      <Backdrop onClick={close} />
      <Container>
        <Header onSave={save} />
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
    </Modal>
  )
}
