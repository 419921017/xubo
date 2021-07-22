package statemgr

import (
	"errors"
	"xubo/chapter07/fsm/state"
)

type StateManager struct {
	stateByName map[string]state.State
	OnChange    func(from, to state.State)
	curr        state.State
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]state.State),
	}
}

func (sm *StateManager) Add(s state.State) {
	// 获取状态的名称
	name := state.StateName(s)
	// 将s转换成能设置名字的接口, 然后调用接口
	s.(interface {
		setName(name string)
	}).setName(name)
	// 根据状态名获取已经添加的状态, 检查该状态是否存在
	if sm.Get(name) != nil {
		panic("duplicate state: " + name)
	}
	sm.stateByName[name] = s
}

// 根据名字获取指定的状态
func (sm *StateManager) Get(name string) state.State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

// 状态没有找到
var ErrStateNotFound = errors.New("state not found")
// 禁止在同状态之间转移
var ErrForbidSameStateTransit = errors.New("forbid same state transit")
// 不能转移到指定状态
var ErrCannotTransitToState = errors.New("cannot transit to state")

func (sm *StateManager) CurrState() state.State {
	return sm.curr
}

// 当前状态是否能转移到目标状态
func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}
	// 相同的状态不用转换
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}
	// 使用当前状态, 检查能否转移到指定名字的状态
	return sm.curr.CanTransitTo(name)
}

// 转移到指定状态
func (sm *StateManager) Transit(name string) error {
	// 获取目标状态
	next := sm.Get(name)
	if next == nil {
		return ErrStateNotFound
	}
	// 记录转移之前的状态
	pre := sm.curr
	// 如果当前有状态
	if sm.curr != nil {
		// 相同状态不需要转换
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}
		// 不能转移到目标状态
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
		// 结束当前状态
		sm.curr.OnEnd()
	}
	// 将当前状态转换成要转移到目标的状态
	sm.curr = next
	// 调用新状态的开始
	sm.curr.OnBegin()

	// 通知回调
	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}
	return nil
}
