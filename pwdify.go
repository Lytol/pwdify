package pwdify

func Encrypt(password string, paths []string) error {
	for _, path := range paths {
		err := EncryptFile(password, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func EncryptFile(password, path string) error {
	// TODO: should encrypt file with the password
	return nil
}
