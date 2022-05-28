import styled from 'styled-components'
import { BackgroundSecondary } from '../../lib/colors'

export const Container = styled.div`
  max-width: 350px;
  margin: auto;
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
  margin: 1rem auto;
  box-sizing: border-box;

  border-radius: 6px;
  background: ${BackgroundSecondary};

  transition: background var(--transition-theme);
`

export const ButtonContainer = styled.div`
  display: flex;
  justify-content: center;
`
