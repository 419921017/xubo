package eventSys



var EventByName = make(map[string][]func(interface{}))


func RegisterEvent(name string, callback func(interface{})) {
	list := EventByName[name]
	list = append(list, callback)

	EventByName[name] = list

}
