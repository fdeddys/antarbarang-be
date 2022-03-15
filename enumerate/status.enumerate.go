package enumerate

type StatusRecord int

const (
	ACTIVE StatusRecord = iota
	NONACTIVE
)

func (s StatusRecord) String() string {
	return [...]string{"ACTIVE", "NONACTIVE"}[s]
}
