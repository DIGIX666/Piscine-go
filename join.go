package piscine

func Join(strs []string, sep string) string {
	resultat := ""
	for _, v := range strs {
		resultat += sep + v
	}
	return resultat[1:]
}
