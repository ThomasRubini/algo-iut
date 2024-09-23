package scan

func (s *impl) UntilEOL() string {
	str := ""
	for {
		if s.Peek() == ";" {
			s.Advance()
			return str[1:]
		} else {
			str += " " + s.Peek()
			s.Advance()
		}
	}
}
