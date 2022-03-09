package main

type SpigotObject interface {
	Update()
	GetNewBridgeUpdate() SpigotObject
	GetData() Data
	SetData()
	GetOldBridgeUpdate() SpigotObject
	GetOldGameUpate() SpigotObject
	GetNewGameUpdate() SpigotObject
}

type Data []interface{}
