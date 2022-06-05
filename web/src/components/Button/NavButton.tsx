import styled from 'styled-components'
import { Accent, BackgroundSecondary, Error } from '../../lib/colors'

type Props = {
  withText?: boolean
  error?: boolean
  primary?: boolean
}

export const NavButton = styled.button<Props>`
  all: unset;
  font: inherit;

  display: flex;
  align-items: center;
  justify-content: center;

  width: ${props => (props.withText ? '' : '2rem')};
  height: 2rem;

  border-radius: ${props => (props.withText ? '20px' : '50%')};
  padding: ${props => (props.withText ? '0 12px' : '')};

  transition: background var(--transition-theme);
  cursor: pointer;

  color: ${props => {
    if (props.error) return Error
    if (props.primary) return Accent
  }};

  &:hover {
    background: ${BackgroundSecondary};
  }

  span {
    padding-left: 0.3rem;
  }
`
