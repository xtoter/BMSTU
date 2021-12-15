package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	log "github.com/mgutz/logxi/v1"
	"github.com/skorobogatov/input"
)

func parserequest(str string) (packagenum int, result string) {
	var x []int
	for i := 0; i < len(str); i++ {
		if str[i] == ',' {
			x = append(x, i)
		}
	}
	packagenum, err0 := strconv.Atoi(str[:x[0]])
	if err0 != nil {
		log.Error("decode message", "error", err0)
	}
	result = str[x[0]+1:]
	return packagenum, result
}

var globalpackage *int

func get(cond *bool, conn *net.UDPConn, numpackage, point1x, point1y, point2x, point2y int) *time.Timer {
	var timer *time.Timer
	timer = time.AfterFunc(5*time.Second, func() {
		if *cond {
			//out := strconv.Itoa(numpackage) + "," + strconv.Itoa(*globalpackage) + "," + strconv.Itoa(point1x) + "," + strconv.Itoa(point1y) + "," + strconv.Itoa(point2x) + "," + strconv.Itoa(point2y) + "," + client
			tojson := Data{numpackage, *globalpackage, point1x, point1y, point2x, point2y, client}
			out, err := json.Marshal(tojson)
			if err != nil {
				log.Error("sending request to server", "error", err)
			}
			conn.Write([]byte(out))
			*globalpackage++
			timer = get(cond, conn, numpackage, point1x, point1y, point2x, point2y)
		}

	})
	return timer
}
func read(conn *net.UDPConn, out string, buf []byte, numpackage, point1x, point1y, point2x, point2y int) string {
	result := ""
	cond := true
	timer := get(&cond, conn, numpackage, point1x, point1y, point2x, point2y)
	bytesRead, err := conn.Read(buf)
	cond = false
	if err != nil {
		log.Error("receiving answer from server", "error", err)
	}
	defer timer.Stop()
	yStr := string(buf[:bytesRead])
	packagenum, result := parserequest(yStr)
	if packagenum != numpackage {
		result = read(conn, out, buf, numpackage, point1x, point1y, point2x, point2y)
	}

	return result
}

var client string

type Data struct {
	Numpackage    int
	Globalpackage int
	Point1x       int
	Point1y       int
	Point2x       int
	Point2y       int
	Client        string
}

func main() {
	var (
		serverAddrStr string
		n             uint
		helpFlag      bool
	)
	numpackage := 0
	temp := 0
	globalpackage = &temp
	flag.StringVar(&serverAddrStr, "server", "127.0.0.1:6000", "set server IP address and port")
	flag.UintVar(&n, "n", 10, "set the number of requests")
	flag.BoolVar(&helpFlag, "help", false, "print options list")

	if flag.Parse(); helpFlag {
		fmt.Fprint(os.Stderr, "client [options]\n\nAvailable options:\n")
		flag.PrintDefaults()
	} else if serverAddr, err := net.ResolveUDPAddr("udp", serverAddrStr); err != nil {
		log.Error("resolving server address", "error", err)
	} else if conn, err := net.DialUDP("udp", nil, serverAddr); err != nil {
		log.Error("creating connection to server", "error", err)
	} else {
		defer conn.Close()
		buf := make([]byte, 128)
		if _, err := conn.Write([]byte("conn")); err != nil {
			log.Error("sending request to server", "error", err)
		}
		bytesRead, err1 := conn.Read(buf)
		if err1 != nil {
			log.Error("sending request to server", "error", err)
		}
		yStr := string(buf[:bytesRead])

		client = yStr
		for i := uint(0); i < n; i++ {
			x := rand.Intn(1000)
			var point1x, point1y, point2x, point2y int
			fmt.Println("Введите центр окружности(в формате x y)")
			input.Scanf("%d%d", &point1x, &point1y)
			fmt.Println("Введите точку на окружности (в формате x y)")
			input.Scanf("%d%d", &point2x, &point2y)
			fmt.Println("")
			tojson := Data{numpackage, *globalpackage, point1x, point1y, point2x, point2y, client}

			out, err := json.Marshal(tojson)
			//fmt.Println(tojson)
			if err != nil {
				log.Error("sending request to server", "error", err)
			}

			//out := strconv.Itoa(numpackage) + "," + strconv.Itoa(*globalpackage) + "," + strconv.Itoa(point1x) + "," + strconv.Itoa(point1y) + "," + strconv.Itoa(point2x) + "," + strconv.Itoa(point2y) + "," + client

			if _, err := conn.Write([]byte(out)); err != nil {
				log.Error("sending request to server", "error", err, "x", x)
			}
			*globalpackage++
			result := read(conn, string(out), buf, numpackage, point1x, point1y, point2x, point2y)
			fmt.Println(result)
			log.Info("successful interaction with server", "result", result)
			numpackage++
		}
	}
}
