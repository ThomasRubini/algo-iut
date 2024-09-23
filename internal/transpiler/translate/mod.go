package translate

func Operator(s string) string {
	if s == "ne_vaut_pas" {
		return "!="
	} else {
		return s
	}
}
