
package server

import (
	"io"
	"log"
	"net"
	"strconv"
	"github.com/Avyukth/dicedb/config"
)
func readCommand(c net.Conn) (string, error) {
	var buf []byte= make([]byte, 1024)
	n,err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil

}

func RunSyncTCPServer() {
	log.Println("Starting a synchronous TCP server on ", config.Host, ":", config.Port)
	var con_clients int=0

	lsnr,err:=net.Listen("tcp",config.Host+":"+strconv.Itoa(config.Port))

	if err!=nil{
		panic(err)
	}

	for{
		c, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}
		con_clients+=1
		log.Println("Client connected with address: ",c.RemoteAddr(), "Total clients connected: ",con_clients)

		for{
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients-=1
				log.Println("Client disconnected with address: ",c.RemoteAddr(), "Total clients connected: ",con_clients)
				if err == io.EOF {
					break
				}
				log.Println("Error reading command: ", err)
			}
			log.Println("Command received: ", cmd)
			if err = respond(cmd,c); err != nil {
				log.Println("Error responding to command: ", err)
			}
		}
	}
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
