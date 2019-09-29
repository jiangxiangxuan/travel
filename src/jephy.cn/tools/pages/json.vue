<template>
    <div class="container json">
        <div class="left">
       <textarea type="text" placeholder="输入json字符串内容" v-model="data"/>
       </div>
        <div class="right">
            <input type="button" @click="copy" value="点此复制">
            <textarea type="text" v-model="jsonData" id="json-view"></textarea>
            <!--<div id="json-view"></div>-->
       </div>
    </div>
</template>

<script>
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
                jsonData: ''
            }
     },
     watch: {
        data (val) {
            try {
                var json = JSON.parse(val)
                this.jsonData = JSON.stringify(json, undefined, 4)
                $('#json-view').val(this.highLight(this.jsonData))
            }catch(err){
                console.log(err)
            }
        }
     },
     mounted () {
     },
     methods: {
        highLight (json) {
			 json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
			 return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
				     var cls = 'number';
				     if (/^"/.test(match)) {
				         if (/:$/.test(match)) {
				             cls = 'key';
				         } else {
				             cls = 'string';
				         }
				     } else if (/true|false/.test(match)) {
				         cls = 'boolean';
				     } else if (/null/.test(match)) {
				         cls = 'null';
				     }
				     return '<span class="' + cls + '">' + match + '</span>';
			 });
        },
        formatJson (jsonStr) {
            var str = JSON.stringify(jsonStr, undefined, 4);
            return this.highLight(str)
        },
        copy () {
            var obj=document.getElementById("json-view")
            obj.select();
            document.execCommand("copy");
        }
     }
    }
</script>

<style scoped>
#json-view {
    display: block;
    font-family: monospace;
    white-space: pre;
    margin: 1em 0px;
    width: 430px;
    height: 565px;
}
.key{
    color: red;
}
.string{
    color:green;
}
.number{
    color:orange;
}
</style>
