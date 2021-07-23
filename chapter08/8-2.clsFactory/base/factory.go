package base

type Class interface {
	Do()
}

var (
	factoryByName = make(map[string]func() Class)
)

func Register(name string, factory func() Class) {
	// fmt.Println("Register factory: ", factory)
	factoryByName[name] = factory
	// fmt.Printf("Register factory: %+v", factoryByName)
}

func Create(name string) Class {
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
