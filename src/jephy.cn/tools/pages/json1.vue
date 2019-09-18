<template>
    <div class="container json">
        <div class="left">
       <textarea type="text" placeholder="输入json字符串内容" v-model="data"/>
       </div>
        <div class="right">
         <button id="collapse-btn">Collapse</button>
         <button id="expand-btn">Expand</button>
         <button id="toggle-btn">Toggle</button>
         <button id="toggle-level1-btn">Toggle level1</button>
         <button id="toggle-level2-btn">Toggle level2</button>
         <div id="json-collapsed"></div>
       </div>
    </div>
</template>

<script>
    // import '../node_modules/jquery.json-viewer/json-viewer/jquery.json-viewer.js'
    export default {
     head () {
        return {
            title: 'json解析'
        }
     },
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
                $("#json-collapsed").jsonViewer(this.jsonData, {collapsed: true, nl2br: true});
            }catch(err){
                console.log(err)
            }
        }
     },
     methods: {

       }
    }
</script>
