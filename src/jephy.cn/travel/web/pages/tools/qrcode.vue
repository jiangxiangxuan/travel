<template>
    <div class="container">
    <div>
    生成二维码：<input type="text" placeholder="生成二维码字符串" v-model="codecontent"><button @click="genqrcode">生成</button>
    </br>
    <div id="qrcode"></div>
     </br>
    </div>
    <!--<div>
    识别二维码：
    <div>
    <input type="text" id="decodecode" placeholder="点击右侧按钮进行识别">
    <input type="file" id="decodefile" v-on:change="getUrl(this,'file-url')"/>
    </div>
    </div>-->
    </div>
</template>

<script>
export default {
  name : 'qrcode',
  data() {
      return {
        codecontent: '',
        qrcode: {}
      }
  },
  components: {

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
    },
    genqrcode() {
        this.qrcode.makeCode(this.codecontent);
    }
  }
}

</script>
