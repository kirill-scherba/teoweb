# Teoweb - teonet web application with webrtc data channel and webasm and vuejs web client  

Teonet webrtc sample application, go server, wasm and vuejs client. Client connect to
server and transfer webrtc data packages.
## Project setup
```
npm install
```

## Run signal server

https://github.com/kirill-scherba/teowebrtc/teowebrtc_signal
```
go run ./cmd/teowebrtc_signal
```

## Run webrtc server
```
cd server
go run .
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
