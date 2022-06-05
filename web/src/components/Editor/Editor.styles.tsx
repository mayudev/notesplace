import styled, { css, keyframes } from 'styled-components'
import {
  Accent,
  BackgroundColor,
  BackgroundSecondary,
  ForegroudNote,
} from '../../lib/colors'
import TextareaAutosize from 'react-textarea-autosize'

const Enter = keyframes`
  from { opacity: 0; transform: scale(0.8); }
  to { opacity: 1; transform: scale(1); }
`

type Props = {
  exiting?: boolean
}

export const Container = styled.div<Props>`
  background: ${BackgroundColor};
  border-radius: 4px;
  min-width: min(500px, 100vw);

  box-shadow: 0 4px 11px rgba(0, 0, 0, 0.25);

  border: 1px solid ${Accent};

  animation: ${Enter} 0.2s;

  @media (orientation: portrait) {
    height: 90vh;
  }

  .exiting {
    background: blue;
  }
`

export const Contents = styled.section`
  padding: 1rem;
  background: ${BackgroundSecondary};
`

export const TitleInput = styled.input`
  display: block;
  width: 100%;
  box-sizing: border-box;

  padding-bottom: 1rem;

  border: none;
  outline: none;
  resize: none;

  font: inherit;
  background: inherit;
  color: inherit;

  font-size: 1.5rem;
  font-weight: 500;
`

export const Textarea = styled(TextareaAutosize)`
  width: 100%;
  box-sizing: border-box;

  border: none;
  outline: none;
  resize: none;

  font: inherit;
  background: inherit;

  min-height: 400px;
  max-height: 60vh;

  font-size: 17px;
  line-height: 1.5rem;

  color: ${ForegroudNote};
`

// export const Textarea = styled.textarea`
//   display: block;
//   width: 100%;
//   box-sizing: border-box;

//   border: none;
//   outline: none;
//   resize: none;

//   font: inherit;
// `
