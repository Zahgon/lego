package challenge

type NetworkStack int

const (
	DualStack NetworkStack = iota
	IPv4Only
	IPv6Only
)

func (s NetworkStack) Network(proto string) string { _ = "STUB: not implemented"; return "" }
