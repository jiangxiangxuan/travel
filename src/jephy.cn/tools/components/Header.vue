<template>
  <header class="header">
    <el-row>
    <el-col :span="24">
    <el-menu :default-active="navselected" :active="navselected" class="el-menu-demo" mode="horizontal" @select="selectItems">
    <el-menu-item index="/qrcode">二维码</el-menu-item>
    <el-menu-item index="/json">json</el-menu-item>
    <el-menu-item index="/encrypt">加密/解密</el-menu-item>
    <el-submenu index="/mdconvert">
          <template slot="title">MarkDown</template>
          <el-menu-item index="/html2markdown">Html转Markdown</el-menu-item>
          <el-menu-item index="/markdown2html">Markdown转Html</el-menu-item>
    </el-submenu>
    </el-menu>
    </el-col>
    </el-row>
  </header>
</template>

<script>
  export default {
          components:{

          },
          data () {
              return {
                  navselected:"/tools/qrcode",
              }
          },
          mounted(){
            this.navselected = window.location.pathname
          },
          methods:{
              getNavType(){

              },
              selectItems(index){
                  this.$router.push(index)
              }
          },
          watch: {

          }
      }
</script>
