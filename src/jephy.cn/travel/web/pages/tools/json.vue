<template>
    <div class="container json">
        <div class="left">
       <textarea type="text" placeholder="输入json字符串内容" v-model="data"/>
       </div>
        <div class="right">
       <json-viewer :value="jsonData"
       :copyable = "copyable"
       boxed
       ></json-viewer>
       </div>
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
