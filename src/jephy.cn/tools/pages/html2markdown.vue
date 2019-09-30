<template>
        <div class="container">
            <textarea v-model="htmlText1" style="width:100%; height:250px;"></textarea>
            <button @click="toMarkdown">转为Markdown</button>
            <textarea v-model="mdText1" style="width:100%; height:250px;"></textarea>
        </div>
</template>

<script>
    import '../assets/css/markdown.css'
    import tomarkdown from '../utils/tomarkdown'
    export default {
         head () {
            return {
                title: 'Html转为Markdown'
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
            toMarkdown(){
                this.mdText1=tomarkdown.toMarkdown(this.htmlText1)
            }
        }
    }
</script>
