package main

import (
    "tank"
    "fmt"
    "os"
  //  "reflect"
)

func main(){
    
     
    if(len(os.Args) == 1){
        fmt.Println("params is empty")

        /*tankClient := &tank.TankClient{make(chan string),"client1"}
        var roomMap = map[string] *tank.TankClient{"client1":tankClient}
        tankRooms := &tank.TankRoom{"title",roomMap}
        fmt.Println(reflect.TypeOf(tankRooms))
        tankRooms.Test()
*/
    } else if(os.Args[1] == "s"){
        tankService := tank.CreateTankService("127.0.0.1:9996")
        tankService.Start()
    } else if(os.Args[1] == "c"){
        tankClient :=tank.CreateClient()
        tankClient.Start()
        
    }
    
}
