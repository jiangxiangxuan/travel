<template>
  <div class="body">
  <Header/>
  <nuxt/>
  <Footer/>
  </div>
</template>

<script>
import Header from '~/components/Header.vue'
import Footer from '~/components/Footer.vue'

export default {
  name : 'default',
  data() {
      return {
      }
  },
  components: {
        Header,
        Footer
  }
}

</script>
