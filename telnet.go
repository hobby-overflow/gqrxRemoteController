package main

import (
	"github.com/ziutek/telnet"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SendMessage(msg string) string {
	t, err := telnet.Dial("tcp", "127.0.0.1:7356")
    checkErr(err)
	defer t.Close()

	println(msg)
	if strings.Contains(msg, "\n") == false {
		msg += "\n"
	}
	_, err = t.Write([]byte(msg))
	checkErr(err)
	response := readResponse(t)
	return response
}

func readResponse(conn *telnet.Conn) string {
	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf)
		checkErr(err)
		return string(buf[:n])
	}
}
