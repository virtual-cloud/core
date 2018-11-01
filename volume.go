package core

// CreateVolume retunts is int
func CreateVolume() int {
	return 1
}

// Foo returns foos
type Foo struct {
	name string
}

// Bar return bars
type Bar struct {
	name string
}

// Man return mans
type Man struct {
	Foo
	Bar
}

var foo Foo
