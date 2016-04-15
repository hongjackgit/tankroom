package tank

import(
    "fmt"
    "net"
    "math/rand"
    "time"
    "bufio"
    "os"
    "encoding/json"
    "reflect"
    "strings"
    "io"
)

type TankClient struct {
     Conn    net.Conn
     UserName string
     Id  int64
     Inmsg chan *Message
     Outmsg chan *Message
}
func CreateClient() *TankClient{
    tcpaddr := "127.0.0.1:9996"
    tcpAddr, _ := net.ResolveTCPAddr("tcp4", tcpaddr)
    c,err := net.DialTCP("tcp", nil, tcpAddr)
    if(err != nil ){
        fmt.Println("start error")
        os.Exit(0);
    }
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    uid := r.Int63()
    name := fmt.Sprintf("name%d",uid)
    return &TankClient{c,name,uid,make(chan *Message,10),make(chan *Message)}
}

func(this *TankClient) Start(){
    
    fmt.Println(this.UserName+" is coming")
    go this.Recv(this.Conn)
    defer this.Conn.Close()
    inputReader := bufio.NewReader(os.Stdin)
    message := CreateMessage()
    message.Id=2
    message.SendUser=this.UserName
    this.Inmsg<- message
    for{
        go this.Send()
        input, _ := inputReader.ReadString('\n')
        message.Msg = strings.Replace(input, "\n", "", -1)
        message.Id = 2
        message.SendUser = this.UserName
        this.Inmsg<- message
    }
}
func dataSize(v reflect.Value) int {
    if v.Kind() == reflect.Slice {
            if s := sizeof(v.Type().Elem()); s >= 0 {
                    return s * v.Len()
            }
            return -1
    }
    return sizeof(v.Type())
}
func(this *TankClient) Send(){
    for{
        msg1 := <-this.Inmsg
        fmt.Println("send",msg1)
            msgjson,_ := json.Marshal(msg1)
            var mm Message
            json.Unmarshal(msgjson,&mm)
            _ ,err := this.Conn.Write([]byte(msgjson))

            if(err != nil){
                fmt.Println("start error")
            }
            //this.Conn.Close()
    }
}

func(this *TankClient) Recv(conn net.Conn){
    for {
		data := make([]byte, 1024)
		buf := make([]byte, 128)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("Recv error")
			}
			data = append(data, buf[:n]...)
			if n != 128 {
				break
			}
		}

		fmt.Println(string(data))
	}
}
func(this *TankClient) Stop(){
    this.Conn.Close()
}
func sizeof(t reflect.Type) int {
                                fmt.Println("---------111111",t.Kind())
        switch t.Kind() {
        case reflect.Array:
                if s := sizeof(t.Elem()); s >= 0 {
                        return s * t.Len()
                }

        case reflect.Struct:
                sum := 0
                for i, n := 0, t.NumField(); i < n; i++ {
                        
                        fmt.Println("---------",t.Field(i).Type)
                        s := sizeof(t.Field(i).Type)
                        fmt.Println("---------3333",sizeof(t.Field(i).Type))
                        if s < 0 {
                                return -1
                        }
                        sum += s
                }
                return sum

        case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
                reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
                reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
                return int(t.Size())
        }

        return -1
}

