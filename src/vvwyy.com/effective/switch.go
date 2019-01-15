package effective

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i<len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1;
	}
	return 0
}

func WhatsType(t interface{}) string {
	switch t.(type) {
	case string:
		return "string"
	case bool:
		return "bool"
	case int:
		return "int"
	default:
		return "unknown"
	}
}

