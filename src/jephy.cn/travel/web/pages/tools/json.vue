<template>
    <div class="container">
       <input type="text" placeholder="json字符串" v-model="data">
       <json-viewer :value="jsonData"></json-viewer>
    </div>
</template>

<script>
    export default {
     components: {
            },
     data() {
            return {
                data: '',
                jsonData:{ }
            }
     },
     watch: {
        data (val) {
            try {
                this.jsonData = JSON.parse(val)
            }catch(err){

            }
        }
     },
     methods: {

       }
    }
</script>
