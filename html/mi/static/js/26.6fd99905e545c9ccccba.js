webpackJsonp([26,32,44],{"+hPY":function(e,t,a){t=e.exports=a("FZ+f")(!1),t.push([e.i,".fade-enter-active[data-v-64e24a1c],.fade-leave-active[data-v-64e24a1c]{-webkit-transition:opacity .5s;transition:opacity .5s}.fade-enter[data-v-64e24a1c],.fade-leave-to[data-v-64e24a1c]{opacity:0}.__upload__[data-v-64e24a1c]{display:inline-block;position:relative;cursor:pointer;border-radius:inherit}.__upload__>[type=file][data-v-64e24a1c]{position:absolute;left:0;top:0;bottom:0;right:0;width:100%;height:100%;opacity:0}.__queue__[data-v-64e24a1c]{width:400px;position:relative}.__item__[data-v-64e24a1c]{position:relative;line-height:30px;height:30px;margin:10px 0;background:#efefef;border-radius:2px;overflow:hidden}.__item__>div[data-v-64e24a1c]{position:absolute;font-size:12px;height:30px}.__name__[data-v-64e24a1c]{z-index:1;left:10px;color:#138d92;right:50px;text-overflow:ellipsis;overflow:hidden}.__name__>svg[data-v-64e24a1c]{display:inline-block;width:12px;height:12px;fill:currentColor;vertical-align:middle;-webkit-transform:translateY(-2px);transform:translateY(-2px)}.__progress__[data-v-64e24a1c]{left:0;top:0;bottom:0;background:#d9eaea;z-index:0}.__percent__[data-v-64e24a1c]{right:50px;top:0;width:40px;height:30px;text-align:center;font-size:10px;color:#138d92}.__cancel__[data-v-64e24a1c]{right:0;top:0;width:30px;height:30px;text-align:center;color:#fff}.__cancel__>svg[data-v-64e24a1c]{width:14px;height:14px;fill:red;-webkit-transform:translateY(4px);transform:translateY(4px)}.__success__[data-v-64e24a1c]{right:0;top:0;width:20px;color:#138d92;-webkit-transform:translateY(2px);transform:translateY(2px)}.__success__>svg[data-v-64e24a1c]{width:12px;height:12px;fill:currentColor}",""])},"3hAG":function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[e.dialogFormVisible?a("div",{staticClass:"modal",staticStyle:{"z-index":"1200"}}):e._e(),e._v(" "),a("el-dialog",{attrs:{modal:!1,title:e.textMap[e.dialogStatus],visible:e.dialogFormVisible,"before-close":e.cancel,size:"small"}},[a("el-form",{ref:"detailForm",staticClass:"small-space",staticStyle:{width:"320px","margin-left":"50px"},attrs:{model:e.detail,rules:e.detailRules,"label-position":"left","label-width":"100px"}},[a("el-form-item",{attrs:{label:"礼品名称",prop:"name"}},[a("el-input",{model:{value:e.detail.name,callback:function(t){e.detail.name=t},expression:"detail.name"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"图片"}},["create"==e.dialogStatus||e.checkPermission(e.permissionConstant.present_u)&&"update"===e.dialogStatus?[a("upload",{attrs:{action:e.image.action,headers:e.uploadHeaders(),disabled:e.image.loading},on:{change:e.image.change,success:e.uploadSuccess,error:e.uploadError}},[a("el-button",{staticStyle:{"margin-bottom":"10px"},attrs:{type:"primary",loading:e.image.loading}},[e._v("上传缩略图")])],1)]:e._e(),e._v(" "),a("img",{staticStyle:{width:"200px",height:"auto",border:"1px solid #bfcbd9"},attrs:{src:e.detail.image,alt:""}})],2),e._v(" "),a("el-form-item",{attrs:{label:"地址",prop:"priority"}},[a("el-input",{model:{value:e.detail.address,callback:function(t){e.detail.address=t},expression:"detail.address"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"权重",prop:"weight"}},[a("el-input",{attrs:{type:"number",min:"0"},model:{value:e.detail.weight,callback:function(t){e.detail.weight=t},expression:"detail.weight"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"需求",prop:"requirement"}},[a("el-input",{attrs:{type:"number",min:"0"},model:{value:e.detail.requirement,callback:function(t){e.detail.requirement=t},expression:"detail.requirement"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"库存",prop:"stock"}},[a("el-input",{attrs:{type:"number",min:"0"},model:{value:e.detail.stock,callback:function(t){e.detail.stock=t},expression:"detail.stock"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"有效时间"}},[a("el-date-picker",{attrs:{clearable:!1,type:"datetime",format:"yyyy-MM-dd HH:mm:ss",placeholder:"选择日期时间"},model:{value:e.detail.expiryDate,callback:function(t){e.detail.expiryDate=t},expression:"detail.expiryDate"}})],1),e._v(" "),a("el-form-item",{attrs:{label:"状态"}},[a("el-select",{staticClass:"filter-item",attrs:{placeholder:"状态"},model:{value:e.detail.hide,callback:function(t){e.detail.hide=t},expression:"detail.hide"}},e._l(e.statusOptions,function(e){return a("el-option",{key:e.key,attrs:{label:e.label,value:e.key}})}))],1)],1),e._v(" "),a("div",{staticClass:"dialog-footer",slot:"footer"},[a("el-button",{on:{click:e.cancel}},[e._v("取 消")]),e._v(" "),"create"==e.dialogStatus?a("el-button",{attrs:{type:"primary"},on:{click:e.create}},[e._v("确 定")]):e._e(),e._v(" "),"update"==e.dialogStatus&&e.checkPermission(e.permissionConstant.present_u)?[a("el-button",{attrs:{type:"primary"},on:{click:e.update}},[e._v("确 定")])]:e._e()],2)],1)],1)},staticRenderFns:[]}},"7oMk":function(e,t,a){var i=a("VU/8")(a("jMdz"),a("3hAG"),null,null,null);e.exports=i.exports},DmoS:function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("div",{staticClass:"__upload__"},[e._t("default"),e._v(" "),a("input",{attrs:{type:"file",disabled:e.disabled,accept:e.accepts,multiple:e.multiple},on:{change:function(t){e.change(t)}}})],2),e._v(" "),e.queue&&e.safeFiles.length>0?a("div",{staticClass:"__queue__"},[a("transition-group",{attrs:{name:"fade",mode:"out-in"}},e._l(e.safeFiles,function(t,i){return t.done?e._e():a("div",{key:i,staticClass:"__item__"},[a("div",{staticClass:"__progress__",style:{width:t.progress+"%"}}),e._v(" "),a("div",{staticClass:"__name__"},[a("svg",{attrs:{viewBox:"0 0 1024 1024"}},[a("path",{attrs:{d:"M636.974339 66.552765 139.739625 66.552765l0 890.583384 742.153844 0L881.893468 304.049854 636.974339 66.552765 636.974339 66.552765zM591.950913 348.826664 591.950913 125.932154 822.506916 348.826664 591.950913 348.826664 591.950913 348.826664zM591.950913 348.826664"}})]),e._v("\n          "+e._s(t.name)+"\n        ")]),e._v(" "),t.progress>0&&t.progress<100?a("div",{staticClass:"__percent__"},[e._v(e._s(t.progress)+"%")]):e._e(),e._v(" "),t.progress<100?a("div",{staticClass:"__cancel__",on:{click:function(t){e.cancel(i)}}},[a("svg",{attrs:{viewBox:"0 0 1024 1024"}},[a("path",{attrs:{d:"M512.325411 0c-282.578844 0-511.653099 229.075279-511.653099 511.653099 0 282.578844 229.074256 511.653099 511.653099 511.653099s511.653099-229.074256 511.653099-511.653099C1023.978511 229.075279 794.904255 0 512.325411 0zM726.690664 761.454422c-4.821819 4.146437-10.754948 6.183839-16.663518 6.183839-7.194866 0-14.352893-3.022847-19.412119-8.906857L509.457084 548.069497 329.953827 760.043283c-5.059226 5.983271-12.272511 9.05626-19.535939 9.05626-5.838985 0-11.710716-1.986237-16.520255-6.058996-10.780531-9.131985-12.123109-25.269523-2.991124-36.051077l184.773284-218.201627L302.096363 306.936601c-9.212826-10.717086-7.995091-26.868951 2.716878-36.080753 10.711969-9.193383 26.881231-8.000208 36.080753 2.716878L509.160325 469.24729l166.813237-196.993606c9.118682-10.799974 25.296129-12.123109 36.051077-2.991124 10.79281 9.131985 12.130272 25.276686 2.998287 36.056194L542.939663 508.529969l186.472995 216.84984C738.619344 736.083591 737.40775 752.23648 726.690664 761.454422z",fill:"#e84122"}})])]):e._e(),e._v(" "),100==t.progress?a("div",{staticClass:"__success__"},[a("svg",{attrs:{viewBox:"0 0 1024 1024"}},[a("path",{attrs:{d:"M512 0C229.239467 0 0 229.239467 0 512s229.239467 512 512 512 512-229.239467 512-512S794.760533 0 512 0z m-68.778667 699.733333l-0.170666-0.136533-0.1536 0.136533L238.933333 490.376533l35.328-34.594133 168.789334 157.934933L752.298667 324.266667 785.066667 359.287467 443.221333 699.733333z"}})])]):e._e()])}))],1):e._e()])},staticRenderFns:[]}},E4LH:function(e,t,a){"use strict";function i(e){return/^(https?|ftp):\/\/([a-zA-Z0-9.-]+(:[a-zA-Z0-9.&%$-]+)*@)*((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(:[0-9]+)*(\/($|[a-zA-Z0-9.,?'\\+&%$#=~_-]+))*$/.test(e)}function n(e,t,a){""===t||/^(((13[0-9]{1})|(15[0-9]{1})|(18[0-9]{1}))+\d{8})$/.test(t)?a():a(new Error("手机号码不合法"))}function s(e){return function(t,a,i){a&&!/^[\d]+$/.test(a)?i(new Error(e)):i()}}function r(e){return function(t,a,i){a&&!/^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*))$/.test(a)?i(new Error(e)):i()}}function l(e,t,a){t>999?a("优先权重最大999"):a()}t.e=i,t.a=n,t.c=s,t.d=r,t.b=l},GMHh:function(e,t){e.exports={render:function(){var e=this,t=e.$createElement;return(e._self._c||t)("Records",{attrs:{containerClass:"app-container"}})},staticRenderFns:[]}},K7IE:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var i=a("Dd8w"),n=a.n(i),s=a("woOf"),r=a.n(s),l=a("NYxO"),o=a("7oMk"),c=a.n(o),d={address:"",expiryDate:"",hide:"",id:"",image:"",name:"",requirement:0,stock:0,weight:0};t.default={components:{PresentDetail:c.a},props:{containerClass:{type:String,default:""},isMain:{type:Boolean,default:!0},presents:{type:Array,default:function(){return[]}}},name:"crp_present",data:function(){return{selections:[],listLoading:!0,textMap:{update:"编辑",create:"创建"},listQuery:{targetPage:1,pageSize:10,keyword:void 0},temp:r()({},d),statusOptions:[{label:"显示",key:"false"},{label:"隐藏",key:"true"}],dialogFormVisible:!1,dialogStatus:"",tableKey:0}},computed:n()({},a.i(l.a)(["present"])),created:function(){this.getList()},filters:{statusFilter:function(e){return!1===e?"显示":"隐藏"}},methods:{getList:function(){var e=this;this.listLoading=!0,this.$store.dispatch("GetAllPresent",this.listQuery).then(function(){e.listLoading=!1},function(){})},handleFilter:function(){this.getList()},handleSizeChange:function(e){this.listQuery.pageSize=e,this.getList()},handleCurrentChange:function(e){this.listQuery.targetPage=e,this.getList()},handleBatchDelete:function(){if(0===this.selections.length)return void this.$message({message:"请选择要删除的奖品",type:"warning"});var e=this.selections.map(function(e){return e.id});this.delete(e,"确认批量删除奖品？")},handleModifyStatus:function(e){this.delete([e.id],"确认删除广告："+e.name+"？")},delete:function(e,t){var a=this;this.$confirm(t).then(function(){a.$store.dispatch("DelPresent",{ids:e}).then(function(){a.$message({message:"操作成功",type:"success"}),a.getList()})},function(){})},handleCreate:function(){this.dialogStatus="create",this.dialogFormVisible=!0},handleUpdate:function(e){var t=this;this.$store.dispatch("GetPresentDetail",e.id).then(function(e){t.temp=r()({},e),t.temp.hide=String(t.temp.hide)}),this.dialogStatus="update",this.dialogFormVisible=!0},submit:function(){this.dialogFormVisible=!1,this.getList(),this.temp=r()({},d)},cancel:function(){this.dialogFormVisible=!1,this.temp=r()({},d)},check:function(e){this.$emit("check",e)},cancelCheck:function(e){this.$emit("cancel-check",e)},has:function(e){return this.presents.some(function(t){return t.id===e})},handleSelectionChange:function(e){this.selections=e}}}},SwZ3:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var i=a("kjTw"),n=a.n(i);t.default={components:{Records:n.a}}},TAej:function(e,t,a){function i(e){a("vB4V")}var n=a("VU/8")(a("jsWf"),a("DmoS"),i,"data-v-64e24a1c",null);e.exports=n.exports},cHx6:function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{class:[e.containerClass]},[a("div",{staticClass:"filter-container"},[a("el-input",{staticClass:"filter-item",staticStyle:{width:"200px"},attrs:{placeholder:"礼品名称"},nativeOn:{keyup:function(t){if(!("button"in t)&&e._k(t.keyCode,"enter",13))return null;e.handleFilter(t)}},model:{value:e.listQuery.keyword,callback:function(t){e.listQuery.keyword=t},expression:"listQuery.keyword"}}),e._v(" "),a("el-button",{directives:[{name:"waves",rawName:"v-waves"}],staticClass:"filter-item",attrs:{type:"primary",icon:"search"},on:{click:e.handleFilter}},[e._v("搜索")]),e._v(" "),e.isMain?[e.checkPermission(e.permissionConstant.present_c)?a("el-button",{staticClass:"filter-item",staticStyle:{"margin-left":"10px"},attrs:{type:"primary",icon:"edit"},on:{click:e.handleCreate}},[e._v("添加")]):e._e(),e._v(" "),e.checkPermission(e.permissionConstant.present_d)?a("el-button",{staticClass:"filter-item",staticStyle:{"margin-left":"10px"},attrs:{type:"danger",icon:"edit"},on:{click:e.handleBatchDelete}},[e._v("批量删除")]):e._e()]:e._e()],2),e._v(" "),a("el-table",{directives:[{name:"loading",rawName:"v-loading.body",value:e.listLoading,expression:"listLoading",modifiers:{body:!0}}],key:e.tableKey,staticStyle:{width:"100%"},attrs:{data:e.present.records,border:"",fit:"","highlight-current-row":""},on:{"selection-change":e.handleSelectionChange}},[e.isMain?a("el-table-column",{attrs:{type:"selection",width:"55"}}):e._e(),e._v(" "),a("el-table-column",{attrs:{align:"center",label:"名称",width:"120"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",{class:{"link-type":e.isMain},on:{click:function(a){e.handleUpdate(t.row)}}},[e._v(e._s(t.row.name))])]}}])}),e._v(" "),a("el-table-column",{attrs:{width:"160",align:"center",label:"图片"},scopedSlots:e._u([{key:"default",fn:function(e){return[a("img",{staticStyle:{width:"120px",height:"auto","padding-top":"5px"},attrs:{src:e.row.image,alt:""}})]}}])}),e._v(" "),a("el-table-column",{attrs:{align:"center",label:"地址"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.address))])]}}])}),e._v(" "),a("el-table-column",{attrs:{align:"center",label:"库存",width:"80"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.stock))])]}}])}),e._v(" "),a("el-table-column",{attrs:{align:"center",label:"需求",width:"80"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.requirement))])]}}])}),e._v(" "),a("el-table-column",{attrs:{align:"center",label:"抽奖权重",width:"120"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.weight))])]}}])}),e._v(" "),a("el-table-column",{attrs:{align:"center",label:"有效时间",width:"120"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.expiryDate))])]}}])}),e._v(" "),e.isMain?[a("el-table-column",{attrs:{"class-name":"status-col",label:"状态",width:"60"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:t.row.hide?"danger":"primary"}},[e._v(e._s(e._f("statusFilter")(t.row.hide)))])]}}])}),e._v(" "),e.checkPermission(e.permissionConstant.present_d)?a("el-table-column",{attrs:{align:"center",label:"操作",width:"150"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-button",{attrs:{size:"small",type:"danger"},on:{click:function(a){e.handleModifyStatus(t.row,!0)}}},[e._v("删除")])]}}])}):e._e()]:[a("el-table-column",{attrs:{align:"center",label:"操作",width:"150"},scopedSlots:e._u([{key:"default",fn:function(t){return[e.has(t.row.id)?a("el-button",{attrs:{size:"small",type:"primary"},on:{click:function(a){e.cancelCheck(t.row)}}},[e._v("取消选中")]):a("el-button",{attrs:{size:"small"},on:{click:function(a){e.check(t.row)}}},[e._v("选中")])]}}])})]],2),e._v(" "),a("div",{directives:[{name:"show",rawName:"v-show",value:!e.listLoading,expression:"!listLoading"}],staticClass:"pagination-container"},[a("el-pagination",{attrs:{"current-page":e.listQuery.currentPage,"page-sizes":[10,20,30,50],"page-size":e.listQuery.pageSize,layout:"total, sizes, prev, pager, next, jumper",total:e.present.pageInfo.totalRow},on:{"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange,"update:currentPage":function(t){e.listQuery.currentPage=t}}})],1),e._v(" "),a("Present-Detail",{attrs:{title:e.textMap[e.dialogStatus],visible:e.dialogFormVisible,"before-close":e.cancel,"dialog-status":e.dialogStatus,detail:e.temp,"status-options":e.statusOptions,"dialog-form-visible":e.dialogFormVisible},on:{submit:function(t){e.submit()},cancel:function(t){e.cancel()}}})],1)},staticRenderFns:[]}},jMdz:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var i=a("woOf"),n=a.n(i),s=a("//Fk"),r=a.n(s),l=a("TAej"),o=a.n(l),c=a("oGH5"),d=a("E4LH");t.default={components:{Upload:o.a},mixins:[new c.a(function(e){1===e.status&&(this.detail.image=e.result),this.image.loading=!1},function(){this.image.loading=!1})],props:{dialogStatus:{type:String,default:"create"},statusOptions:{type:Array,default:function(){return[]}},dialogFormVisible:{type:Boolean,default:!1},detail:{type:Object,default:function(){return{id:"",name:"",hide:"",image:"",address:"",expiryDate:"",requirement:0,stock:0,weight:0}}}},data:function(){var e=this;return{isMain:!1,image:{action:"http://djq.tunnel.qydev.com/mi/presentAction/uploadImage",loading:!1,change:function(){e.image.loading=!0}},textMap:{update:"编辑",create:"创建"},detailRules:{name:[{required:!0,min:3,max:32,message:"奖品名称长度3到32位",trigger:"blur"}],requirement:[{validator:d.c("需求只能为数字"),trigger:"blur"}],stock:[{validator:d.c("库存只能为数字"),trigger:"blur"}],weight:[{validator:d.c("权重只能为数字"),trigger:"blur"}]}}},methods:{validate:function(){var e=this;return new r.a(function(t,a){if(e.image.loading)return e.$message.warning("正在上传图片缩略图，请稍后提交"),void a();e.$refs.detailForm.validate(function(e){e?t():a()})})},create:function(){var e=this,t=this;t.validate().then(function(){var a=n()({},t.detail);delete a.id,t.$store.dispatch("CreatePresent",a).then(function(){e.$notify({title:"成功",message:"创建成功",type:"success",duration:2e3}),e.$emit("submit")},function(){})},function(){})},update:function(){var e=this,t=this;t.validate().then(function(){var a=n()({},t.detail);t.$store.dispatch("UpdatePresentDetail",a).then(function(){t.$notify({title:"成功",message:"更新成功",type:"success",duration:2e3}),e.$emit("submit")},function(){})},function(){})},cancel:function(){this.$emit("cancel")}},watch:{dialogFormVisible:function(){this.$refs.detailForm&&this.$refs.detailForm.resetFields()}}}},jsWf:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.default={props:{action:{type:String},disabled:{type:Boolean,default:!1},headers:{type:Object,default:function(){return{}}},data:{type:Object,default:function(){return{}}},dataType:{type:String,default:"json"},auto:{type:Boolean,default:!0},name:{type:String,default:"theFile"},limit:{type:Number,default:20},accepts:{type:Array,default:function(){return["image/jpeg","image/png"]}},multiple:{type:Boolean,default:!1},chunked:{type:Boolean,default:!1},queue:{type:Boolean,default:!1}},data:function(){return{files:[],safeFiles:[],unsafeFiles:[],input:null}},methods:{change:function(e){this.input=e.target,this.files=[],this.unsafeFiles=[],this.files=e.target.files||e.dataTransfer.files;for(var t,a=0;t=this.files[a];a++)void 0==this.limit?this.addFile(t):t.size<=1024*this.limit*1024?this.addFile(t):this.unsafeFiles.push(t);this.$emit("change",this.safeFiles),this.auto&&(this.upload(),this.safeFiles=[])},addFile:function(e){for(var t,a=!1,i=0;t=this.safeFiles[i];i++)e.name==t.name&&(a=!0);a||(e.progress=0,this.safeFiles.push(e))},upload:function(){var e=this;if(this.unsafeFiles.length>0)return this.$emit("error","limit",this.unsafeFiles),!1;if(0==this.safeFiles.length)return this.$emit("error","empty"),!1;if(void 0==this.action)return this.$emit("error","action"),!1;this.input.value="";for(var t,a=0;t=this.safeFiles[a];a++)!function(t){t.size>5242880&&e.chunked?e.uploadChunked(t):e.uploadNormal(t)}(t)},uploadNormal:function(e){var t=this;if(e.xhr=new XMLHttpRequest,e.xhr.upload){var a=new FormData;if(a.append(this.name,e),void 0!=this.data)for(var i in this.data)a.append(i,this.data[i]);if(e.xhr.upload.progress=function(a){e.progress=Math.floor(a.loaded/a.total*100),100==e.progress&&window.setTimeout(function(){e.done=!0,t.$forceUpdate()},1e3),t.$emit("progress",e.progress),t.$forceUpdate()},e.xhr.onreadystatechange=function(a){if(4==e.xhr.readyState)if(200==e.xhr.status){var i=e.xhr.responseText;"json"==t.dataType&&(i=JSON.parse(i)),t.$emit("success",i)}else t.$emit("error","server",t.file)},e.xhr.open("POST",this.action,!0),void 0!=this.headers)for(var n in this.headers)e.xhr.setRequestHeader(n,this.headers[n]);e.xhr.withCredentials=!1,e.xhr.send(a)}},uploadChunked:function(e){var t=this,a=Math.random().toString(36).substr(2),i=Math.ceil(e.size/1048576),n=0,s=0;!function r(){var l=1048576*n,o=l+1048576;if(e.xhr=new XMLHttpRequest,e.xhr.upload){var c=new FormData;if(c.append("chunked","true"),c.append("chunkedIndex",n),c.append("chunkedID",a),c.append("chunkedTotal",i),c.append("chunkedData",e.slice(l,o)),void 0!=t.data)for(var d in t.data)c.append(d,t.data[d]);if(e.xhr.upload.addEventListener("progress",function(a){s+=a.loaded,e.progress=Math.floor(s/e.size*100),100==e.progress&&window.setTimeout(function(){e.done=!0,t.$forceUpdate()},1e3),t.$emit("progress",e.progress),t.$forceUpdate()},!1),e.xhr.onreadystatechange=function(a){if(4==e.xhr.readyState)if(200==e.xhr.status){var s=e.xhr.responseText;"json"==t.dataType&&(s=JSON.parse(s)),n++,n<i&&r(),t.$emit("success",s)}else t.$emit("error","server",t.file)},e.xhr.open("POST",t.action,!0),void 0!=t.headers)for(var u in t.headers)e.xhr.setRequestHeader(u,t.headers[u]);e.xhr.withCredentials=!1,e.xhr.send(c)}}()},submit:function(){this.upload()},cancel:function(e){this.safeFiles[e].xhr&&this.safeFiles[e].xhr.abort(),this.safeFiles.splice(e,1),this.input.value=""}}}},kjTw:function(e,t,a){var i=a("VU/8")(a("K7IE"),a("cHx6"),null,null,null);e.exports=i.exports},"lJ+M":function(e,t,a){var i=a("VU/8")(a("SwZ3"),a("GMHh"),null,null,null);e.exports=i.exports},oGH5:function(e,t,a){"use strict";var i=a("Zrlr"),n=a.n(i),s=function e(t,a){return n()(this,e),{methods:{uploadHeaders:function(){return{}},uploadSuccess:function(e){var a=this;2===e.status?a.$messageBox.confirm("你已被登出，可以取消继续留在该页面，或者重新登录","确定登出",{confirmButtonText:"重新登录",cancelButtonText:"取消",type:"warning"}).then(function(){a.$store.dispatch("FedLogOut").then(function(){location.reload()})}):1!==e.status&&a.$message({message:e.message||e.msg,type:"error",duration:2e3}),t.call(a,e)},uploadError:function(e){var t=this;"limit"===e&&t.$message({message:"图片大小不能超过20M",type:"error"}),"empty"===e&&t.$message({message:"请选择图片",type:"warning"}),"server"===e&&t.$message({message:"服务器错误，请联系管理员处理",type:"error"}),a.call(t,e)}}}};t.a=s},vB4V:function(e,t,a){var i=a("+hPY");"string"==typeof i&&(i=[[e.i,i,""]]),i.locals&&(e.exports=i.locals);a("rjj0")("14bc98a0",i,!0)}});