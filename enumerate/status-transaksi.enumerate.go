package enumerate

type StatusTransaksi int

const (
	NEW StatusRecord = iota
	ON_PROCCESS
	ON_THE_WAY
	DONE
)

func (s StatusTransaksi) String() string {
	return [...]string{"NEW", "ON_PROCCESS", "ON_THE_WAY", "DONE"}[s]
}
