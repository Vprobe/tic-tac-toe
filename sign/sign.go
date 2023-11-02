package sign

const (
	EmptyKey int    = 0
	EmptyVal string = " "

	CrossKey  int    = 1
	CrossVal  string = "X"
	CrossName string = "Crosses"

	NoughtKey  int    = 2
	NoughtVal  string = "O"
	NoughtName string = "Noughts"

	// Side validation
	MinSide int = 1
	MaxSide int = 2
)

var signs = map[int]string{EmptyKey: EmptyVal, CrossKey: CrossVal, NoughtKey: NoughtVal}
var signNames = map[int]string{CrossKey: CrossName, NoughtKey: NoughtName}

func GetSignByNumber(num int) string {
	val, ok := signs[num]
	if ok {
		return val
	}

	return signs[0]
}

func GetSignNameByNumber(num int) string {
	val, ok := signNames[num]
	if ok {
		return val
	}

	return signNames[0]
}
