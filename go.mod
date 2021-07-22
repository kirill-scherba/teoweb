module github.com/kirill-scherba/teoweb

// replace github.com/kirill-scherba/teowebrtc/teowebrtc_client => ../teowebrtc/teowebrtc_client
// replace github.com/kirill-scherba/teowebrtc/teowebrtc_wasm => ../teowebrtc/teowebrtc_wasm

go 1.16

require (
	github.com/kirill-scherba/teowebrtc/teowebrtc_client v0.0.0-20210722080139-cd6e9b8e8636
	github.com/kirill-scherba/teowebrtc/teowebrtc_server v0.0.0-20210722080139-cd6e9b8e8636
	github.com/kirill-scherba/teowebrtc/teowebrtc_wasm v0.0.0-20210722080139-cd6e9b8e8636
)
