<template>
  <div class="footer">
    Copyright ©2019-2029 iteam. All Rights Reserved. Current version is 1.0.0.
    <a id="beian" href="http://www.beian.miit.gov.cn" target="_blank">粤ICP备19053447号</a>
  </div>
</template>

<script>
    export default {
     head () {
        return {
            beiantext: ""
        }
     },
     components: {

     },
     mounted () {
        this.beiantext = "粤ICP备19053447号"
        if(document.domain == "www.jephy.cn"){
            this.beiantext = "粤ICP备19053447号-2"
        }else if(document.domain == "www.jephy.top"){
            this.beiantext = "粤ICP备19053447号-1"
        }
        $("#beian").text(this.beiantext)
     },
     data() {
     },
     watch: {
     },
     methods: {

     },
     async asyncData ({params, error}) {
        console.log(this.$nuxt._route.params.id)
        console.log(params.id)
     }
    }
</script>
