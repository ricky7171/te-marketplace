package library_wrapper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type MyJwt interface {
	GenerateStandardClaims(hours int) jwt.RegisteredClaims
	NewToken(method jwt.SigningMethod, claims jwt.Claims) (string, error)
	ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error)
}

type MyJwtImpl struct {
}

func NewMyJwtImpl() *MyJwtImpl {
	return &MyJwtImpl{}
}

func (myJwt *MyJwtImpl) GenerateStandardClaims(hours int) jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Local().Add(time.Hour * time.Duration(hours)),
		},
	}
}

func (myJwt *MyJwtImpl) NewToken(method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(method, claims).SignedString([]byte(SECRET_KEY))
}

func (myJwt *MyJwtImpl) ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, keyFunc, options...)
}
