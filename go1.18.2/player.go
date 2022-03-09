package main

type Player struct {
	// RX = X+XO
	// RY = Y+YO
	// RZ = Z+ZO
	Location Location
}

func (P Player) Move(TargetLocation Location) {}
func (P Player) Teleport()                    {}
func (P Player) SendMessage()                 {}
func (P Player) DisplayMessage()              {}
func (P Player) SetHealth()                   {}
func (P Player) SetHunger()                   {}
func (P Player) SetAI(A EntityAI)             {}
func (P Player) GetInventory()                {}
