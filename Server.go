package Godis

import "net"

func ListenAndServer(addr string){
	listen, err := net.Listen("TCP", addr)
	defer listen.Close()
	if err!=nil{
		Error(err.Error())
	}
	for {
		conn, err := listen.Accept()
		if err != nil{
			Error(err.Error())
		}
		go doHandle(conn)
	}
}
func doHandle(con net.Conn){
	
}
