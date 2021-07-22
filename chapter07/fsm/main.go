package main

import (
	"fmt"
	"xubo/chapter07/fsm/info"
	"xubo/chapter07/fsm/state"
	statemgr "xubo/chapter07/fsm/stateMgr"
)

type IdleState struct {
	info.StateInfo
}

func (i *IdleState) OnBegin() {
	fmt.Println("IdleState begin")
}

func (i *IdleState) OnEnd()  {
	fmt.Println("IdleState end")
}


type MoveState struct {
	info.StateInfo
}
func (m *MoveState) OnBegin() {
	fmt.Println("MoveState begin")
}
// 允许状态之间互相转换
func (m *MoveState)EnableSameTransit() bool  {
	return true
}

type JumpState struct {
	info.StateInfo
}

func (j *JumpState) OnBegin()  {

}

func (j *JumpState) CanTransitTo(name string) bool  {
	return name != "MoveState"
}


func main() {
	sm := statemgr.NewStateManager()
	sm.OnChange = func(from, to state.State) {
		fmt.Printf("%s --> %s \n\n", state.StateName(from), state.StateName(to))
	}

	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")
}

func transitAndReport(sm *statemgr.StateManager, target string) {
	if err := sm.Transit(target); err!=nil {
	fmt.Printf("FAILED! %s --> %s, %s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}