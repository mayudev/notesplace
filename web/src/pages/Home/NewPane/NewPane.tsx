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
  const [title, setTitle] = useState('')
  const [password, setPassword] = useState('')

  const isEnabled = title.length > 0 && (level === 0 || password.length > 0)

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
          value={title}
          onChange={e => setTitle(e.target.value)}
        />
        <ProtectionLevelChooser level={level} onChoose={n => setLevel(n)} />
        {level !== 0 && (
          <Input
            maxLength={512}
            placeholder="Input password..."
            type="password"
            value={password}
            onChange={e => setPassword(e.target.value)}
          />
        )}
      </Container>
      <ButtonContainer>
        <Button disabled={!isEnabled}>Create</Button>
      </ButtonContainer>
    </div>
  )
}
