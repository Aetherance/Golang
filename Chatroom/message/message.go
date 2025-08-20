package message

import "encoding/json"

type Message struct {
	From string
	Text string
}

func SerMessage(from string,text string) []byte {
	data,_ := json.Marshal(Message{from,text})
	return data
}

func UnSerMessage(data []byte) Message {
	message := Message{}
	json.Unmarshal(data,&message)
	return message
}