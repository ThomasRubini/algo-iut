package translate

func Operator(s string) string {
	switch s {
	case "ne_vaut_pas":
		return "!="
	case "vaut":
		return "=="
	default:
		return s
	}
}
