package main

import (
	"fmt"
	"io"
	"net"
	"os"
)
func SendFile(path string,conn net.Conn)  {
	f,err := os.Open(path)
	if err != nil {
		fmt.Println("wrong")
		return
	}
	defer f.Close()
	//读文件内容
	buf := make([]byte,1024*4)
	for{
		n,err := f.Read(buf)
		if err != nil {
			if err == io.EOF{
				fmt.Println("文件发送完毕")
			}else {
				fmt.Println("wrong")
			}

			return
		}
		conn.Write(buf[:n])
	}
}
func main(){
	fmt.Println("请输入需要传输的文件")
	var path string
	fmt.Scan(&path)
	//获取文件名
	info,err := os.Stat(path)
	if err != nil {
		fmt.Println("wrong")
		return
	}
	fmt.Println("name = ",info.Name())
	//主动链接服务器
	conn,err1 := net.Dial("tcp","127.0.0.1:8000")
	if err1 != nil{
		fmt.Println("err = ",err1)
		return
	}
	defer conn.Close()
	//发送文件名
	_,err = conn.Write([]byte(info.Name()))
	if err1 != nil{
		fmt.Println("err = ",err1)
		return
	}
	//接收对方的回复，如果回复OK，说明对方已经准备好，可以发文件
	var n int
	buf := make([]byte,1024)
	n,err  = conn.Read(buf)
	if err1 != nil{
		fmt.Println("err = ",err1)
		return
	}
	if "ok" == string(buf[:n]){
		//发送文件内容
		SendFile(path,conn)

	}

}
