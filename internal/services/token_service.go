package services

import (
	"time"

	"first_socket/internal/domain"
	"first_socket/internal/res"
	errorsRes "first_socket/internal/res/errors"

	"github.com/golang-jwt/jwt"
)

type tokenService struct {
}

func (service *tokenService) CreateTokens(login string) (string, string, error) {
	accessClaims := &domain.TokenClaims{
		Login: login,
		Type:  "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Hour,
			).Unix(),
		},
	}
	refreshClaims := &domain.TokenClaims{
		Login: login,
		Type:  "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Hour * 168,
			).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		accessClaims,
	)
	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		refreshClaims,
	)

	accessTokenString, err1 := accessToken.SignedString([]byte(res.JWTKEY))
	refreshTokenString, err2 := refreshToken.SignedString([]byte(res.JWTKEY))

	if err1 != nil || err2 != nil {
		return "", "", errorsRes.TokenGenerateError
	}

	return accessTokenString, refreshTokenString, nil
}

func (service *tokenService) VerifyToken(tokenString string) (string, error) {
	claims, err := service.parseToken(tokenString)

	if err != nil {
		return "", err
	}

	return claims.Login, nil
}

func (service *tokenService) VerifyAccessToken(tokenString string) (string, error) {
	claims, err := service.parseToken(tokenString)

	if err != nil {
		return "", err
	}

	if claims.Type == "access" {
		return claims.Login, nil
	}

	return "", errorsRes.InvalidTokenError
}

func (service *tokenService) RefreshTokens(refreshTokenString string) (string, string, error) {
	claims, err := service.parseToken(refreshTokenString)

	if err != nil {
		return "", "", err
	}

	if claims.Type == "refresh" {
		return service.CreateTokens(claims.Login)
	}

	return "", "", errorsRes.InvalidTokenError
}

func (service *tokenService) parseToken(tokenString string) (*domain.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&domain.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errorsRes.InvalidTokenError
			}
			return []byte(res.JWTKEY), nil
		},
	)

	if err != nil {
		return nil, errorsRes.InvalidTokenError
	}

	if claims, ok := token.Claims.(*domain.TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errorsRes.InvalidTokenError
}

func NewTokenService() ITokenService {
	return &tokenService{}
}
