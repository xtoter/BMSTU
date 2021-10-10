package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"proto"
	"strconv"
	"time"

	log "github.com/mgutz/logxi/v1"
)

type Client struct {
	logger log.Logger     // Объект для печати логов
	conn   *net.TCPConn   // Объект TCP-соединения
	enc    *json.Encoder  // Объект для кодирования и отправки сообщений
	number int            // Количество напоминаний
	events []proto.Events // События (Событие,время)
}

func NewClient(conn *net.TCPConn) *Client {
	return &Client{
		logger: log.New(fmt.Sprintf("client %s", conn.RemoteAddr().String())),
		conn:   conn,
		enc:    json.NewEncoder(conn),
		number: 0,
		events: []proto.Events{},
	}
}
func (client *Client) serve() {
	defer client.conn.Close()
	decoder := json.NewDecoder(client.conn)
	for {
		var req proto.Request
		if err := decoder.Decode(&req); err != nil {
			fmt.Println(err)
			client.logger.Error("cannot decode message", "reason", fmt.Sprint(err))
			break
		} else {
			client.logger.Info("received command", "command", req.Command)
			if client.handleRequest(&req) {
				client.logger.Info("shutting down connection")
				break
			}
		}
	}
}
func check(client *Client, i int) {
	timeelem, _ := strconv.Atoi(client.events[i].Time)
	for int(time.Now().Unix()) != timeelem {
	}
	client.logger.Info("sent reminder")
	client.respond("result", &proto.Events{
		Eventmessage: fmt.Sprint(client.events[i].Eventmessage),
		Time:         fmt.Sprint(time.Now().Unix()),
	})
}

func (client *Client) handleRequest(req *proto.Request) bool {
	switch req.Command {
	case "quit":
		client.respond("ok", nil)
		return true
	case "add":
		errorMsg := ""
		if req.Data == nil {
			errorMsg = "data field is absent"
		} else {
			var event proto.Events
			if err := json.Unmarshal(*req.Data, &event); err != nil {
				errorMsg = "data field"
			} else {
				client.logger.Info("performing addition", "value", "msg : "+event.Eventmessage+", time: "+event.Time+" sec")
				timeevent, cond := strconv.Atoi(event.Time)
				if cond == nil {
					event.Time = fmt.Sprint(timeevent + int(time.Now().Unix()))
					client.events = append(client.events, event)
					client.number++
					go check(client, client.number-1)
				} else {
					errorMsg = "invalid expression"
				}

			}
		}
		if errorMsg == "" {
			client.respond("ok", nil)
		} else {
			client.logger.Error("addition failed", "reason", errorMsg)
			client.respond("failed", errorMsg)
		}

	case "check":
		if client.number == 0 {
			client.logger.Info("error, no events")
			client.respond("failed", "404")
		} else {
			client.logger.Info("send all events")
			client.respond("result", &proto.Events{
				Eventmessage: fmt.Sprint(client.events),
				Time:         fmt.Sprint(time.Now().Unix()),
			})
		}
	default:
		client.logger.Error("unknown command")
		client.respond("failed", "unknown command")
	}
	return false
}

func (client *Client) respond(status string, data interface{}) {
	var raw json.RawMessage
	raw, _ = json.Marshal(data)
	client.enc.Encode(&proto.Response{status, &raw})
}

func main() {
	// Работа с командной строкой, в которой может указываться необязательный ключ -addr.
	var addrStr string
	flag.StringVar(&addrStr, "addr", "127.0.0.1:6000", "specify ip address and port")
	flag.Parse()

	if addr, err := net.ResolveTCPAddr("tcp", addrStr); err != nil {
		log.Error("address resolution failed", "address", addrStr)
	} else {
		log.Info("resolved TCP address", "address", addr.String())

		if listener, err := net.ListenTCP("tcp", addr); err != nil {
			log.Error("listening failed", "reason", err)
		} else {
			for {
				if conn, err := listener.AcceptTCP(); err != nil {
					log.Error("cannot accept connection", "reason", err)
				} else {
					log.Info("accepted connection", "address", conn.RemoteAddr().String())
					go NewClient(conn).serve()
				}
			}
		}
	}
}
