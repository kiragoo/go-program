package str

func LC389(s, t string) (res byte) {
	for i := range s {
		res ^= s[i] ^ t[i]
	}
	return res ^ t[len(s)-1]
}
