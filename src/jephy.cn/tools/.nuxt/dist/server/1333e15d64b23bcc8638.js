exports.ids=[4],exports.modules={24:function(e,t,o){"use strict";o.r(t);var d={name:"qrcode",head:()=>({title:"二维码"}),data:()=>({codecontent:"",qrcode:{}}),components:{},watch:{},mounted(){this.qrcode=new QRCode("qrcode",{text:"二维码",width:200,height:200,colorDark:"#000000",colorLight:"#ffffff",correctLevel:QRCode.CorrectLevel.H})},methods:{getUrl(e,param){analyticCode.getUrl(param,document.getElementById("decodefile"),function(e,t){document.getElementById("decodecode").value=e})},genqrcode(){this.qrcode.clear(),this.qrcode.makeCode(this.codecontent)}}},c=o(1),component=Object(c.a)(d,function(){var e=this.$createElement;return(this._self._c||e)("div",{staticClass:"container qrcode"},[this._ssrNode('<div class="left"><textarea type="text" placeholder="输入二维码内容">'+this._ssrEscape(this._s(this.codecontent))+'</textarea><br> <button>生成</button></div> <div class="right">\n生成结果:<br> <div id="qrcode"></div> <br></div>')])},[],!1,null,null,"7607e768");t.default=component.exports}};
//# sourceMappingURL=1333e15d64b23bcc8638.js.map