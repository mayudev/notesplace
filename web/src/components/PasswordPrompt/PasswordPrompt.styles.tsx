import styled from 'styled-components'
import { Accent, BackgroundColor } from '../../lib/colors'
import { Appear } from '../Modal'

export const Contents = styled.div`
  background: ${BackgroundColor};

  border-radius: 4px;
  border: 1px solid ${Accent};

  min-width: 400px;
  animation: ${Appear} 0.3s ease-in-out forwards;
`

export const Message = styled.div`
  padding: 1rem;
  font-size: 1.2rem;
`

export const IncorrectMessage = styled.div`
  span {
    padding: 0 1rem;
    padding-bottom: 1rem;
    display: block;
  }

  .incorrect-enter {
    opacity: 0;
    margin-top: -30px;
  }

  .incorrect-enter-active {
    opacity: 1;
    margin-top: 0;
    transition: 200ms;
  }

  .incorrect-exit-active {
    opacity: 0;
    margin-top: -35px;
    transition: 200ms;
  }
`

export const Buttons = styled.div`
  display: flex;
  flex-direction: row-reverse;

  padding: 1rem;
`
