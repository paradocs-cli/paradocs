package pkr

type PackerAttributes []PackerAttribute
type PackerAttribute struct {
	Name          string
	Value         string
	Filename      string
	PositionStart struct {
		Line   int
		Column int
	}
	PositionEnd struct {
		Line   int
		Column int
	}
}
