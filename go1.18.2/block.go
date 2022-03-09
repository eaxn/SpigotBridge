package main

type Block struct {
	Location   BlockLocation
	LastUpdate *Block
}

func (B Block) Update() {
	// ask server for updates
}

func (B Block) SetMaterial() {
	// B.Update()
	// // Set material
	// B.Update()
}
