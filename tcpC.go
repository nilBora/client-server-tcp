package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
    "encoding/gob"
    "bytes"
)

type Message struct {
	ID   string
	Data string
}

func main() {
    arguments := os.Args
    if len(arguments) == 1 {
            fmt.Println("Please provide host:port.")
            return
    }

    CONNECT := arguments[1]
    c, err := net.Dial("tcp", CONNECT)
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">> ")
        text, _ := reader.ReadString('\n')

        msg := Message{ID: "ID1", Data: text}

        binBuf := new(bytes.Buffer)
        gobobj := gob.NewEncoder(binBuf)
        gobobj.Encode(msg)

        //fmt.Fprintf(c, msg+"\n")
        c.Write(binBuf.Bytes())

        message, _ := bufio.NewReader(c).ReadString('\n')
        fmt.Print("->: " + message)
        if strings.TrimSpace(string(text)) == "STOP" {
            fmt.Println("TCP client exiting...")
            return
        }
    }
}
