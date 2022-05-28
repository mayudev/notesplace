import styled from 'styled-components'
import { BackgroundSecondary } from '../../lib/colors'

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
`

export const Pane = styled.div``
