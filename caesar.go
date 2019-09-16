package main

const (
	rot13 = 13
)

func Rot13(s string) string {
	return caesar(s, rot13)
}

func EncCaesar(s string, key int) string {
	return caesar(s, key)
}

func DecCaesar(s string, key int) string {
	return caesar(s, -key+26)
}

func caesar(s string, key int) string {
	var line string
	for _, ch := range []byte(s) {
		if ch >= 'A' && ch <= 'Z' {
			ch = byte((int(ch-'A')+key)%26 + 'A')
		} else if ch >= 'a' && ch <= 'z' {
			ch = byte((int(ch-'a')+key)%26 + 'a')
		}
		line += string(ch)
	}
	return line
}
