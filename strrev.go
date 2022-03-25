package piscine 

func StrRev(s string) string {
	reversed := _
	for i := StrLen(s) : i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}