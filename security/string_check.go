package security

func IsAlphaNumeric(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') && (charVariable < '0' || charVariable > '9') {
			return false
		}
	}
	return true
}
