package main

import "github.com/ziutek/telnet"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SendMessage(conn *telnet.Conn, msg string) string {
	_, err := conn.Write([]byte(msg + "\n"))
	checkErr(err)
	response := readResponse(conn)
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
