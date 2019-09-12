<template>
    <div class="container">
       <input type="text" placeholder="json字符串" v-model="data">
       <json-viewer :value="jsonData"
       :copyable = "copyable"
       boxed
       ></json-viewer>
    </div>
</template>

<script>
    export default {
     components: {
            },
     data() {
            return {
                data: '',
                copyable:{copyText: '复制', copiedText: '复制'},
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
