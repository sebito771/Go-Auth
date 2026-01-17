package ports

type PassWordHaser interface{
	Hash(password string) (string ,error)
	Compare(password string, hash string)error
}
