<template>
    <div class="container encrypt">
            <div class="row">
                <label>明文</label>
                <div class="right">
                <textarea v-model="model.beforeTxt" placeholder="加密/解密字符串..."></textarea>
                </div>
            </div>
            <div class="row">
                <label>加密类型</label>
                <div class="right">
                    <ul>
                    <li v-for="item in encryptTypes">
                    <input class="encrypttype" name="encrypttype" type="radio" @input="currentHandler(item)">{{item}}</input>
                    </li>
                    </ul>
                </div>
            </div>
            <div class="row">
                <label>密匙</label>
                <div class="right">
                    <input class="input" type="text" v-model="model.pwdTxt" placeholder="密匙"/>
                </div>
            </div>
            <div class="row">
                <div>
                    <button @click="encrypt(model.currentType)">加密</button>
                    <button @click="decrypt(model.currentType)">解密</button>
                </div>
            </div>
            <div class="row">
                <label>密文</label>
                <div class="right">
                <textarea v-model="model.afterTxt" placeholder="加密后字符串..."></textarea>
                </div>
            </div>
    </div>
</template>
<script>
    import CryptoJS from 'crypto-js'
    export default {
        data() {
            return {
                model:{
                    beforeTxt:'',
                    afterTxt:'',
                    pwdTxt:'',
                    currentType:'AES',
                },
                encryptTypes:['AES','DES','RC4','Rabbit','TripleDES','Base64'],
                hashTypes:['MD5','SHA1','SHA224','SHA256','SHA384','SHA512','HmacSHA1','HmacSHA224','HmacSHA256','HmacSHA384','HmacSHA512','HmacMD5'],
                hashPwdTypes:['HmacSHA1','HmacSHA224','HmacSHA256','HmacSHA384','HmacSHA512','HmacMD5'],
                tabType:0,
            }
        },
        watch:{
            '$route':'fetchData'
        },
        mounted () {
            $(".encrypttype").eq(0).attr('checked', 'true');
        },
        created() {
            this.fetchData()
        },
        methods: {
            fetchData(){
                this.model={
                    beforeTxt:'',
                    afterTxt:'',
                    pwdTxt:'',
                    currentType:'',
                };
                this.tabType=this.$route.query.tab||1;
                if(this.tabType==1){
                    this.model.currentType=this.encryptTypes[0];
                }
                else{
                    this.model.currentType=this.hashTypes[0]
                }
            },
            switchTab(tab){
                this.$options.filters.openRoute({path:this.$route.fullpath,query:{tab:tab}})
            },
            currentHandler(e){
                this.model.currentType=e;
                console.log(e);
            },
            encrypt(type){
                let self=this;
                type=type||self.encryptTypes[0];
                let data={
                    beforeTxt:this.model.beforeTxt||'',
                    afterTxt:this.model.afterTxt||'',
                    pwdTxt:this.model.pwdTxt||''
                };
                if(type=='Base64'){
                    var str=CryptoJS.enc.Utf8.parse(data.beforeTxt);
                    self.model.afterTxt=CryptoJS.enc.Base64.stringify(str);
                    return;
                }
                self.model.afterTxt=CryptoJS[type].encrypt(data.beforeTxt,data.pwdTxt).toString();

            },
            decrypt(type){
                let self=this;
                type=type||self.encryptTypes[0];
                let data={
                    afterTxt:this.model.afterTxt||'',
                    pwdTxt:this.model.pwdTxt||'',
                };
                if(type=='Base64'){
                    var words  = CryptoJS.enc.Base64.parse(data.afterTxt);
                    self.model.beforeTxt=words.toString(CryptoJS.enc.Utf8);

                    return;
                }
                self.model.beforeTxt=CryptoJS[type].decrypt(data.afterTxt,data.pwdTxt).toString(CryptoJS.enc.Utf8);

            },
            hashEncrypt(type){
                let self=this;
                type=type||self.hashTypes[0];
                let data={
                    beforeTxt:this.model.beforeTxt||'',
                    afterTxt:this.model.afterTxt||'',
                    pwdTxt:this.model.pwdTxt||''
                };
                self.model.afterTxt=CryptoJS[type](data.beforeTxt,data.pwdTxt).toString();
            }
        }
    }

</script>

<style scoped>
.encrypt ul {
    margin: 0;
}
.encrypt ul li {
    list-style: none;
    padding: 0px;
    margin: 0px;
    width: auto;
    margin-left: 5px;
    float: left;
}
.encrypt .row  {
    width: 100%;
    float: left;
    margin-top: 10px;
}
.encrypt .row label {
    float: left;
}
.encrypt .row .input {
    float: left;
    margin-left: 10px;
    width: 600px;
}
.encrypt .row textarea {
    float: left;
    margin-left: 10px;
    width: 600px;
    height: 60px;
}
</style>
