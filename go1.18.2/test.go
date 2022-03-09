package main

import (
	"fmt"
	"time"
)

var ClientPort = 12833
var ServerPort = AutoPort("0.0.0.0", time.Millisecond*4)
var PingTimeout = time.Second * 100

func main() {

	RegisterEventHandler(EventHandler{

		OnEvent: func(E *Event) {
			X := BlockBreakEvent(*E).GetBlock().Location.X
			Y := BlockBreakEvent(*E).GetBlock().Location.Y
			Z := BlockBreakEvent(*E).GetBlock().Location.Z
			fmt.Println("Block got broken: X", X, "Y", Y, "Z", Z)
		},
	}, EventBlockBreak)

	C := GetNewClient("0.0.0.0:" + fmt.Sprint(ClientPort))
	Ping := C.Ping("0.0.0.0:"+fmt.Sprint(ServerPort), PingTimeout)
	if Ping < 0 {
		fmt.Println("Impossible ping: No connection possible (bad connection / too high ping) | Timeout: " + PingTimeout.String())
		return
	} else {
		if Ping < 100 && Ping >= 20 {
			fmt.Println("Good ping: Good internet connection.")
		} else if Ping < 20 && Ping >= 5 {
			fmt.Println("Very good ping: Good internet connection.")
		} else if Ping < 5 && Ping <= 100 {
			fmt.Println("Special ping: Connection to 0.0.0.0 ;D")
		} else if Ping > 100 {
			fmt.Println("Medium ping: Medium internet connection.")
		} else if Ping > 400 {
			fmt.Println("Bad ping: Bad internet connection.")
		} else {
			return
		}
	}
	fmt.Println("Connected to the server!")
	C.Listen("0.0.0.0:"+fmt.Sprint(ServerPort), time.Millisecond*0)
	// RegisterEventHandler(EventHandler{
	// 	OnEvent: func(E *Event) {
	// 		fmt.Println(BlockPlaceEvent(*E).GetBlock().Location)
	// 	},
	// }, EventBlockBreak)
	// RegisterEventHandler(EventHandler{
	// 	OnEvent: func(E *Event) {
	// 		fmt.Println(BlockPlaceEvent(*E).GetBlock().Location, "2")
	// 	},
	// }, EventBlockBreak)
	// CallEvent(EventBlockBreak, Event{
	// 	EventData: []Data{
	// 		Data{BlockLocation{*big.NewInt(0), *big.NewInt(0), *big.NewInt(4)}, Block{}},
	// 	},
	// })
}
