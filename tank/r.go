package tank

import(
    "fmt"
    "net"
)

type TankRoom struct {
    Id int64
    Title string
    Conn map[int64]map[net.Conn]    string
}

func CreateTankRoom() *TankRoom{
    
    return &TankRoom{}
}
func(this *TankRoom) InRoom(msg *Message,conn net.Conn) *TankRoom {
    if(this.Conn == nil){
        this.Conn = make(map[int64]map[net.Conn] string)
    }
    if(this.Conn[msg.Id] == nil){
        this.Conn[msg.Id] = make(map[net.Conn] string)
    }
    this.Conn[msg.Id][conn]=msg.SendUser
    this.Id=msg.Id
    this.Title="DOTA"
    return this
}

func(this *TankRoom) Broad(msg *Message){
    fmt.Printf("before %vx \n",this)
    for k,val := range this.Conn[msg.Id] {
        _ ,err := k.Write([]byte(msg.SendUser+" say: "+msg.Msg))
        if(err != nil){
            this.OutRoom(msg.Id,k)
            fmt.Println("Send error",err.Error(),val)
        }
    }
}

func(this *TankRoom) OutRoom(id int64,conn net.Conn){
    delete(this.Conn[id],conn)
    
}
