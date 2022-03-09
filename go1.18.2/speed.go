package main

type Speed float32
type BlocksPerSecond int
type BlocksPerMinute int
type BlocksPerHour int
type BlocksPerDay int

func (S Speed) ToBlocksPerSecond() BlocksPerSecond {
	return BlocksPerSecond(S)
}

func (S Speed) ToBlocksPerMinute() BlocksPerMinute {
	return BlocksPerMinute(S * 60)
}

func (S Speed) ToBlocksPerHour() BlocksPerHour {
	return BlocksPerHour(S * 3600)
}
