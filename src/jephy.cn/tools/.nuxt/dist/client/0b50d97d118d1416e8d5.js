(window.webpackJsonp=window.webpackJsonp||[]).push([[5],{276:function(e,t,o){"use strict";o.r(t);var n={name:"qrcode",head:function(){return{title:"二维码"}},data:function(){return{codecontent:"",qrcode:{}}},components:{},watch:{},mounted:function(){this.qrcode=new QRCode("qrcode",{text:"二维码",width:200,height:200,colorDark:"#000000",colorLight:"#ffffff",correctLevel:QRCode.CorrectLevel.H})},methods:{getUrl:function(e,param){analyticCode.getUrl(param,document.getElementById("decodefile"),function(e,t){document.getElementById("decodecode").value=e})},genqrcode:function(){this.qrcode.clear(),this.qrcode.makeCode(this.codecontent)}}},c=o(13),component=Object(c.a)(n,function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{staticClass:"container qrcode"},[o("div",{staticClass:"left"},[o("textarea",{directives:[{name:"model",rawName:"v-model",value:e.codecontent,expression:"codecontent"}],attrs:{type:"text",placeholder:"输入二维码内容"},domProps:{value:e.codecontent},on:{input:function(t){t.target.composing||(e.codecontent=t.target.value)}}}),o("br"),e._v(" "),o("button",{on:{click:e.genqrcode}},[e._v("生成")])]),e._v(" "),e._m(0)])},[function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"right"},[this._v("\n生成结果:"),t("br"),this._v(" "),t("div",{attrs:{id:"qrcode"}}),this._v(" "),t("br")])}],!1,null,null,null);t.default=component.exports}}]);