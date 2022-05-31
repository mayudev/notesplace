import styled from 'styled-components'
import { BackgroundSecondary } from '../../lib/colors'

export const NavButton = styled.button`
  all: unset;
  font: inherit;

  display: flex;
  align-items: center;
  justify-content: center;

  width: 2rem;
  height: 2rem;

  border-radius: 50%;
  transition: background var(--transition-theme);
  cursor: pointer;

  &:hover {
    background: ${BackgroundSecondary};
  }
`
