package main

import (
	"go-redis-protocol/protocol"
	"log"
	"net"
	"os"
)

var (
	Address = "127.0.0.1:6379"
	Network = "tcp"
)

func main()  {
	//args[0] = path
	arg := os.Args[1:]
	redisConn, err := net.Dial(Network, Address)

	defer redisConn.Close()

	if err != nil {
		log.Fatalf("Conn err: %v", err)
	}

	//reqCommand := protocol.GetRequest([]string{
	//	"auth",
	//	"mmclick",
	//})
	command := make([]byte, 1024)
	//_, err = redisConn.Write(reqCommand)
	//n, err := redisConn.Read(command)
	//if err != nil {
	//	log.Fatalf("Conn Read err: %v", err)
	//}

	reqCommand := protocol.GetRequest(arg)

	_, err = redisConn.Write(reqCommand)
	n, err = redisConn.Read(command)
	if err != nil {
		log.Fatalf("Conn Read err: %v", err)
	}

	str , _ := protocol.GetReply(command[:n])

	log.Printf("Reply: %v", str)
}
