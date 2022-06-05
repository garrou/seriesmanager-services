package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtHelper interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
	ExtractUserId(authorization string) string
}

type jwtClaims struct {
	UserId string `json:"id"`
	jwt.StandardClaims
}

type jwtHelper struct {
	secretKey string
	issuer    string
}

func (j *jwtHelper) GenerateToken(userId string) string {
	claims := &jwtClaims{
		userId,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			Issuer:    j.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		panic(err.Error())
	}
	return t
}

func (j *jwtHelper) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtHelper) ExtractUserId(authorization string) string {
	token, _ := j.ValidateToken(authorization)
	claims := token.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%s", claims["id"])
}

func NewJwtHelper() JwtHelper {
	return &jwtHelper{
		secretKey: os.Getenv("JWT_SECRET"),
		issuer:    os.Getenv("JWT_ISSUER"),
	}
}
