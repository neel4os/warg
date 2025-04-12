// https://nuxt.com/docs/api/configuration/nuxt-config
import materialIcons from 'quasar/icon-set/svg-material-icons'
import materialIconsRound from 'quasar/icon-set/svg-material-icons-round'
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  ssr: true,
  modules: [
    '@nuxt/content',
    '@nuxt/eslint',
    '@nuxt/fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxt/scripts',
    '@nuxt/test-utils',
    'nuxt-quasar-ui',
    'nuxt-oidc-auth'
  ],
  oidc: {
    defaultProvider: 'keycloak',
    providers: {
      keycloak: {
        audience: 'account',
        baseUrl: '',
        clientId: '',
        clientSecret: '',
        redirectUri: 'http://localhost:9999/auth/keycloak/callback',
        userNameClaim: 'preferred_username',
        logoutRedirectUri: 'http://localhost:9999/',  // Target of your post logout redirection
      },
    },
    middleware:{
      globalMiddlewareEnabled: true,
      customLoginPage: true,
    }
  },
  quasar: {
    plugins: ['Dark','Dialog', 'Notify'],
    sassVariables: true,
    iconSet: {
      ...materialIcons,
      colorPicker: materialIconsRound.colorPicker
    },
    extras: {
      fontIcons: ["material-icons"],
      font: 'roboto-font',
      animations: 'all',
    },
    appConfigKey: 'nuxtQuasarCustom',
    config: {
      dark: true,
      brand: {
        primary: '#26a69a',
        secondary: '#26c6da',
        accent: '#9C27B0',
        dark: '#1d1d1d',
        positive: '#21BA45',
        negative: '#C10015',
        info: '#31CCEC',
        warning: '#F2C037'
      }
    }
  }
})