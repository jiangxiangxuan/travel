
export default {
  mode: 'universal',
  /*
  ** Headers of the page
  */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    script: [
      { src: '/qrcode/qrcode.min.js', type: 'text/javascript', charset: 'utf-8', ssr: false},
      { src: '/qrcode/llqrcode.js', type: 'text/javascript', charset: 'utf-8', ssr: false},
      { src: '/qrcode/webqr.js', type: 'text/javascript', charset: 'utf-8', ssr: false},
      { src: '/qrcode/analyticCode.js', type: 'text/javascript', charset: 'utf-8', ssr: false}
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: '/css/layout.css' }
    ]
  },
  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },
  /*
  ** Global CSS
  */
  css: [
    'element-ui/lib/theme-chalk/index.css'
  ],
  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    {src: '~/plugins/ElementUI', ssr: true },
    {src: '~/plugins/JsonView', ssr: false }
  ],
  /*
  ** Nuxt.js dev-modules
  */
  devModules: [
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
  ],
  /*
  ** Build configuration
  */
  build: {
    // 防止element-ui被多次打包
    vendor: ['element-ui','vue-json-views'],
    /*
    ** You can extend webpack config here
    */
    extend (config, ctx) {
    }
  }
}
