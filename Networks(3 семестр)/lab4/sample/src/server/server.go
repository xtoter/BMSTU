package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"

	log "github.com/mgutz/logxi/v1"
)

type Data struct {
	Numpackage    int
	Globalpackage int
	Point1x       int
	Point1y       int
	Point2x       int
	Point2y       int
	Client        string
}

func parserequest(str string) (packagenum, globalpackage, point1x, point1y, point2x, point2y, numclient int) {
	var m Data
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		log.Error("creating listening connection", "error", err)
		return
	}

	res, _ := strconv.Atoi(m.Client)

	return m.Numpackage, m.Globalpackage, m.Point1x, m.Point1y, m.Point2x, m.Point2y, res
}
func calculation(point1x, point1y, point2x, point2y int) float64 {
	return math.Pi * (math.Pow(float64(point1x-point2x), 2) + math.Pow(float64(point1y-point2y), 2))
}
func main() {
	client := 0
	var global []int
	result := float64(0)
	results := map[int]float64{}
	var (
		serverAddrStr string
		helpFlag      bool
	)
	flag.StringVar(&serverAddrStr, "addr", "127.0.0.1:6000", "set server IP address and port")
	flag.BoolVar(&helpFlag, "help", false, "print options list")

	if flag.Parse(); helpFlag {
		fmt.Fprint(os.Stderr, "server [options]\n\nAvailable options:\n")
		flag.PrintDefaults()
	}
	serverAddr, err := net.ResolveUDPAddr("udp", serverAddrStr)
	if err != nil {
		log.Error("resolving server address", "error", err)
		return
	}
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Error("creating listening connection", "error", err)
		return
	}

	log.Info("server listens incoming messages from clients")
	buf := make([]byte, 128)
	for {
		if bytesRead, addr, err := conn.ReadFromUDP(buf); err != nil {
			log.Error("receiving message from client", "error", err)
		} else {
			s := string(buf[:bytesRead])
			if s == "conn" {
				log.Info("new client", client)
				_, err2 := conn.WriteToUDP([]byte(strconv.Itoa(client)), addr)
				if err2 != nil {
					log.Error("sending message to client", "error", err, "client", addr.String())
				}
				client++
				global = append(global, -1)
				continue

			} else {
				numpackage, globalpackage, point1x, point1y, point2x, point2y, numclient := parserequest(s)
				if globalpackage > global[numclient] {
					global[numclient] = globalpackage
					if results[numpackage*100+numclient] == 0 {

						result = calculation(point1x, point1y, point2x, point2y)
						results[numpackage*100+numclient] = result
						_, err1 := conn.WriteToUDP([]byte(strconv.Itoa(numpackage)+","+strconv.FormatFloat(result, 'f', 6, 64)), addr)
						if err1 != nil {
							log.Error("sending message to client", "error", err, "client", addr.String())
						} else {
							log.Info("successful interaction to client", "result", result, "client", addr.String())
						}
					} else {
						result := results[numpackage*100+numclient]
						_, err2 := conn.WriteToUDP([]byte(strconv.Itoa(numpackage)+","+strconv.FormatFloat(result, 'f', 6, 64)), addr)
						if err2 != nil {
							log.Error("sending message to client", "error", err, "client", addr.String())
						} else {
							log.Info("successful interaction to client", "result", result, "client", addr.String())
						}
					}
				} else {
					log.Error("duplicate package", "error")
				}
			}
		}
	}
}
