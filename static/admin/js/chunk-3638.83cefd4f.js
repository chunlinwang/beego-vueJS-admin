(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-3638"],{GRjg:function(t,e,n){"use strict";var o=n("nDc9");n.n(o).a},Y5bG:function(t,e,n){"use strict";n.d(e,"a",function(){return i}),Math.easeInOutQuad=function(t,e,n,o){return(t/=o/2)<1?n/2*t*t+e:-n/2*(--t*(t-2)-1)+e};var o=window.requestAnimationFrame||window.webkitRequestAnimationFrame||window.mozRequestAnimationFrame||function(t){window.setTimeout(t,1e3/60)};function i(t,e,n){var i=document.documentElement.scrollTop||document.body.parentNode.scrollTop||document.body.scrollTop,a=t-i,r=0;e=void 0===e?500:e;!function t(){r+=20,function(t){document.documentElement.scrollTop=t,document.body.parentNode.scrollTop=t,document.body.scrollTop=t}(Math.easeInOutQuad(r,i,a,e)),r<e?o(t):n&&"function"==typeof n&&n()}()}},nDc9:function(t,e,n){},uzFz:function(t,e,n){"use strict";n.r(e);var o=n("xMja"),i={name:"PageList",components:{Pagination:n("Mz3J").a},filters:{statusFilter:function(t){return{published:"success",draft:"info",deleted:"danger"}[t]}},data:function(){return{list:null,total:0,listLoading:!0,listQuery:{page:1,limit:20}}},created:function(){this.getList()},methods:{getList:function(){var t=this;this.listLoading=!0,Object(o.b)(this.listQuery).then(function(e){t.list=e.data.items,t.total=e.data.total,t.listLoading=!1})},handleSizeChange:function(t){this.listQuery.limit=t,this.getList()},handleCurrentChange:function(t){this.listQuery.page=t,this.getList()}}},a=(n("GRjg"),n("KHd+")),r=Object(a.a)(i,function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-container"},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],staticStyle:{width:"100%"},attrs:{data:t.list,border:"",fit:"","highlight-current-row":""}},[n("el-table-column",{attrs:{label:t.$t("product.id"),align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",[t._v(t._s(e.row._id))])]}}])}),t._v(" "),n("el-table-column",{attrs:{label:t.$t("product.number"),align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",[t._v(t._s(e.row.number))])]}}])}),t._v(" "),n("el-table-column",{attrs:{label:t.$t("product.price"),align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",[t._v(t._s(e.row.price)+" €")])]}}])}),t._v(" "),n("el-table-column",{attrs:{label:t.$t("product.onSale"),align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",[t._v(t._s(e.row.on_sale))])]}}])}),t._v(" "),n("el-table-column",{attrs:{label:t.$t("product.price_on_sale"),align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",[t._v(t._s(e.row.price_on_sale)+" €")])]}}])}),t._v(" "),n("el-table-column",{attrs:{label:t.$t("product.actions"),align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("router-link",{attrs:{to:"/product/edit/"+e.row._id}},[n("el-button",{attrs:{type:"primary",size:"small",icon:"el-icon-edit"}},[t._v(t._s(t.$t("edit")))])],1)]}}])})],1),t._v(" "),n("pagination",{directives:[{name:"show",rawName:"v-show",value:t.total>0,expression:"total>0"}],attrs:{total:t.total,page:t.listQuery.page,limit:t.listQuery.limit},on:{"update:page":function(e){t.$set(t.listQuery,"page",e)},"update:limit":function(e){t.$set(t.listQuery,"limit",e)},pagination:t.getList}})],1)},[],!1,null,"3be74acd",null);r.options.__file="list.vue";e.default=r.exports},xMja:function(t,e,n){"use strict";n.d(e,"b",function(){return i}),n.d(e,"c",function(){return a}),n.d(e,"a",function(){return r}),n.d(e,"e",function(){return u}),n.d(e,"d",function(){return l});var o=n("t3Un");function i(t){return Object(o.a)({url:"/products/list",method:"get",params:t})}function a(t){return Object(o.a)({url:"/products/"+t,method:"get"})}function r(t){return Object(o.a)({url:"/products/",method:"post",data:t})}function u(t,e){return Object(o.a)({url:"/products/update/"+t,method:"put",data:e})}function l(t){return Object(o.a)({url:"/products",method:"get",params:t})}}}]);
//# sourceMappingURL=chunk-3638.83cefd4f.js.map