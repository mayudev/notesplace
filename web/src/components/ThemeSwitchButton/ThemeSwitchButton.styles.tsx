import styled from 'styled-components'
import { BackgroundSecondary } from '../../lib/colors'

export const Container = styled.div`
  width: 2rem;
  height: 2rem;
`

export const Button = styled.button`
  all: unset;
  font: inherit;

  display: flex;
  align-items: center;
  justify-content: center;

  width: 100%;
  height: 100%;

  border-radius: 50%;
  transition: background var(--transition-theme);
  cursor: pointer;

  &:hover {
    background: ${BackgroundSecondary};
  }
`
