package decrypt

func StringDecryptor(s string) string {
	decryptedString := ""
	for _, c := range s {
		decryptedString += string(c - 8)
	}
	return decryptedString
}
