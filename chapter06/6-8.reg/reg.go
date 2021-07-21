package reg

import (
	"fmt"
	eventSys "xubo/chapter06/6-7.eventSys"
)

func CallEvent(name string, param interface{}) {
	list := eventSys.EventByName[name]

	fmt.Println("list len", len(list))
	for _, callback := range list {
		callback(param)
	}
}
