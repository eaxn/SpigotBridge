package main

type Saferoutine struct {
	Execute  func()
	Finished bool
}

var CurrentSaferoutines = 0

func (S *Saferoutine) Run() {
	CurrentSaferoutines++
	S.Execute()
	S.Finished = true
	CurrentSaferoutines--
}

func (S *Saferoutine) HasBeenExecuted() bool {
	return S.Finished
}

// Example:
/*

S := Saferoutine{}
S.Execute = func() {
	time.Sleep(time.Second*1)
}
go S.Run()
for !S.HasBeenExecuted() {

}

*/
