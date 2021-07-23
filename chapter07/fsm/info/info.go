package info

type StateInfo struct {
	name string
}

func (s *StateInfo) Name() string {
	return s.name
}
func (s *StateInfo) setName(name string) {
	s.name = name
}

func (s *StateInfo) EnableSameTransit() bool {
	return false
}
func (s *StateInfo) OnBegin() {

}
func (s *StateInfo) OnEnd() {

}

// 默认可以转移到任何状态
func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}
