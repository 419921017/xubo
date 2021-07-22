package main

import (
	"fmt"
	"sort"
)

type HeroKind int

const (
	NoneHeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

// func (h Heros) Len() int {
// 	return len(h)
// }

// func (h Heros) Less(i, j int) bool {
// 	if h[i].Kind != h[j].Kind {
// 		return h[i].Kind < h[j].Kind
// 	}
// 	return h[i].Name < h[j].Name
// }

// func (h Heros) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

func main() {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Slice(heros, func(i, j int) bool {
		a := heros[i]
		b := heros[j]
		if a.Kind != b.Kind {
			return a.Kind < b.Kind
		}
		return a.Name < b.Name
	})

	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
