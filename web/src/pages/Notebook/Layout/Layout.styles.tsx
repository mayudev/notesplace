import styled from 'styled-components'
import { BackgroundNavbar } from '../../../lib/colors'

export const Nav = styled.nav`
  display: flex;
  align-items: center;

  padding: 0 1rem;

  background: ${BackgroundNavbar};
  box-shadow: 0px 0px 3px rgba(0, 0, 0, 0.25);
  transition: var(--transition-theme);
`

export const Title = styled.h2`
  margin: 10px 1rem;
  font-weight: 300;
`
