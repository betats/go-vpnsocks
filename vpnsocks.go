package main

import "github.com/armon/go-socks5" 

func main() {
	// openvpn start
	// Create an instance of the openvpn struct
	p := openvpn.NewStaticKeyClient("localhost", "pre-shared.key")

	// Start the openvpn process. Note that this method do not block so the program will continue at once.
	p.Start()

	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "127.0.0.1:8000"); err != nil {
		panic(err)
	}
}
