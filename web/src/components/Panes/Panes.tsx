import { TransitionStatus } from 'react-transition-group'
import styled, { keyframes } from 'styled-components'
import { BackgroundSecondary } from '../../lib/colors'

const BlinkAnimation = keyframes`
  from {
    opacity: 0;
    color: inherit;
  }

  25% {
    opacity: 1;
  }

  70% {
    color: var(--color-error)
  }

  to {
    color: inherit;
  }
`

export const Blink = styled.div<{ state: TransitionStatus }>`
  text-align: center;
  padding: 12px;

  transition: 0.3s ease-out;
  animation: ${BlinkAnimation} 1s linear;
  opacity: ${({ state }) => (state === 'exiting' ? 0 : 1)};
`

export const Container = styled.div`
  max-width: 350px;
  margin: 1rem auto;
`

export const PaneHeading = styled.p`
  margin: 0;
  font-weight: 300;
  font-size: 1.7rem;
  text-align: center;
`

export const PaneSubheading = styled.p`
  margin: 0;
  margin-top: 3px;
  text-align: center;
`

export const Input = styled.input`
  all: unset;

  font: inherit;
  font-size: 1.2rem;

  display: block;
  width: 100%;

  padding: 12px;
  box-sizing: border-box;

  border-radius: 6px;
  background: ${BackgroundSecondary};

  transition: background var(--transition-theme);
`

export const ButtonContainer = styled.div`
  display: flex;
  justify-content: center;
`
