import { useState } from 'react'

type Props = {
  onSubmit: (success: boolean, input: string) => void
}

export default function PasswordPrompt(props: Props) {
  const [password, setPassword] = useState('')

  return (
    <div>
      <span>a password prompt</span>
      <input
        type="password"
        value={password}
        onChange={e => setPassword(e.target.value)}
      />
      <button onClick={() => props.onSubmit(false, '')}>Cancel</button>
      <button onClick={() => props.onSubmit(true, password)}>Continue</button>
    </div>
  )
}
