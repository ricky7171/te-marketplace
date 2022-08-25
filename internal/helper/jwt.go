package helper

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/ricky7171/te-marketplace/internal/library_wrapper"

	"github.com/golang-jwt/jwt/v4"
)

type HelperJwt interface {
	GenerateToken(email string, userId int) (signedToken string, signedRefreshToken string, err error)
	ValidateToken(signedToken string) (claims *SignedTokenDetails, err error)
	ValidateRefreshToken(signedToken string) (claims *SignedRefreshTokenDetails, err error)
}

type HelperJwtImpl struct {
	myJwt library_wrapper.MyJwt
}

// SignedTokenDetails is representation of JWT Token payload
type SignedTokenDetails struct {
	Email string
	ID    string
	jwt.RegisteredClaims
}

// SignedRefreshTokenDetails is representation of JWT Refresh Token payload
type SignedRefreshTokenDetails struct {
	Email string
	ID    string
	jwt.RegisteredClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func NewHelperJwtImpl(myJwt library_wrapper.MyJwt) *HelperJwtImpl {
	return &HelperJwtImpl{
		myJwt: myJwt,
	}
}

// GenerateAllTokens generates both the detailed token and refresh token
func (helper *HelperJwtImpl) GenerateToken(email string, userId int) (signedToken string, signedRefreshToken string, err error) {
	if email == "" || userId < 1 {
		return "", "", errors.New("email or userid cannot be empty")
	}
	userIdString := strconv.Itoa(userId)

	//1. generate claims for token payload
	//token will expired 24 hours
	claims := &SignedTokenDetails{
		Email:            email,
		ID:               userIdString,
		RegisteredClaims: helper.myJwt.GenerateStandardClaims(24),
	}

	//2. generate refresh claims for refresh token payload
	//refresh token will expired 168 hours (1 week)
	refreshClaims := &SignedRefreshTokenDetails{
		ID:               userIdString,
		RegisteredClaims: helper.myJwt.GenerateStandardClaims(168),
	}

	//3. generate token and refresh token according to claims & refreshClaims
	token, err := helper.myJwt.NewToken(jwt.SigningMethodHS256, claims)
	if err != nil {
		token = ""
		return "", "", err
	}

	refreshToken, err := helper.myJwt.NewToken(jwt.SigningMethodHS256, refreshClaims)
	if err != nil {
		token = ""
		refreshToken = ""
	}

	return token, refreshToken, err
}

//ValidateToken validates the jwt token
//convert token jadi data user
func (helper *HelperJwtImpl) ValidateToken(signedToken string) (claims *SignedTokenDetails, err error) {
	if signedToken == "" {
		return nil, errors.New("empty signed token")
	}
	token, err := helper.myJwt.ParseWithClaims(signedToken, &SignedTokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedTokenDetails)

	if !ok {
		return nil, errors.New("the token is invalid")
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, errors.New("the token is expired")
	}

	return claims, nil
}

//convert refresh_token to SignedRefreshTokenDetails that contain user id
func (helper *HelperJwtImpl) ValidateRefreshToken(signedToken string) (claims *SignedRefreshTokenDetails, err error) {
	if signedToken == "" {
		return nil, errors.New("empty signed refresh token")
	}

	token, err := helper.myJwt.ParseWithClaims(
		signedToken,
		&SignedRefreshTokenDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedRefreshTokenDetails)

	if !ok {
		return nil, errors.New("the refresh token is invalid")
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, errors.New("the refresh token is expired")
	}

	return claims, nil
}
