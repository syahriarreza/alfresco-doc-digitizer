const path = require('path')

export default {
  /*
   ** Nuxt rendering mode
   ** See https://nuxtjs.org/api/configuration-mode
   */
  mode: 'spa',
  devTools: process.env.NODE_ENV !== 'production',
  /* Server and variables */
  server: {
    port: process.env.LISTEN_PORT || 3000,
    host: process.env.LISTEN_HOST || 'localhost',
  },
  env: {
    baseUrl:
      process.env.SITE_BASE_URL || 'https://local.dev.kano.app:3000/giias/',
  },
  router: {
    base: '/admin/',
  },
  /*
   ** Nuxt target
   ** See https://nuxtjs.org/api/configuration-target
   */
  target: 'server',
  /*
   ** Headers of the page
   ** See https://nuxtjs.org/api/configuration-head
   */
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      {
        charset: 'utf-8',
      },
      {
        name: 'viewport',
        content: 'width=device-width, initial-scale=1',
      },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      },
    ],
    link: [
      {
        rel: 'icon',
        type: 'image/png',
        href: 'favicon.png',
      },
    ],
  },
  /*
   ** Global CSS
   */
  css: ['../../webs/shared/assets/base.scss'],
  /*
   ** Plugins to load before mounting the App
   ** https://nuxtjs.org/guide/plugins
   */
  plugins: [
    '~/plugins/cookie.js',
    '~/plugins/wysiwyg.js',
    '~/plugins/validation.js',
    '~/plugins/casl.js',
    '~/plugins/localStorage.js',
    '../../webs/shared/plugins/axios.js',
    '../../webs/shared/plugins/tool.js',
    { src: '~/plugins/mask.js', ssr: false },
    {
      src: '../../webs/shared/plugins/swiper.js',
      ssr: false,
    },
    { src: '~/plugins/vue-html2pdf', mode: 'client' },
  ],
  /*
   ** Auto import components
   ** See https://nuxtjs.org/api/configuration-components
   */
  components: true,
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: ['@nuxtjs/vuetify', '@nuxtjs/moment'],
  /*
   ** Nuxt.js modules
   */
  modules: [
    [
      'cookie-universal-nuxt',
      {
        alias: 'cookie',
      },
    ],
    '@nuxtjs/axios',
    'nuxt-webfontloader',
    'vuetify-dialog/nuxt',
  ],
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {
    baseURL: process.env.SITE_API_URL || 'https://local.dev.kano.app:36080/v1/',
  },
  /*
   ** vuetify module configuration
   ** https://github.com/nuxt-community/vuetify-module
   */
  vuetify: {
    customVariables: ['../../webs/shared/assets/variables.scss'],
    treeShake: true,
    theme: {
      dark: false,
      themes: {
        light: {
          primary: '#3F51B5',
          secondary: '#757575',
          tertiary: '#009688',
          accent: '#FFC107',
          error: '#F44336',
          info: '#2196F3',
          success: '#4CAF50',
          warning: '#FFC107',
        },
        dark: {
          primary: '#303F9F',
          secondary: '#757575',
          tertiary: '#009688',
          accent: '#FFC107',
          error: '#D32F2F',
          info: '#1976D2',
          success: '#388E3C',
          warning: '#FFA000',
        },
      },
      options: {
        customProperties: true,
      },
    },
    defaultAssets: {
      /*
       ** Remove roboto as default font
       ** therefor it can use webfontloader
       */
      font: false,
      icons: 'mdi',
    },
  },

  vuetifyDialog: {
    property: '$dialog',
    confirm: {
      actions: {
        false: 'Tidak',
        true: {
          text: 'Iya',
          color: 'primary',
        },
      },
      title: 'Konfirmasi',
      width: 400,
      persistent: true,
    },
  },

  /*
   ** Build configuration
   ** See https://nuxtjs.org/api/configuration-build/
   */
  build: {
    extend(config) {
      config.resolve.alias['~shared'] = path.resolve(__dirname, '../shared')
    },
  },
  webfontloader: {
    google: {
      families: [
        'Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900',
        'Open+Sans:ital,wght@0,300;0,400;0,600;0,700;0,800;1,300;1,400;1,600;1,700;1,800',
      ],
    },
  },
  publicRuntimeConfig: {
    api: {
      baseUrl: process.env.API_BASE_URL || 'https://local.dev.kano.app:36080/',
      asset: process.env.API_ASSET || 'v1/asset/',
    },
  },
  privateRuntimeConfig: {},
}
