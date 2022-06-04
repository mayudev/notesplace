import styled, { keyframes } from 'styled-components'

export const Modal = styled.div`
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;

  display: flex;
  justify-content: center;
  align-items: center;
`

export const Appear = keyframes`
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
