<template>
  <header class="header">
    <el-row>
    <el-col :span="24">
    <el-menu :default-active="navselected" :active="navselected" class="el-menu-demo" mode="horizontal" @select="selectItems">
    <el-menu-item index="/qrcode">二维码</el-menu-item>
    <el-menu-item index="/json">json</el-menu-item>
    <!--
    <el-menu-item index="3">客户管理</el-menu-item>
    <el-menu-item index="4">系统设置</el-menu-item>
    <el-menu-item index="5">活动发布</el-menu-item>
    <el-submenu index="6">
          <template slot="title">选项4</template>
          <el-menu-item index="2-4-1">选项1</el-menu-item>
          <el-menu-item index="2-4-2">选项2</el-menu-item>
          <el-menu-item index="2-4-3">选项3</el-menu-item>
        </el-submenu>
      </el-submenu>
      -->
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
