package tank

import(
    "fmt"
    "net"
)

type TankRoom struct {
    Id int64
    Title string
    Conn map[int64]map[string] net.Conn
}

func CreateTankRoom() *TankRoom{
    
    return &TankRoom{}
}
func(this *TankRoom) InRoom(msg *Message,conn net.Conn) *TankRoom {
    if(this.Conn == nil){
        this.Conn = make(map[int64]map[string] net.Conn)
    }
    if(this.Conn[msg.Id] == nil){
        this.Conn[msg.Id] = make(map[string] net.Conn)
    }
    this.Conn[msg.Id][msg.SendUser]=conn
    this.Id=msg.Id
    this.Title="DOTA"
    fmt.Println(this.Conn)
    
    return this
}

func(this *TankRoom) Broad(msg *Message){
    fmt.Println("msg:",msg)
    for _,val := range this.Conn[msg.Id] {
        _ ,err := val.Write([]byte(msg.SendUser+" say: "+msg.Msg))
        if(err != nil){
            delete(this.Conn[msg.Id],msg.SendUser)
            fmt.Println("Send error",err.Error(),this.Conn[msg.Id])
        }
    }
}

func(this *TankRoom) OurRoom(msg *Message,conn net.Conn){
    fmt.Println(msg.SendUser," out ")
}
