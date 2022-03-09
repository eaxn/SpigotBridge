# SpigotBridge (Go)

## Events and EventHandlers

### Information about EventHandlers
You can possibly create MAX_INT64+1 / MAX_INT32+1 event handlers for MAX_INT64+1 / MAX_INT32+1 possible event types per event and if that isn't enough, you can contribute to the infinite array.

### Create own EventHandlers

#### Create a custom event type
```go
CustomEventType := EventType(2008)
RegisterEventType(CustomEventType, "CustomEventType")
```

### Use EventHandlers

Register an EventHandler (EventBlockBreak)
```go
RegisterEventHandler(EventHandler{
    OnEvent: func(E *Event) {
	    fmt.Println(BlockPlaceEvent(*E).GetBlock().Location)
	},
}, EventBlockBreak)
```
Register a second EventHandler (EventBlockBreak)
```go
RegisterEventHandler(EventHandler{
	OnEvent: func(E *Event) {
	    fmt.Println(BlockPlaceEvent(*E).GetBlock().Location, "2")
	},
}, EventBlockBreak)
```
Call the event and the event handlers meanwhile
```go
CallEvent(EventBlockBreak, Event{
	EventData: []Data{
	    Data{BlockLocation{*big.NewInt(0), *big.NewInt(0), *big.NewInt(4)}, Block{}},
	},
})
```