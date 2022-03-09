package main

type EventType uint

var EventBlockPlace EventType = 2
var EventBlockBreak EventType = 3
var EventBlockClick EventType = 4

type Event struct {
	EventType EventType
	EventData []Data
}

type BlockPlaceEvent Event
type BlockBreakEvent Event

func (B BlockPlaceEvent) GetBlock() Block {
	return B.EventData[0][0].(Block)
}

func (B BlockBreakEvent) GetBlock() Block {
	return B.EventData[0][0].(Block)
}

type EventHandler struct {
	OnEvent   func(E *Event)
	EventType EventType
}

func (E EventHandler) Call(Event *Event) {
	E.OnEvent(Event)
}

func (E EventType) GetName() string {
	if E == 0 {
		return "TickEvent"
	} else if E == 1 {
		return "TestEvent"
	} else {
		return PublicEventTypes[E]
	}
}

func RegisterEventType(EventType EventType, Name string) {
	PublicEventTypes[EventType] = Name
}

func RegisterEventHandler(EventHandler EventHandler, EventType EventType) {
	PublicHandlers[EventType] = append(PublicHandlers[EventType], EventHandler)
}

var PublicEventTypes = map[EventType]string{}
var PublicHandlers = map[EventType][]EventHandler{} // EventType -> [] EventHandlers

func CallEvent(EventType EventType, Event Event) {
	for x := 0; x < len(PublicHandlers[EventType]); x++ {
		S := Saferoutine{}
		S.Execute = func() {
			PublicHandlers[EventType][x].Call(&Event)
		}
		go S.Run()
		for !S.HasBeenExecuted() {

		}
	}
}
