package ports


type TokenBlackList interface {
	Add(token string, ttl int64) error
	IsBlackListed(token string) (bool, error)
}