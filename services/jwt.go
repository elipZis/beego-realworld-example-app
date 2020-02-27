package services

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/elipZis/beego-realworld-example-app/models"
	"strconv"
	"time"
)

// Some config variables
var JwtSecret = []byte(beego.AppConfig.String("jwt::secret"))
var JwtHeader = beego.AppConfig.String("jwt::header")
var JwtScheme = beego.AppConfig.String("jwt::scheme")

type JwtService struct {
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (service *JwtService) ValidateTokenFromContext(ctx *context.Context) (jwt.MapClaims, error) {
	tokenString, err := service.GetTokenStringFromHeader(ctx)
	if err == nil {
		return service.ValidateToken(tokenString)
	}
	return nil, err
}

func (service *JwtService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := service.ExtractToken(tokenString)
	if err != nil {
		return nil, err
	}
	// Validate that it is ok and return the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("error.token.invalid")
}

func (service *JwtService) GenerateToken(user *models.User) (string, error) {
	// Create the JWT claims
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // Set expiration time
			Id:        strconv.FormatInt(int64(user.Id), 10),
			IssuedAt:  time.Now().Unix(),
			Issuer:    beego.AppConfig.String("appname"),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (service *JwtService) ExtractToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}

		return JwtSecret, nil
	})

	if err == nil {
		return token, nil
	}
	return nil, err
}

func (service *JwtService) GetTokenStringFromHeader(ctx *context.Context) (string, error) {
	auth := ctx.Request.Header.Get(JwtHeader)
	l := len(JwtScheme)
	if len(auth) > l+1 && auth[:l] == JwtScheme {
		return auth[l+1:], nil
	}
	return "", errors.New("error.token.none")
}
