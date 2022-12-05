package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtService struct {
	secretkey string
	issure    string
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretkey: "secret-key",
		issure:    "aluno-api",
	}
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretkey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}
		return []byte(s.secretkey), nil
	})
	return err == nil
}
