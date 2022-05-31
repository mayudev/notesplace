import styled, { keyframes } from 'styled-components'
import { Accent, BackgroundColor } from '../../lib/colors'

export const Container = styled.div`
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;

  display: flex;
  justify-content: center;
  align-items: center;
`

const Appear = keyframes`
  from {
    opacity: 0;
    transform: scale(0.7);
  }

  30% {
    transform: scale(1.1);
  }

  to {
    opacity: 1;
  }
  `

export const Backdrop = styled.div`
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;

  z-index: -1;

  background: #00000070;
  animation: ${Appear} 0.2s linear forwards;
`

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

export const Buttons = styled.div`
  display: flex;
  flex-direction: row-reverse;

  padding: 1rem;
`
