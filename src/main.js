import './wasm_exec'
import Vue from 'vue'
import App from './App.vue'
import router from './router'

// Start wasm
console.log("start main.wasm")

/* eslint no-undef: "off"*/
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
  go.run(result.instance);

  console.log("loaded main.wasm")

  // Go functions and data
  Vue.prototype.$go = {
    // getVal: getVal,
    // printMessage: printMessage,
    // data: data
    // SendCmd: SendCmd
  }
  // Vue.prototype.Send = Send
  // Vue.prototype.SendCmd = SendCmd

// });

  Vue.config.productionTip = false

  new Vue({
    router,
    render: h => h(App)
  }).$mount('#app')

});