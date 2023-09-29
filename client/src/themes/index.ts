import type { ThemeDefinition } from 'vuetify'
import { merge } from 'lodash'

const general: ThemeDefinition = {
  colors: {
    error: '#D50000',
    success: '#00C853',
    info: '#00BFA5',
    warning: '#FFD600'
  }
}

const light: ThemeDefinition = {
  dark: false,
  colors: {
    background: '#F5F5F5',
    surface: '#FFFFFF',
    primary: '#004D40',
    secondary: '#00796B',
    chart: '#515152'
  }
}

const dark: ThemeDefinition = {
  dark: true,
  colors: {
    background: '#212121',
    surface: '#424242',
    primary: '#26C6DA',
    secondary: '#0096a8',
    chart: '#656565'
  }
}

merge(light, general)
merge(dark, general)

export { light, dark }
