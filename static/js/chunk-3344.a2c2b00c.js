(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-3344"],{"8+B/":function(t,e,a){"use strict";var o=a("8wW5");a.n(o).a},"8wW5":function(t,e,a){},FTaA:function(t,e,a){"use strict";a.r(e);var o=a("omC7"),r=a.n(o),n=a("FzGJ"),s=a.n(n);a("KCx4"),a("rLTS"),a("rseV"),a("FTbg"),a("hiQM"),a("DFAU");a("VIyY");var i={name:"JsonEditor",props:["value"],data:function(){return{jsonEditor:!1}},watch:{value:function(t){t!==this.jsonEditor.getValue()&&this.jsonEditor.setValue(r()(this.value,null,2))}},mounted:function(){var t=this;this.jsonEditor=s.a.fromTextArea(this.$refs.textarea,{lineNumbers:!0,mode:"application/json",gutters:["CodeMirror-lint-markers"],theme:"rubyblue",lint:!0}),this.jsonEditor.setValue(r()(this.value,null,2)),this.jsonEditor.on("change",function(e){t.$emit("changed",e.getValue()),t.$emit("input",e.getValue())})},methods:{getValue:function(){return this.jsonEditor.getValue()}}},l=(a("nLDF"),a("ZrdR")),m=Object(l.a)(i,function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"json-editor"},[e("textarea",{ref:"textarea"})])},[],!1,null,"1455b086",null);m.options.__file="index.vue";var u={name:"JsonEditorDemo",components:{JsonEditor:m.exports},data:function(){return{value:JSON.parse('[{"items":[{"market_type":"forexdata","symbol":"XAUUSD"},{"market_type":"forexdata","symbol":"UKOIL"},{"market_type":"forexdata","symbol":"CORN"}],"name":""},{"items":[{"market_type":"forexdata","symbol":"XAUUSD"},{"market_type":"forexdata","symbol":"XAGUSD"},{"market_type":"forexdata","symbol":"AUTD"},{"market_type":"forexdata","symbol":"AGTD"}],"name":"贵金属"},{"items":[{"market_type":"forexdata","symbol":"CORN"},{"market_type":"forexdata","symbol":"WHEAT"},{"market_type":"forexdata","symbol":"SOYBEAN"},{"market_type":"forexdata","symbol":"SUGAR"}],"name":"农产品"},{"items":[{"market_type":"forexdata","symbol":"UKOIL"},{"market_type":"forexdata","symbol":"USOIL"},{"market_type":"forexdata","symbol":"NGAS"}],"name":"能源化工"}]')}}},d=(a("8+B/"),Object(l.a)(u,function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"components-container"},[t._m(0),t._v(" "),a("div",{staticClass:"editor-container"},[a("json-editor",{ref:"jsonEditor",model:{value:t.value,callback:function(e){t.value=e},expression:"value"}})],1)])},[function(){var t=this.$createElement,e=this._self._c||t;return e("code",[this._v("JsonEditor is base on  "),e("a",{attrs:{href:"https://github.com/codemirror/CodeMirror",target:"_blank"}},[this._v("CodeMirrorr")]),this._v(" , lint base on json-lint ")])}],!1,null,"4af1a8e4",null));d.options.__file="jsonEditor.vue";e.default=d.exports},nLDF:function(t,e,a){"use strict";var o=a("srwF");a.n(o).a},srwF:function(t,e,a){}}]);