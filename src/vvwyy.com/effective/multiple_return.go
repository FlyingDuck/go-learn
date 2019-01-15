package effective

func NextNum(b []byte, i int) (int, int) {
	for ; i<len(b) && !isDigit(b[i]); i++ {
	}
	
	x := 0
	for ; i<len(b) && isDigit(b[i]) ; i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func isDigit(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}
