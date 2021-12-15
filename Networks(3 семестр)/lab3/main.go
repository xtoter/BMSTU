package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/skorobogatov/input"
)

type List_client struct {
	Client []string
}
type Note struct {
	Surname string
	Email   string
}
type Data struct {
	Mas_note []Note
}

var newclient bool
var changedata bool

func readconsole() {
	for {
		var command string
		fmt.Println("Введите команду (add,del,look)")
		input.Scanf("%s", &command)
		if command == "add" {

			fmt.Println("Введите Фамилию")
			var surname, email string
			input.Scanf("%s", &surname)
			fmt.Println("Введите почту")
			input.Scanf("%s", &email)
			curnote := Note{surname, email}
			changedata = true
			notes = append(notes, curnote)

		} else if command == "del" {

			var num int
			fmt.Println("Введите номер записи")
			input.Scanf("%d", &num)
			copy(notes[num:], notes[num+1:])
			notes = notes[:len(notes)-1]
			changedata = true
		} else if command == "look" {
			fmt.Println(notes)
			fmt.Println(clients)
		}
	}
}
func newconnections(ln net.Listener, clientsreader map[string]*bufio.Reader, serversconn map[string]net.Conn) {
	for {
		conncur, err := ln.Accept()

		if err != nil {
			fmt.Println("err3", err)
		}
		read := bufio.NewReader(conncur)
		addr, _ := read.ReadString('\n')
		fmt.Println("newclient", addr)
		var m Data
		m.Mas_note = notes
		text, err := json.Marshal(m)
		if err != nil {
			fmt.Println("error1")
		}
		var ser List_client
		ser.Client = clients
		serv, _ := json.Marshal(ser)
		com, _ := read.ReadString('\n')
		if strings.HasPrefix(com, "new") {
			fmt.Println("senddata")
			conncur.Write([]byte(string(serv) + "\n"))
			conncur.Write([]byte(string(text) + "\n"))

		}
		reader := bufio.NewReader(conncur)
		//addr := conncur.RemoteAddr().String()
		clients = append(clients, addr)
		clientsreader[addr] = reader
		serversconn[addr] = conncur
		go newdataasyn(addr, clientsreader, serversconn)
	}
}

var currentadr string
var clients []string

func deletedata(curclient string) {
	for i := 0; i < len(clients); i++ {
		if clients[i] == curclient {
			copy(clients[i:], clients[i+1:])
			clients[len(clients)-1] = ""
			clients = clients[:len(clients)-1]
		}
	}
}
func newdataasyn(client string, clientsreader map[string]*bufio.Reader, serversconn map[string]net.Conn) {
	for {
		test := client
		reader, err := clientsreader[test]
		if !err {
		} else {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Потерян клиент ", client)
				delete(clientsreader, test)
				delete(serversconn, test)
				//
				deletedata(test)
				//
			}
			var new_data Data
			json.Unmarshal([]byte(message), &new_data)
			if err != nil {
				break
			}
			notes = new_data.Mas_note //получили новые данные (синхронизировали)
			fmt.Println("get newdata")
		}
	}
}
func newdata(clientsreader map[string]*bufio.Reader, serversconn map[string]net.Conn) {
	active := make(map[string]bool)
	for {
		for i := 0; i < len(clients); i++ {
			if clients[i] != currentadr {
				if !active[clients[i]] {
					active[clients[i]] = true
					//go newdataasyn(i, &clientsreader, serversconn, active)
				}
			}

		}
	}
}

var notes []Note

func newserv(new []string) {
	for i := 1; i < len(new); i++ {
		new[i] = strings.TrimSpace(new[i])
		cond := false
		for _, v := range clients {
			if v == new[i] {
				cond = true
			}
		}
		if cond == false {
			clients = append(clients, new[i])
			fmt.Println("get new server", new[i])
		}
	}

}
func main() {

	newclient = false
	changedata = false
	if len(os.Args) < 2 {
		fmt.Println("error, need arguments")
		return
	}

	clients = os.Args[1:]
	currentadr := os.Args[1]
	fmt.Println(clients)
	fmt.Println("starting on ", currentadr)
	ln, err1 := net.Listen("tcp", currentadr)
	if err1 != nil {
		fmt.Println(err1)
	}

	serversconn := make(map[string]net.Conn)
	serverreader := make(map[string]*bufio.Reader)
	for i := 1; i < len(clients); i++ { //проходим по всем клиентам
		fmt.Println("подключаюсь к ", clients[i])
		conn, err2 := net.Dial("tcp", clients[i])
		if err2 != nil {
			fmt.Println("error conection to ", clients[i])
			copy(clients[i:], clients[i+1:])
			clients[len(clients)-1] = ""
			clients = clients[:len(clients)-1]
			i--
		} else {
			conn.Write([]byte(currentadr + "\n"))
			if i == 1 {
				conn.Write([]byte("new" + "\n"))
				t := bufio.NewReader(conn)
				message, _ := t.ReadString('\n')
				//fmt.Println(message)
				var new_data List_client
				json.Unmarshal([]byte(message), &new_data)
				fmt.Println(new_data)
				newserv(new_data.Client)
				message1, _ := t.ReadString('\n')
				var new_data1 Data
				json.Unmarshal([]byte(message1), &new_data1)
				notes = new_data1.Mas_note //получили новые данные (синхронизировали)
				fmt.Println("get firstdata")

			} else {
				conn.Write([]byte("old" + "\n"))
			}
			fmt.Println("successfully")
			serversconn[clients[i]] = conn
			reader := bufio.NewReader(conn)

			serverreader[clients[i]] = reader
			go newdataasyn(clients[i], serverreader, serversconn)
		}
	}

	go readconsole()
	go newconnections(ln, serverreader, serversconn)
	//go newdata(serverreader, serversconn)
	for {
		if changedata { //рассылаем во все соединения новые данные
			fmt.Println("send data", clients)
			changedata = false
			for i := 0; i < len(clients); i++ {
				if clients[i] != currentadr {
					conn, err4 := serversconn[clients[i]]
					if !err4 {
						fmt.Println("error2")
					} else {
						var m Data
						m.Mas_note = notes
						text, err := json.Marshal(m)
						if err != nil {
							fmt.Println("error1")
						}
						conn.Write([]byte(string(text) + "\n"))
					}
				}
			}

		}

	}
}
