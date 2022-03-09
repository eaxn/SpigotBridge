package main

type EntityAI struct {
	Tick func()
	// Damage       func()
	// Death        func()
	// Hunger       func()
	// Effect       func()
	// EntityNearby func()
	EventHandler []EventHandler
	Entity       *Entity
}
