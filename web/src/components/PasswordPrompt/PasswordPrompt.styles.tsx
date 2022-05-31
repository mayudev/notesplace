import styled from 'styled-components'
import { Accent, BackgroundColor } from '../../lib/colors'

export const Container = styled.div`
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;

  display: flex;
  justify-content: center;
  align-items: center;
`

export const Backdrop = styled.div`
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;

  background: #00000070;
  z-index: -1;
`

export const Contents = styled.div`
  background: ${BackgroundColor};

  border-radius: 4px;
  border: 1px solid ${Accent};

  min-width: 400px;
`

export const Message = styled.div`
  padding: 1rem;
  font-weight: 300;
  font-size: 1.2rem;
`

export const Buttons = styled.div`
  display: flex;
  flex-direction: row-reverse;

  padding: 1rem;
`
