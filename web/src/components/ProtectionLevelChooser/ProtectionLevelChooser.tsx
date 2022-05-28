import { ProtectionLevel } from '../../features/notebook/notebook.slice'
import { Form, Radio } from './ProtectionLevelChooser.styles'

export default function ProtectionLevelChooser({
  level,
  onChoose,
}: {
  level: ProtectionLevel
  onChoose: (level: ProtectionLevel) => void
}) {
  const labels = [
    'No protection',
    'Require password to edit notes',
    'Require password for access',
  ]

  return (
    <Form>
      {labels.map((label, i) => (
        <Radio key={i}>
          <input
            type="radio"
            value={i}
            checked={level === i}
            onChange={() => onChoose(i)}
          />
          <span>{label}</span>
        </Radio>
      ))}
    </Form>
  )
}
