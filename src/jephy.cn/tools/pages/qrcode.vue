<template>
    <div class="container qrcode">
    <div class="left">
    <textarea type="text" placeholder="输入二维码内容" v-model="codecontent"/>
    </br>
    <button @click="genqrcode">生成</button>
    </div>
    <div class="right">
    生成结果:</br>
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
  head () {
    return {
        title: '二维码'
    }
  },
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
              text: '二维码',
              width: 200,
              height: 200,
              colorDark : '#000000',
              colorLight : '#ffffff',
              correctLevel : QRCode.CorrectLevel.H
        });

        // this.qrcode.clear();
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
        this.qrcode.clear();
        this.qrcode.makeCode(this.codecontent);
    }
  }
}

</script>
