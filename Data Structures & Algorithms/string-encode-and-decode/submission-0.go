type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if strs == nil {
		return "␀NIL"  // Special marker
	}
	if len(strs) == 0 {
		return "␀EMPTY"
	}
	
	// Use length-prefixed encoding
	return strings.Join(strs, "␟") // Use rare Unicode delimiter
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "␀NIL" {
		return nil
	}
	if encoded == "␀EMPTY" {
		return []string{}
	}
	
	return strings.Split(encoded, "␟")
}