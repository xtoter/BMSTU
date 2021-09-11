package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"proto"

	"github.com/skorobogatov/input"
)

func get(decoder *json.Decoder) {
	for {
		var resp proto.Response
		if err := decoder.Decode(&resp); err != nil {
			fmt.Printf("error: %v\n", err)
		}

		switch resp.Status {
		case "ok":
			fmt.Printf("ok\n")
		case "failed":
			if resp.Data == nil {
				fmt.Printf("error: data field is absent in response\n")
			} else {
				var errorMsg string
				if err := json.Unmarshal(*resp.Data, &errorMsg); err != nil {
					fmt.Printf("error: malformed data field in response\n")
				} else {
					fmt.Printf("failed: %s\n", errorMsg)
				}
			}
		case "result":
			if resp.Data == nil {
				fmt.Printf("error: data field is absent in response\n")
			} else {
				var event proto.Events
				if err := json.Unmarshal(*resp.Data, &event); err != nil {
					fmt.Printf("error: malformed data field in response\n")
				} else {
					events := event.Eventmessage
					fmt.Println(events)
					for i := 0; i < len(events); i++ {
					}
				}
			}
		default:
			fmt.Printf("error: server reports unknown status %q\n", resp.Status)
		}
	}
}
func interact(conn *net.TCPConn) {
	defer conn.Close()
	encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
	go get(decoder)
	for {
		command := input.Gets()
		switch command {
		case "":
			continue
		case "quit":
			send_request(encoder, "quit", nil)
			return
		case "add":
			var event proto.Events
			fmt.Printf("message = ")
			event.Eventmessage = input.Gets()
			fmt.Printf("time (in sec) = ")
			event.Time = input.Gets()
			send_request(encoder, "add", &event)
		case "check":
			send_request(encoder, "check", nil)
		default:
			fmt.Printf("error: unknown command\n")
			continue
		}

	}
}

func send_request(encoder *json.Encoder, command string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	encoder.Encode(&proto.Request{command, &raw})
}

func main() {
	var addrStr string
	flag.StringVar(&addrStr, "addr", "127.0.0.1:6000", "specify ip address and port")
	flag.Parse()
	if addr, err := net.ResolveTCPAddr("tcp", addrStr); err != nil {
		fmt.Printf("error: %v\n", err)
	} else if conn, err := net.DialTCP("tcp", nil, addr); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		interact(conn)
	}
}
