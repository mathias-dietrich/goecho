/*
 echo.go  Simple echo server - echoes back to all connected clients
 author: mathias.dietrioch@devops.ooo
*/

package main

import (
    "net"
    "bufio"
    "strconv"
    "fmt"
)

// change the port to whatever you need
const PORT = 2000

var clients []net.Conn  = make([]net.Conn,0)  

func main() {
    server, _ := net.Listen("tcp", ":" + strconv.Itoa(PORT))
    if server == nil {
        panic("couldn't start listening: " )
    }
    conns := clientConns(server)
    for {
        go handleConn(<-conns)
    }
}

func clientConns(listener net.Listener) chan net.Conn {
    ch := make(chan net.Conn)
    i := 0
    go func() {
        for {
            client,_  := listener.Accept()
            if client == nil {
                fmt.Printf("couldn't accept: ")
                continue
            }
		clients = append(clients, client)
            i++
            fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
            ch <- client
        }
    }()
    return ch
}

func handleConn(client net.Conn) {
    b := bufio.NewReader(client)
    for {
        line,_  := b.ReadBytes('\n')
        for i:=range clients{

            sent, err := clients[i].Write(line)
            if(sent==0){
                fmt.Printf("have 0 bytes, removing client")
                clients = append(clients[:i], clients[i+1:]...)
                return
            }
            if err != nil {
                fmt.Printf("error sending %v, removing client\n", err)
                clients = append(clients[:i], clients[i+1:]...)
                return
            }else{
                fmt.Printf("sent bytes %v to %v\n", sent, clients[i].RemoteAddr())
            }
       }
    }
}
