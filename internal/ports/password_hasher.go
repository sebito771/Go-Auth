package ports

type PassWordHaser interface{
	Hash(password string) (string ,error)
}
