import {
  Center,
  Container,
  Heading,
  Hero,
  Pane,
  Panes,
  Subheading,
} from './Home.styles'
import JoinPane from './JoinPane/JoinPane'
import NewPane from './NewPane/NewPane'
import TopBar from './TopBar/TopBar'

export default function Home() {
  return (
    <Container>
      <TopBar />
      <Hero>
        <Heading>There are notes, and this is their place.</Heading>
        <Subheading>
          Share notes with just a link. And a password, if you want.
        </Subheading>
      </Hero>
      <Center>Start now</Center>
      <Panes>
        <Pane>
          <JoinPane />
        </Pane>
        <Pane>
          <NewPane />
        </Pane>
      </Panes>
    </Container>
  )
}
