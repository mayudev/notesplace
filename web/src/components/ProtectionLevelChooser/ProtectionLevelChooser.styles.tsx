import styled from 'styled-components'
import { Accent } from '../../lib/colors'

export const Radio = styled.label`
  display: block;
  cursor: pointer;
  padding: 0.2rem 0;

  & > input[type='radio'] {
    display: none;
  }

  & > input[type='radio'] + *::before {
    content: '';
    display: inline-block;
    vertical-align: bottom;
    width: 1rem;
    height: 1rem;
    margin-right: 0.3rem;
    border-radius: 50%;
    border-style: solid;
    border-width: 0.1rem;
    border-color: ${Accent};
  }

  & > input[type='radio']:checked + *::before {
    background: radial-gradient(${Accent} 30%, transparent 40%, transparent);
  }

  & > input[type='radio']:checked + span {
    color: ${Accent};
  }

  & > span {
    transition: 0.2s;
  }
`
