import styled from 'styled-components'
import { Accent, BackgroundSecondary } from '../../lib/colors'

type Props = {
  primary?: boolean
}

const Button = styled.button<Props>`
  background: ${props => (props.primary ? 'palevioletred' : 'transparent')};
  font: inherit;
  color: ${Accent};

  font-size: 1em;

  outline: none;
  border: none;
  border-radius: 3px;

  padding: 8px 16px;

  &:active,
  &:focus {
    outline: none;
  }

  transition: background 0.2s ease-in;
  cursor: pointer;

  &:hover {
    background: ${BackgroundSecondary};
  }
`

export default Button
