(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-695a"],{"2TQq":function(t,e,o){"use strict";var a=o("P2sY"),r=o.n(a),n=o("glbJ"),s=o("fL+G"),i=o("Grqa"),l=o("uARZ"),c=o("Yfch"),u=o("RaKr"),d=o("gg54"),m={_id:"",code:"",type:"",value:"",active:"",min_consomm:"",begin_date:"",end_date:""},p={name:"PromoCodeDetail",components:{Tinymce:n.a,MDinput:i.a,Upload:s.a,Sticky:l.a},props:{isEdit:{type:Boolean,default:!1}},data:function(){var t=this,e=function(e,o,a){""===o?(t.$message({message:e.field+"为必传项",typlengthe:"error"}),a(new Error(e.field+"为必传项"))):a()};return{postForm:r()({},m),loading:!1,userListOptions:[],rules:{image_uri:[{validator:e}],title:[{validator:e}],content:[{validator:e}],source_uri:[{validator:function(e,o,a){o?Object(c.b)(o)?a():(t.$message({message:"外链url填写不正确",type:"error"}),a(new Error("外链url填写不正确"))):a()},trigger:"blur"}]},tempRoute:{}}},computed:{lang:function(){return this.$store.getters.language}},created:function(){if(this.isEdit){var t=this.$route.params&&this.$route.params.id;this.fetchData(t)}else this.postForm=r()({},m);this.tempRoute=r()({},this.$route)},methods:{fetchData:function(t){var e=this;Object(u.c)(t).then(function(t){e.postForm=t.data,e.setTagsViewTitle()}).catch(function(t){console.log(t)})},setTagsViewTitle:function(){var t="zh"===this.lang?"编辑文章":"Edit Article",e=r()({},this.tempRoute,{title:t+"-"+this.postForm.id});this.$store.dispatch("updateVisitedView",e)},submitForm:function(){var t=this;this.$refs.postForm.validate(function(e){if(!e)return console.log("error submit!!"),!1;t.$route.params.id?Object(u.d)(t.$route.params.id,t.postForm).then(function(e){t.loading=!0;var o="成功",a="发布文章成功",r="success";-1===e.data.code&&(r="error",o="failed",a="save failed"),t.$notify({title:o,message:a,type:r,duration:2e3}),t.loading=!1},function(){t.$notify({title:"failed",message:"save failed",type:"error",duration:2e3})}):Object(u.a)(t.postForm).then(function(e){t.loading=!0;var o="成功",a="发布文章成功",r="success";-1===e.data.code&&(r="error",o="failed",a="save failed"),t.$notify({title:o,message:a,type:r,duration:2e3}),t.loading=!1,t.$router.push("/promoCode/list")},function(){t.$notify({title:"failed",message:"save failed",type:"error",duration:2e3})})})},draftForm:function(){0!==this.postForm.content.length&&0!==this.postForm.title.length?(this.$message({message:"保存成功",type:"success",showClose:!0,duration:1e3}),this.postForm.status="draft"):this.$message({message:"请填写必要的标题和内容",type:"warning"})},getRemoteUserList:function(t){var e=this;Object(d.a)(t).then(function(t){t.data.items&&(e.userListOptions=t.data.items.map(function(t){return t.name}))})}}},f=(o("uwli"),o("KHd+")),h=Object(f.a)(p,function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",{staticClass:"createPost-container"},[o("el-form",{ref:"postForm",staticClass:"form-container",attrs:{model:t.postForm,rules:t.rules}},[o("sticky",{attrs:{"class-name":"sub-navbar "+t.postForm.status}},[o("el-button",{directives:[{name:"loading",rawName:"v-loading",value:t.loading,expression:"loading"}],staticStyle:{"margin-left":"10px"},attrs:{type:"success"},on:{click:t.submitForm}},[t._v("发布\n      ")])],1),t._v(" "),o("div",{staticClass:"createPost-main-container"},[o("el-row",[o("div",{staticClass:"postInfo-container"},[o("el-row",[o("el-col",{attrs:{span:24}},[o("el-form-item",{attrs:{label:t.$t("promoCode.title"),"label-width":"80px"}},[o("el-input",{attrs:{type:"text",autosize:"",placeholder:"code"},model:{value:t.postForm.code,callback:function(e){t.$set(t.postForm,"code",e)},expression:"postForm.code"}})],1)],1)],1),t._v(" "),o("el-row",[o("el-col",{attrs:{span:24}},[o("el-form-item",{attrs:{label:t.$t("promoCode.type"),"label-width":"80px"}},[o("el-input",{attrs:{placeholder:t.$t("promoCode.type"),type:"text"},model:{value:t.postForm.type,callback:function(e){t.$set(t.postForm,"type",e)},expression:"postForm.type"}})],1)],1)],1),t._v(" "),o("el-row",[o("el-col",{attrs:{span:24}},[o("el-form-item",{attrs:{label:t.$t("promoCode.value"),"label-width":"80px"}},[o("el-input-number",{attrs:{placeholder:t.$t("promoCode.value")},model:{value:t.postForm.value,callback:function(e){t.$set(t.postForm,"value",e)},expression:"postForm.value"}})],1)],1)],1),t._v(" "),o("el-row",[o("el-col",{attrs:{span:24}},[o("el-form-item",{attrs:{label:t.$t("promoCode.min_consomm"),"label-width":"80px"}},[o("el-input-number",{attrs:{placeholder:t.$t("promoCode.min_consomm")},model:{value:t.postForm.min_consomm,callback:function(e){t.$set(t.postForm,"min_consomm",e)},expression:"postForm.min_consomm"}})],1)],1)],1),t._v(" "),o("el-row",[o("el-col",{attrs:{span:24}},[o("el-form-item",{attrs:{label:t.$t("promoCode.begin_date"),"label-width":"80px"}},[o("el-date-picker",{attrs:{placeholder:t.$t("promoCode.begin_date"),type:"date"},model:{value:t.postForm.begin_date,callback:function(e){t.$set(t.postForm,"begin_date",e)},expression:"postForm.begin_date"}})],1)],1)],1),t._v(" "),o("el-row",[o("el-col",{attrs:{span:24}},[o("el-form-item",{attrs:{label:t.$t("promoCode.end_date"),"label-width":"80px"}},[o("el-date-picker",{attrs:{placeholder:t.$t("promoCode.end_date"),type:"date"},model:{value:t.postForm.end_date,callback:function(e){t.$set(t.postForm,"end_date",e)},expression:"postForm.end_date"}})],1)],1)],1)],1)])],1)],1)],1)},[],!1,null,"cd2289a2",null);h.options.__file="PromoCodeDetail.vue";e.a=h.exports},A0ZI:function(t,e,o){},J71L:function(t,e,o){"use strict";o.r(e);var a={name:"CreateForm",components:{PromoCodeDetail:o("2TQq").a}},r=o("KHd+"),n=Object(r.a)(a,function(){var t=this.$createElement;return(this._self._c||t)("promo-code-detail",{attrs:{"is-edit":!1}})},[],!1,null,null,null);n.options.__file="create.vue";e.default=n.exports},MSNs:function(t,e,o){"use strict";o.d(e,"a",function(){return r});var a=o("t3Un");function r(){return Object(a.a)({url:"/qiniu/upload/token",method:"get"})}},RaKr:function(t,e,o){"use strict";o.d(e,"b",function(){return r}),o.d(e,"c",function(){return n}),o.d(e,"a",function(){return s}),o.d(e,"d",function(){return i});var a=o("t3Un");function r(t){return Object(a.a)({url:"/pages/list",method:"get",params:t})}function n(t){return Object(a.a)({url:"/pages/"+t,method:"get"})}function s(t){return Object(a.a)({url:"/pages/",method:"post",data:t})}function i(t,e){return Object(a.a)({url:"/pages/update/"+t,method:"put",data:e})}},Yfch:function(t,e,o){"use strict";function a(t){return["admin"].indexOf(t.trim())>=0}function r(t){return/^(https?|ftp):\/\/([a-zA-Z0-9.-]+(:[a-zA-Z0-9.&%$-]+)*@)*((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(:[0-9]+)*(\/($|[a-zA-Z0-9.,?'\\+&%$#=~_-]+))*$/.test(t)}o.d(e,"a",function(){return a}),o.d(e,"b",function(){return r})},gg54:function(t,e,o){"use strict";o.d(e,"a",function(){return r});var a=o("t3Un");function r(t){return Object(a.a)({url:"/search/user",method:"get",params:{name:t}})}},uwli:function(t,e,o){"use strict";var a=o("A0ZI");o.n(a).a}}]);
//# sourceMappingURL=chunk-695a.be7124a7.js.map