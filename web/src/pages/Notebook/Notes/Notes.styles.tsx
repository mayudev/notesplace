import styled from 'styled-components'
import { Accent, BackgroundSecondary } from '../../../lib/colors'

export const Container = styled.div`
  display: grid;

  grid-template-columns: repeat(3, 1fr);
  gap: 10px;

  margin: 10px;

  // Editor animation
  .editor-enter {
    opacity: 0;
    transform: scale(0.8);
  }

  .editor-enter-active {
    opacity: 1;
    transform: scale(1);
    transition: opacity 0.2s, transform 0.2s;
  }

  .editor-exit {
    opacity: 1;
  }

  .editor-exit-active {
    opacity: 0;
    transition: opacity 0.2s;
  }
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
