/* Build:

GOOS=js GOARCH=wasm go build -o ../public/main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ..

# install goexec: go get -u github.com/shurcooL/goexec
goexec 'http.ListenAndServe(`:8088`, http.FileServer(http.Dir(`.`)))'

*/
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kirill-scherba/teowebrtc/teowebrtc_client"
	"github.com/kirill-scherba/teowebrtc/teowebrtc_wasm"
)

func main() {
	addr := "localhost:8080"
	name := "teoweb-1"
	server := "server-2"

	// Subscribe to data object
	var subscr = teowebrtc_client.NewSubscribe()
	teowebrtc_wasm.SetFuncs(subscr)

	var id = 0
	var finish = true
	var connected bool

connect:

	// If we was connected than set connected flag to false and wait finish
	// flag will sets to true
	if connected {
		connected = false
		for !finish {
			time.Sleep(100 * time.Microsecond)
		}
	}

	// Connect to teo webrtc application (server)
	// log.Println("Start connection")
	err := teowebrtc_client.Connect(addr, name, server, func(peer string, d *teowebrtc_client.DataChannel) {
		log.Println("Connected to", peer)

		// On open Send messages to created data channel
		d.OnOpen(func() {
			connected = true
			finish = false
			teowebrtc_wasm.SetDataChannel(d)
			for connected {
				id++
				msg := fmt.Sprintf("Hello from %s with id %d!", name, id)
				err := d.Send([]byte(msg))
				if err != nil {
					log.Printf("Send error: %s\n", err)
					continue
				}
				log.Printf("Send: %s", msg)
				time.Sleep(5000 * time.Millisecond)
			}
			log.Println("Finish send messages")
			finish = true
		})

		// On data connection closed
		// This callback does not work in webasm so connect and finish bools
		// flag used to finish previouse connection before reconnect
		d.OnClose(func() {
			log.Println("Connection closed")
			connected = false
		})

		// On message received callback
		d.OnMessage(func(data []byte) {

			// Process subscribers
			if subscr.Process(data) {
				return
			}

			log.Printf("Got: %s\n", data)

			// Call javascript function to set data
			teowebrtc_wasm.SetData(data)
		})
	})
	if err != nil {
		log.Println("connect error:", err)
	}

	// Reconnect after five seconds
	time.Sleep(5 * time.Second)
	goto connect
}
