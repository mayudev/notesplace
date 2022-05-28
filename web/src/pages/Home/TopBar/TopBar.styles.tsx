import styled from 'styled-components'
import { BackgroundNavbar } from '../../../lib/colors'

export const Nav = styled.nav`
  display: flex;
  align-items: center;
  justify-content: space-between;

  padding: 0 1rem;

  background: ${BackgroundNavbar};
  box-shadow: 0px 0px 3px rgba(0, 0, 0, 0.25);
  transition: var(--transition-theme);
`

export const Title = styled.h2`
  margin: 10px 0;
  font-weight: 300;
`

export const Fill = styled.div`
  width: 2rem;
`
