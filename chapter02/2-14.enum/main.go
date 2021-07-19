package main

import "fmt"

type Weapon int
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	const (
		Arrow Weapon = iota
		Shuriken
		SniperRifle
		Rifle
		Blower
	)

	fmt.Println(Arrow,
		Shuriken,
		SniperRifle,
		Rifle,
		Blower)

	var weapon Weapon = Blower

	fmt.Println(weapon)

	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)

	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)

	fmt.Printf("%s %d", CPU, CPU)

}
