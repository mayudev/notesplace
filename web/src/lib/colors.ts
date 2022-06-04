import theme from 'styled-theming'

export const BackgroundColor = theme('mode', {
  light: '#ffffff',
  dark: '#18191a',
})

export const BackgroundNavbar = theme('mode', {
  light: '#ffffff',
  dark: '#242526',
})

export const BackgroundSecondary = theme('mode', {
  light: '#ebedf0',
  dark: '#242526',
})

export const ForegroundColor = theme('mode', {
  light: '#222427',
  dark: '#f5f6f7',
})

export const ForegroudNote = theme('mode', {
  light: '#303030',
  dark: '#f5f6f7',
})

export const Accent = theme('mode', {
  light: '#1c7ed6',
  dark: '#5facec',
})

export const Error = theme('mode', {
  light: '#ec5f5f',
  dark: '#e87979',
})
