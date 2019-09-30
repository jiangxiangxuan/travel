import Vue from 'vue'
import Router from 'vue-router'
import { interopDefault } from './utils'

const _0d434994 = () => interopDefault(import('..\\pages\\html2markdown.vue' /* webpackChunkName: "pages_html2markdown" */))
const _6560ea9a = () => interopDefault(import('..\\pages\\json.vue' /* webpackChunkName: "pages_json" */))
const _46c5fa87 = () => interopDefault(import('..\\pages\\json1.vue' /* webpackChunkName: "pages_json1" */))
const _1579cde8 = () => interopDefault(import('..\\pages\\markdown2html.vue' /* webpackChunkName: "pages_markdown2html" */))
const _133bd327 = () => interopDefault(import('..\\pages\\mdconvert1.vue' /* webpackChunkName: "pages_mdconvert1" */))
const _5bfc9fc0 = () => interopDefault(import('..\\pages\\qrcode.vue' /* webpackChunkName: "pages_qrcode" */))
const _59c40690 = () => interopDefault(import('..\\pages\\index.vue' /* webpackChunkName: "pages_index" */))

Vue.use(Router)

if (process.client) {
  if ('scrollRestoration' in window.history) {
    window.history.scrollRestoration = 'manual'

    // reset scrollRestoration to auto when leaving page, allowing page reload
    // and back-navigation from other pages to use the browser to restore the
    // scrolling position.
    window.addEventListener('beforeunload', () => {
      window.history.scrollRestoration = 'auto'
    })

    // Setting scrollRestoration to manual again when returning to this page.
    window.addEventListener('load', () => {
      window.history.scrollRestoration = 'manual'
    })
  }
}
const scrollBehavior = function (to, from, savedPosition) {
  // if the returned position is falsy or an empty object,
  // will retain current scroll position.
  let position = false

  // if no children detected and scrollToTop is not explicitly disabled
  if (
    to.matched.length < 2 &&
    to.matched.every(r => r.components.default.options.scrollToTop !== false)
  ) {
    // scroll to the top of the page
    position = { x: 0, y: 0 }
  } else if (to.matched.some(r => r.components.default.options.scrollToTop)) {
    // if one of the children has scrollToTop option set to true
    position = { x: 0, y: 0 }
  }

  // savedPosition is only available for popstate navigations (back button)
  if (savedPosition) {
    position = savedPosition
  }

  return new Promise((resolve) => {
    // wait for the out transition to complete (if necessary)
    window.$nuxt.$once('triggerScroll', () => {
      // coords will be used if no selector is provided,
      // or if the selector didn't match any element.
      if (to.hash) {
        let hash = to.hash
        // CSS.escape() is not supported with IE and Edge.
        if (typeof window.CSS !== 'undefined' && typeof window.CSS.escape !== 'undefined') {
          hash = '#' + window.CSS.escape(hash.substr(1))
        }
        try {
          if (document.querySelector(hash)) {
            // scroll to anchor by returning the selector
            position = { selector: hash }
          }
        } catch (e) {
          console.warn('Failed to save scroll position. Please add CSS.escape() polyfill (https://github.com/mathiasbynens/CSS.escape).')
        }
      }
      resolve(position)
    })
  })
}

export function createRouter() {
  return new Router({
    mode: 'history',
    base: decodeURI('/'),
    linkActiveClass: 'nuxt-link-active',
    linkExactActiveClass: 'nuxt-link-exact-active',
    scrollBehavior,

    routes: [{
      path: "/html2markdown",
      component: _0d434994,
      name: "html2markdown"
    }, {
      path: "/json",
      component: _6560ea9a,
      name: "json"
    }, {
      path: "/json1",
      component: _46c5fa87,
      name: "json1"
    }, {
      path: "/markdown2html",
      component: _1579cde8,
      name: "markdown2html"
    }, {
      path: "/mdconvert1",
      component: _133bd327,
      name: "mdconvert1"
    }, {
      path: "/qrcode",
      component: _5bfc9fc0,
      name: "qrcode"
    }, {
      path: "/",
      component: _59c40690,
      name: "index"
    }],

    fallback: false
  })
}
