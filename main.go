package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

//go:embed index.html index.js
var indexHTML embed.FS

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	mux := http.NewServeMux()

	// upgrade /socket connections from websocket
	mux.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		// upgrade to websocket or bail out
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Unable to upgrade connection: %s", err)
			w.WriteHeader(400)
			return
		}
		defer conn.Close()

		log.Printf("[%s]: upgraded to websocket", r.RemoteAddr)
		for {
			// read the next message from the client
			typ, bytes, err := conn.ReadMessage()
			if err != nil {
				// client is leaving "normally"
				if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
					log.Printf("[%s]: closing websocket", r.RemoteAddr)
					break
				}

				// something else happened
				log.Printf("[%s]: error reading: %s", r.RemoteAddr, err)
				break
			}

			// echo whatever you received back
			err = conn.WriteMessage(typ, bytes)
			if err != nil {
				log.Printf("[%s]: error writing: %s", r.RemoteAddr, err)
				break
			}
		}
	})

	// serve static files on "/"
	mux.Handle("/", http.FileServer(http.FS(indexHTML)))

	// and serve
	log.Printf("listening on %s", flagBind)
	http.ListenAndServe(flagBind, mux)
}

var flagBind string = "localhost:8080"

func init() {
	var legalFlag bool = false
	flag.BoolVar(&legalFlag, "legal", legalFlag, "Display legal notices and exit")
	defer func() {
		if legalFlag {
			fmt.Print(LegalText())
			os.Exit(0)
		}
	}()

	defer flag.Parse()

	flag.StringVar(&flagBind, "bind", flagBind, "address and port to bind to")
}
