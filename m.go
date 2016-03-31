package tank

import()


type Message struct{
    Msg string
    Id int64
    SendUser string
}

func CreateMessage() *Message {
    
    return &Message{}
}
