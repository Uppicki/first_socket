package services

type ITokenService interface {
	CreateTokens(string) (string, string, error)
	VerifyToken(string) (string, error)
	VerifyAccessToken(string) (string, error)
	RefreshTokens(string) (string, string, error)
}

type IUserService interface {
}
