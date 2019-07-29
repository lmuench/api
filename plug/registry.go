package plug

var Registry []Plug

func Register(plug Plug) {
	Registry = append(Registry, plug)
}
