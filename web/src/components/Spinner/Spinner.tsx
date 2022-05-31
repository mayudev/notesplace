import styled, { keyframes } from 'styled-components'
import { Accent, ForegroundColor } from '../../lib/colors'

const spin = keyframes`
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
`

type Props = {
  primary?: boolean
  width: number
  borderWidth: number
}

export default styled.div<Props>`
  width: ${props => props.width}px;
  height: ${props => props.width}px;

  border: ${props => props.borderWidth}px solid transparent;
  border-right: ${props => props.borderWidth}px solid
    ${props => (props.primary ? Accent : ForegroundColor)};

  border-left: ${props => props.borderWidth}px solid
    ${props => (props.primary ? Accent : ForegroundColor)};

  border-radius: 50%;
  animation: ${spin} 0.8s linear infinite;
`
