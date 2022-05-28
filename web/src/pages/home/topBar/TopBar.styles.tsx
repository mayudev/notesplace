import styled from 'styled-components'
import { BackgroundSecondary } from '../../../lib/colors'

export const Nav = styled.nav`
  display: flex;
  justify-content: center;

  background: ${BackgroundSecondary};
  box-shadow: 0px 0px 3px rgba(0, 0, 0, 0.25);
`

export const Title = styled.h2`
  margin: 10px 0;
  font-weight: 300;
`
