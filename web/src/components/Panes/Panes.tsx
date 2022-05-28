import styled from 'styled-components'
import { BackgroundSecondary } from '../../lib/colors'

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
  max-width: 350px;
  margin: 1rem auto;

  border-radius: 6px;
  background: ${BackgroundSecondary};
`

export const ButtonContainer = styled.div`
  display: flex;
  justify-content: center;
`
