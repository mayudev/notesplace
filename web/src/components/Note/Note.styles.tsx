import styled, { keyframes } from 'styled-components'
import { Accent, BackgroundSecondary } from '../../lib/colors'

const Appear = keyframes`
  from { opacity: 0; transform: translateY(30px) }
  to { opacity: 1} 
`

export const Container = styled.div`
  animation: ${Appear} 0.4s ease-out;

  background: ${BackgroundSecondary};
  border-radius: 6px;

  transition: var(--transition-theme);
  cursor: pointer;

  padding: 12px;

  &:hover {
    color: ${Accent};
  }
`

export const Title = styled.div`
  font-weight: 600;
  font-size: 1.1rem;
`

export const Content = styled.div``
