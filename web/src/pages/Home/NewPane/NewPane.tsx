import { useState } from 'react'
import Button from '../../../components/Button/Button'
import {
  ButtonContainer,
  Container,
  Input,
  PaneHeading,
  PaneSubheading,
} from '../../../components/Panes/Panes'
import ProtectionLevelChooser, {
  ProtectionLevel,
} from '../../../components/ProtectionLevelChooser/ProtectionLevelChooser'

export default function NewPane() {
  const [level, setLevel] = useState(ProtectionLevel.None)

  return (
    <div>
      <PaneHeading>Create a new notebook</PaneHeading>
      <PaneSubheading>
        Those settings <strong>can not</strong> be changed later.
      </PaneSubheading>
      <Container>
        <Input
          maxLength={256}
          placeholder="Name your notebook..."
          type="text"
        />
        <ProtectionLevelChooser level={level} onChoose={n => setLevel(n)} />
        {level !== 0 && (
          <Input
            maxLength={512}
            placeholder="Input password..."
            type="password"
          />
        )}
      </Container>
      <ButtonContainer>
        <Button>Create</Button>
      </ButtonContainer>
    </div>
  )
}
