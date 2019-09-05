<template>
  <div class="container">
    <Header/>
    <v-tab :items="tabItems">
            <div slot="generatecode">
            <label class="layui-form-label">文本/Url</label>
            <input type="text" class="layui-input" placeholder="将要生成二维码的文本或链接" v-model="dataTxt"/>

                <div class="layui-form-item layui-form-text">
                    <label class="layui-form-label">颜色背景</label>
                    <div class="layui-input-inline">
                        <input type="text" class="layui-input" placeholder="颜色 默认#ffffff" v-model="color"/>
                    </div>
                    <div class="layui-input-inline">
                        <input type="text" class="layui-input" placeholder="背景 默认#1aa094" v-model="bgcolor"/>
                    </div>
                </div>
                <div class="layui-form-item">
                    <div class="layui-input-block">
                        <button type="button" v-on:click="convert">生成</button>
                    </div>
                </div>
                <div class="layui-form-item">
                    <div class="layui-input-block">
                        <img :src="qrcodeImg" />
                    </div>
                </div>
            </div>
        </v-tab>
    <Footer/>
  </div>
</template>

<script>
  import jrQrcode from 'jr-qrcode';
  import Header from '~/components/Header.vue'
  import Footer from '~/components/Footer.vue'

  export default {
    data(){
        return {
            tabItems:[
                {
                    Name:'generatecode',
                    Title:'二维码生成'
                }
            ],
            qrcodeImg:'',
            dataTxt:'',
            color:'',
            bgcolor:''
        }
    },
    components: {
        Header,
        Footer
    },
    created(){
    },
    watch:{
        'dataTxt':'zidong'
    },
    methods:{
        zidong(){
            this.convert(this.dataTxt,this.color,this.bgcolor);
        },
        convert(){
            var imgBase64 = jrQrcode.getQrBase64(this.dataTxt,{
                padding		: 10,   //二维码四边空白，默认为10px
                width		: 256,  //二维码图片宽度，默认为256px
                height		: 256,  //二维码图片高度，默认为256px
                correctLevel	: QRErrorCorrectLevel.H,    //二维码容错level，默认为高
                background      : this.color||"#ffffff",    //二维码颜色
                foreground      : this.bgcolor||"#1aa094"     //二维码背景颜色
            });
            this.qrcodeImg=imgBase64;
        }
    },
    async asyncData ({query}) {
          console.log(query.id);
          return {
            query: query
          }
      }
  }
</script>
