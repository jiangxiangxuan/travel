<template>
    <div style="width:100%;height:100%;" class="layui-tab-brief">
        <div style="width:100%;height:100%;" v-if="tabType==2">
            <div class="layui-form-item layui-form-text">
                <label class="layui-form-label">明文</label>
                <div class="layui-input-block">
                <textarea v-model="model.beforeTxt" placeholder="加密/解密字符串..." class="layui-textarea"></textarea>
                </div>
            </div>
            <div class="layui-form-item layui-form-text">
                <label class="layui-form-label">加密类型</label>
                <div class="layui-input-block">
                    <v-radio :value="model.currentType" :optionVals="hashTypes" @input="currentHandler"></v-radio>
                </div>
            </div>
            <div class="layui-form-item layui-form-text" v-if="hashPwdTypes.indexOf(model.currentType)>-1">
                <label class="layui-form-label">密匙</label>
                <div class="layui-input-block">
                    <input type="text" v-model="model.pwdTxt" placeholder="密匙" class="layui-input" />
                </div>
            </div>
            <div class="layui-form-item layui-form-text">
                <div class="layui-input-block">
                    <v-button @click="hashEncrypt(model.currentType)">
                        <span><i class="layui-icon">&#xe61a; </i>加密</span>
                    </v-button>
                </div>
            </div>
            <div class="layui-form-item layui-form-text">
                <label class="layui-form-label">密文</label>
                <div class="layui-input-block">
                <textarea v-model="model.afterTxt" placeholder="加密后字符串..." class="layui-textarea"></textarea>
                </div>
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
                    currentType:'',
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