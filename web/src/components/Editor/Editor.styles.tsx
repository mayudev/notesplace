import styled from 'styled-components'
import {
  Accent,
  BackgroundColor,
  BackgroundSecondary,
  ForegroudNote,
} from '../../lib/colors'
import TextareaAutosize from 'react-textarea-autosize'

export const Container = styled.div`
  background: ${BackgroundColor};
  border-radius: 4px;
  min-width: 600px;

  box-shadow: 0 4px 11px rgba(0, 0, 0, 0.25);

  border: 1px solid ${Accent};
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
