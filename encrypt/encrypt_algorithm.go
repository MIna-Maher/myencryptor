package encrypt

func StringEncryptor(s string) string {
	encryptedString := ""
	for _, c := range s {
		encryptedString += string(c + 10)
	}
	return encryptedString
}
