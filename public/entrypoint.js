
// // Monitor data
// var data = { dcAnswer: "Wait loading...", num: 0 }
// function setData (d) {
//   data = JSON.parse(d)
//   console.log("setData", d)
// }
// setData('{"num":0}')
// console.log(data)

// // Start wasm
// console.log("start main.wasm")
// const go = new Go();
// WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
//     go.run(result.instance);
//     printMessage("Hello Go from JS!");
//     console.log("loaded main.wasm")
// });

// Global data
// var setData;
// var data;
// var SendCmd;
// var SetCallback;

function SendCmdPromise(cmd, msg) {
    return new Promise((resolve, reject) => {
        SendCmd(cmd, msg, (err, message) => {
            if (err) {
                reject(err);
                return;
            }
            resolve(message);
        });
    });
}