import styled from 'styled-components'
import { Accent, BackgroundSecondary } from '../../../lib/colors'

export const Container = styled.div`
  display: grid;

  grid-template-columns: repeat(3, 1fr);
  gap: 10px;

  margin: 10px;
`

export const CreateNote = styled.button`
  all: unset;
  display: flex;
  justify-content: center;
  align-items: center;

  padding: 38px;

  background: ${BackgroundSecondary};
  border-radius: 6px;
  border: 2px solid transparent;

  transition: var(--transition-theme);
  cursor: pointer;

  &:hover {
    border: 2px solid ${Accent};
    color: ${Accent};
  }
`
