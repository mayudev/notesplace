import styled, { keyframes } from 'styled-components'
import { BackgroundSecondary } from '../../lib/colors'

const Appear = keyframes`
  from { opacity: 0; transform: scale(0.9); }
  to { opacity: 1; }
`

export const Container = styled.div`
  animation: ${Appear} 0.34s ease-out;
`

export const Hero = styled.section`
  max-width: 500px;
  margin: 15vh auto;
`

export const Heading = styled.div`
  text-align: center;
  font-size: 2.5rem;
  font-weight: 500;
`

export const Subheading = styled.div`
  text-align: center;
  font-size: 1.2rem;

  margin-top: 1rem;
`

export const Center = styled.div`
  text-align: center;
  font-weight: 300;
  font-size: 1.2rem;

  padding: 1rem;
  margin-bottom: 1rem;

  background: ${BackgroundSecondary};

  transition: background var(--transition-theme);
`

export const Panes = styled.section`
  display: grid;
  grid-template-columns: 50% 50%;

  max-width: 1080px;
  margin: auto;
`

export const Pane = styled.div``
