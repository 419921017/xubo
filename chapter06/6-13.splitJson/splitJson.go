package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32
	ResX, ResY int
}

type Battery struct {
	Capacity int
}

func genJsonData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchId bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			2910,
		},
		HasTouchId: true,
	}
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	jsonData := genJsonData()
	fmt.Println(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchId bool
	}{}

	json.Unmarshal(jsonData, &screenAndTouch)
	fmt.Printf("%+v\n", screenAndTouch)

	batteryAndTouch := struct {
		Battery
		HasTouchId bool
	}{}
	json.Unmarshal(jsonData, &batteryAndTouch)
	fmt.Printf("%+v\n", batteryAndTouch)

}
