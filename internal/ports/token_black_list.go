package ports


type TokenBlackList interface {
	Add(token string) error
	IsBlackListed(token string) (bool, error)
}