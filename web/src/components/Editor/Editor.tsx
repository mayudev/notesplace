import { Backdrop, Modal } from '../Modal'
import { Container, Contents, Textarea, TitleInput } from './Editor.styles'
import Header from './Header/Header'

export default function Editor() {
  return (
    <Modal>
      <Backdrop />
      <Container>
        <Header />
        <Contents>
          <TitleInput placeholder="Title" />
          <Textarea placeholder="Note"></Textarea>
        </Contents>
      </Container>
    </Modal>
  )
}
