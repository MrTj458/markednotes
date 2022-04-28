package token

import (
	"time"

	"github.com/MrTj458/markednotes"
	"github.com/golang-jwt/jwt"
)

type Jwt struct {
	signingKey []byte
}

type userClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func New(key []byte) *Jwt {
	return &Jwt{key}
}

func (j *Jwt) NewToken(userId int) (string, error) {
	c := userClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString(j.signingKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (j *Jwt) Parse(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (any, error) {
		return j.signingKey, nil
	})
	if err != nil {
		e, ok := err.(*jwt.ValidationError)
		if !ok {
			return 0, err
		}

		switch e.Errors {
		case jwt.ValidationErrorExpired:
			return 0, markednotes.ErrTokenExpired
		default:
			return 0, err
		}
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok || !token.Valid {
		return 0, err
	}

	return claims.UserID, nil
}
