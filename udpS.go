package main

import (
    "fmt"
    "math/rand"
    "net"
    "os"
    "strconv"
    "strings"
    "time"
    "bytes"
    "encoding/gob"
)

type Message struct {
	Uuid   string
	Data string
}

func random(min, max int) int {
    return rand.Intn(max-min) + min
}

func main() {
    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide a port number!")
        return
    }
    PORT := ":" + arguments[1]

    s, err := net.ResolveUDPAddr("udp4", PORT)
    if err != nil {
        fmt.Println(err)
        return
    }

    connection, err := net.ListenUDP("udp4", s)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer connection.Close()
    buffer := make([]byte, 1024)
    rand.Seed(time.Now().Unix())

    for {
        _, addr, err := connection.ReadFromUDP(buffer)

        tmpbuff := bytes.NewBuffer(buffer)
        tmpstruct := new(Message)
        // creates a decoder object
        gobobjdec := gob.NewDecoder(tmpbuff)
        // decodes buffer and unmarshals it into a Message struct
        gobobjdec.Decode(tmpstruct)


       // fmt.Print("-> ", string(buffer[0:n-1]))
        fmt.Print("-> ", tmpstruct.Data)
        fmt.Print("-> ", tmpstruct.Uuid)
        if strings.TrimSpace(string(tmpstruct.Data)) == "STOP" {
            fmt.Println("Exiting UDP server!")
            return
        }

        data := []byte(strconv.Itoa(random(1, 1001)))
        fmt.Printf("data: %s\n", string(data))
        _, err = connection.WriteToUDP(data, addr)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}
