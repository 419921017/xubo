package main

import (
	"fmt"
	eventSys "xubo/chapter06/6-7.eventSys"
	reg "xubo/chapter06/6-8.reg"
)

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event: ", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("global event: ", param)
}

func main() {
	a := new(Actor)
	eventSys.RegisterEvent("OnSkill", a.OnEvent)

	eventSys.RegisterEvent("OnSkill", GlobalEvent)

	reg.CallEvent("OnSkill", 100)
}
