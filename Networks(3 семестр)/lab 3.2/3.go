package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"

	"github.com/skorobogatov/input"
	"github.com/sparrc/go-ping"
)

type transport struct {
	current *http.Request
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	return http.DefaultTransport.RoundTrip(req)
}

func (t *transport) GotConn(info httptrace.GotConnInfo) {
	fmt.Println(fmt.Sprint(t.current.URL))

	pinger, err := ping.NewPinger(t.current.Host)

	pinger.Count = 1

	if err != nil {

		fmt.Printf("ERROR: %s\n", err.Error())

		return

	}

	pinger.OnRecv = func(pkt *ping.Packet) {

		fmt.Println(pkt.IPAddr)
		//fmt.Println(pkt)

	}

	pinger.Run()
}

func main() {
	t := &transport{}
	fmt.Println("Введите хост")
	var host string
	input.Scanf("%s", &host)
	req, _ := http.NewRequest("GET", "https://"+host, nil)
	trace := &httptrace.ClientTrace{
		GotConn: t.GotConn,
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := &http.Client{Transport: t}
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
	}
}
