import { Radio } from './ProtectionLevelChooser.styles'

export enum ProtectionLevel {
  None,
  ReadOnly,
  Protected,
}

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
    <form>
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
    </form>
  )
}
