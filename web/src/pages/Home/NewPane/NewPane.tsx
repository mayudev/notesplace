import { SerializedError } from '@reduxjs/toolkit'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAppDispatch } from '../../../app/hooks'
import Button from '../../../components/Button/Button'
import {
  ButtonContainer,
  Container,
  Input,
  PaneHeading,
  PaneSubheading,
} from '../../../components/Panes/Panes'
import ProtectionLevelChooser from '../../../components/ProtectionLevelChooser/ProtectionLevelChooser'
import Spinner from '../../../components/Spinner/Spinner'
import { createNotebook } from '../../../features/notebook/notebook.thunks'
import { ProtectionLevel } from '../../../features/notebook/notebook.types'

export default function NewPane() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const [level, setLevel] = useState(ProtectionLevel.None)
  const [name, setName] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)

  const isEnabled = name.length > 0 && (level === 0 || password.length > 0)

  const create = async () => {
    if (!isEnabled) return

    setLoading(true)

    try {
      const result = await dispatch(
        createNotebook({
          name,
          protectionLevel: level,
          password,
        })
      ).unwrap()

      navigate('/nb/' + result.id)
    } catch (e: unknown) {
      const err = e as SerializedError
      console.log(err.message)
      // TODO error handling
    }

    setLoading(false)
  }

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
          value={name}
          onChange={e => setName(e.target.value)}
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
        <Button disabled={!isEnabled} onClick={create}>
          {loading ? <Spinner width={16} borderWidth={2} primary /> : 'Create'}
        </Button>
      </ButtonContainer>
    </div>
  )
}
