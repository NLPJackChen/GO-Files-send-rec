package main

import (
	"fmt"
	"io"
	"net"
	"os"
)
func Recvfile(filename string,conn net.Conn) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				return
			}
			if n==0{
				break
			}
			f.Write((buf[:n]))
		}

	}
}
func main()  {
	//监听
	listenner,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("wrong2")
		return
	}
	defer listenner.Close()

	conn,err1 := listenner.Accept()
	if err1 != nil {
		fmt.Println("wrong1")
		return
	}
	defer conn.Close()
	buf := make([]byte,1024)
	var n int
	n,err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ",err)
		return
	}
	filename := string(buf[:n])
	conn.Write([]byte("ok"))
	Recvfile(filename,conn)


}
