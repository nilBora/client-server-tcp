package main

import (
   // "bufio"
    "fmt"
    "net"
    "os"
    "strings"
    "time"
    "bytes"
    "encoding/gob"
)

type Message struct {
	Uuid   string
	Data string
}

var count = 1

func main() {
    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide port number")
        return
    }

    PORT := ":" + arguments[1]
    l, err := net.Listen("tcp", PORT)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer l.Close()

    c, err := l.Accept()
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        tmp := make([]byte, 500)
        c.Read(tmp)
        tmpbuff := bytes.NewBuffer(tmp)
        tmpstruct := new(Message)
        // creates a decoder object
        gobobjdec := gob.NewDecoder(tmpbuff)
        // decodes buffer and unmarshals it into a Message struct
        gobobjdec.Decode(tmpstruct)

        if err != nil {
            fmt.Println(err)
            return
        }
        if strings.TrimSpace(string(tmpstruct.Data)) == "STOP" {
            fmt.Println("Exiting TCP server!")
            return
        }

        fmt.Print("-> Uuid: ", string(tmpstruct.Uuid), " Message: ",string(tmpstruct.Data))
        t := time.Now()
        myTime := t.Format(time.RFC3339) + "\n"
        c.Write([]byte(myTime))
    }
}
