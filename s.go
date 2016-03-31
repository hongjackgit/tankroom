package tank

import (
    "net"
    "fmt"
    "encoding/json"
)

type tankService struct{
    Address string
}

var  tr  TankRoom

func CreateTankService(address string) *tankService{
    return &tankService{address}
}
func(tankservice *tankService) Start(){
    tcpAddr,err := net.ResolveTCPAddr("tcp4",tankservice.Address)
    if(err != nil){
        fmt.Println("error")
    }
    l ,_ := net.ListenTCP("tcp4",tcpAddr)
    for{
        conn,err := l.Accept()
        if(err != nil ){
            fmt.Println("conn is error ",err.Error())
        }
        buf :=make([]byte,2048)
        len,err := conn.Read(buf)
        flag := checkError(err)
        if flag == 0 {
                break
        }
        var msg Message
        json.Unmarshal(buf[0:len],&msg)
        tr.InRoom(&msg,conn)
        go tankservice.Send(conn)
    }
}


func(this *tankService) Send(conn net.Conn){
    for {
            buf :=make([]byte,2048)
            len,err := conn.Read(buf)
            flag := checkError(err)
            if flag == 0 {
                    break
            }
            var msg Message
            json.Unmarshal(buf[0:len],&msg)
            go tr.Broad(&msg)
        }
}

func checkError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			return 0
		}
		return -1
	}
	return 1
}
