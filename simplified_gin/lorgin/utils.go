package lorgin

import "path"

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

func joinPaths(a, b string) string {
	if b == "" {
		return a
	}
	finalPath := path.Join(a, b)
	if lastChar(b) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func assert1(guard bool, text string) {
	if !guard {
		panic(text)
	}
}
