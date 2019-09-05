<template>
  <div>
    <img :src="qrcodeImg" />
  </div>
</template>

<script>
if (process.BROWSER_BUILD) {
  require('jr-qrcode')
}
var jrQrcode = require('jr-qrcode');

export default {
  name : 'test',
  data() {
      return {
        qrcodeImg:'',
            dataTxt:'test',
            color:'',
            bgcolor:'',
      }
    },
       components: {
         jrQrcode,
       },
  mounted () {
        var imgBase64 = jrQrcode.getQrBase64(this.dataTxt,{
                padding		: 10,   //二维码四边空白，默认为10px
                width		: 256,  //二维码图片宽度，默认为256px
                height		: 256,  //二维码图片高度，默认为256px
                correctLevel	: QRErrorCorrectLevel.H,    //二维码容错level，默认为高
                background      : this.color||"#ffffff",    //二维码颜色
                foreground      : this.bgcolor||"#1aa094"     //二维码背景颜色
            });
            this.qrcodeImg=imgBase64;
  },
  methods: {

  }
}

</script>
