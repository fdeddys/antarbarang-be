package enumerate

type StatusRecord int

const (
	NONACTIVE StatusRecord = iota
	ACTIVE
)

func (s StatusRecord) String() string {
	return [...]string{"ACTIVE", "NONACTIVE"}[s]
}
