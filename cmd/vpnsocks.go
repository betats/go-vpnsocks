package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/armon/go-socks5"
	"github.com/mysteriumnetwork/go-openvpn/openvpn3"
)

type callbacks interface {
	openvpn3.Logger
	openvpn3.EventConsumer
	openvpn3.StatsConsumer
}

type loggingCallbacks struct {
}

func (lc *loggingCallbacks) Log(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Println("Openvpn log >>", line)
	}
}

func (lc *loggingCallbacks) OnEvent(event openvpn3.Event) {
	fmt.Printf("Openvpn event >> %+v\n", event)
}

func (lc *loggingCallbacks) OnStats(stats openvpn3.Statistics) {
	fmt.Printf("Openvpn stats >> %+v\n", stats)
}

var _ callbacks = &loggingCallbacks{}

// StdoutLogger represents the stdout logger callback
type StdoutLogger func(text string)

// Log logs the given string to stdout logger
func (lc StdoutLogger) Log(text string) {
	lc(text)
}

func main() {
	username := os.Getenv("OPENVPN_USERNAME")
	password := os.Getenv("OPENVPN_PASSWORD")
	if username == "" || password == "" {
		panic("OPENVPN_USERNAME or OPENVPN_PASSWORD is empty")
	}
	// openvpn start
	// Create an instance of the openvpn struct
	var logger StdoutLogger = func(text string) {
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			fmt.Println("Library check >>", line)
		}
	}

	openvpn3.SelfCheck(logger)

	bytes, err := ioutil.ReadFile("/etc/openvpn/profile.ovpn")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	config := openvpn3.NewConfig(string(bytes))

	session := openvpn3.NewSession(config, openvpn3.UserCredentials{Username: username, Password: password}, &loggingCallbacks{})
	session.Start()

	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy port 1080
	if err := server.ListenAndServe("tcp", "0.0.0.0:1080"); err != nil {
		panic(err)
	}

	// openvpn session
	err = session.Wait()
	if err != nil {
		fmt.Println("Openvpn3 error: ", err)
	} else {
		fmt.Println("Graceful exit")
	}

}
