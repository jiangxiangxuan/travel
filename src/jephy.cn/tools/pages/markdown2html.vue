<template>
        <div class="container">
            <textarea v-model="mdText2" style="width:100%; height:250px;"></textarea>
            <button @click="toHtml">转为Html</button>
            <textarea v-model="htmlText2" style="width:100%; height:250px;"></textarea>
            <h2 style="display:none;" class="site-tips">预览</h2>
            <div style="display:none;" v-html="htmlText2" class="layui-input-block artcontent" >
            </div>
        </div>
</template>

<script>
    import '../assets/css/markdown.css'
    export default {
        head () {
                return {
                    title: 'Markdown转为Html'
                }
        },
        data () {
            return {
                htmlText1:'',
                mdText1:'',
                mdText2:'',
                htmlText2:''
            }
        },
        methods: {
            convert(md){
                var marked = require('marked');
                marked.setOptions({
                    renderer: new marked.Renderer(),
                    gfm: true,
                    tables: true,
                    breaks: false,
                    pedantic: false,
                    sanitize: false,
                    smartLists: true,
                    smartypants: false
                });
                return marked(md);
            },
            toHtml(){
                this.htmlText2=this.convert(this.mdText2);
            }
        }
    }
</script>
