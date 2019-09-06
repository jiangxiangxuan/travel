<template>
  <div class="container">
    <Header/>
    生成二维码：
    <div id="qrcode"></div>
    识别二维码：
    <div>
    <input type="text" id="decodecode" placeholder="点击右侧按钮进行识别">
    <input type="file" id="decodefile" v-on:change="getUrl(this,'file-url')"/>
    </div>
    <Footer/>
  </div>
</template>

<script>
import Header from '~/components/Header.vue'
import Footer from '~/components/Footer.vue'

export default {
  name : 'test',
  data() {
      return {
        qrcode: {}
      }
  },
  components: {
        Header,
        Footer
  },
  watch:{

  },
  mounted () {
        this.qrcode = new QRCode('qrcode', {
              text: '永远相信美好的事情即将发生',
              width: 200,
              height: 200,
              colorDark : '#000000',
              colorLight : '#ffffff',
              correctLevel : QRCode.CorrectLevel.H
        });

        this.qrcode.clear();
        // this.qrcode.makeCode('new content');
  },
  methods: {
    getUrl(e,param){
    		analyticCode.getUrl(
    			param,document.getElementById('decodefile'),
    			function(imgMsg,url){
    			document.getElementById('decodecode').value=imgMsg;
    		}
    		)
    	}
  }
}

</script>
