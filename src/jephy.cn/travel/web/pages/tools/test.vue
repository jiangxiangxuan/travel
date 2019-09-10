<template>
  <div class="container">
    <Header/>
    test111:{{query.id}}
    <Footer/>
  </div>
</template>

<script>
  import Header from '~/components/Header.vue'
  import Footer from '~/components/Footer.vue'

  export default {
    components: {
        Header,
        Footer
    },
    async asyncData ({query}) {
          console.log(query.id);
          return {
            query: query
          }
      }
  }
</script>
