package baton

func keyJoin(k1, k2 string) []byte {
	if (k1[len(k1) -1 ] == '/' && k2[0] != '/') || (k1[len(k1) -1 ] != '/' && k2[0] == '/') {
		return []byte(k1 + k2)
	}

	if k1[len(k1) -1 ] == '/' && k2[0] == '/' {
		if len(k2) == 1 {
			return []byte(k1)
		}
		return []byte(k1 + k2[1:])
	}


	return []byte(k1 + "/" + k2)
}
