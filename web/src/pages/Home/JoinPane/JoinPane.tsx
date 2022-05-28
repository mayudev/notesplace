import Button from '../../../components/Button/Button'
import {
  ButtonContainer,
  Input,
  PaneHeading,
  PaneSubheading,
} from '../../../components/Panes/Panes'

export default function JoinPane() {
  return (
    <div>
      <PaneHeading>Open notebook</PaneHeading>
      <PaneSubheading>
        Just enter the URL or the ID you received!
      </PaneSubheading>
      <Input placeholder="Existing notebook URL or ID" type="text" />
      <ButtonContainer>
        <Button>Enter</Button>
      </ButtonContainer>
    </div>
  )
}
