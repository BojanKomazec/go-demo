package httpserverdemo

import (
	"fmt"
	"log"
	"net/http"
)

// If httpServerDir has value "." server will serve files from the current directory (where this application is run from).
// If httpServerDir has value "/dirXYZ" server will serve files from the root directory "dirXYZ"
// Don't use "~/..." to specify home directory, use "/home/user/...".
//
// It might be necessary to open port 8080 on the machine running HTTP server.
// For Ubuntu, please refer to this article: https://www.bojankomazec.com/2019/12/how-to-open-ports-on-ubuntu.html
// It is not necessary for HTTP Server to run as root.
//
// To access HTTP Server from another machine, open on it a browser and type: IP:8080 where IP is the IP address of the machine
// which is running HTTP Server. Use ifconfig on HTTP Server machine to find out that IP. Both machines have to be on the same LAN.
func runServer(httpServerDir string) {
	fmt.Printf("Serving files in the directory %s on port 8080\n", httpServerDir)
	http.Handle("/", http.FileServer(http.Dir(httpServerDir)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// ShowDemo func
func ShowDemo(httpServerDir string) {
	runServer(httpServerDir)
}
