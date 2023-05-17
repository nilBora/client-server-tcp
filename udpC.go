package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
    "encoding/gob"
    "bytes"
    "github.com/google/uuid"
)

type Message struct {
	Uuid   string
	Data string
}

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
            fmt.Println("Please provide a host:port string")
            return
        }
        CONNECT := arguments[1]

        s, err := net.ResolveUDPAddr("udp4", CONNECT)
        c, err := net.DialUDP("udp4", nil, s)
        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
        defer c.Close()

        for {
            reader := bufio.NewReader(os.Stdin)
            fmt.Print(">> ")
            text, _ := reader.ReadString('\n')

            msg := Message{Uuid: uuid.New().String(), Data: text}

            binBuf := new(bytes.Buffer)
            gobobj := gob.NewEncoder(binBuf)
            gobobj.Encode(msg)

            //data := []byte(text + "\n")
            //_, err = c.Write(data)
            _, err = c.Write(binBuf.Bytes())

            if strings.TrimSpace(string(text)) == "STOP" {
                fmt.Println("Exiting UDP client!")
                return
            }

            if err != nil {
                fmt.Println(err)
                return
            }

            buffer := make([]byte, 1024)
            n, _, err := c.ReadFromUDP(buffer)
            if err != nil {
                fmt.Println(err)
                return
            }
            fmt.Printf("Reply: %s\n", string(buffer[0:n]))
        }
}
