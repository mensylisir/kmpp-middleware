package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

func CreateToken(user entity.User) (string, error) {
	exp := viper.GetInt("jwt.exp")
	secretKey := []byte(viper.GetString("jwt.secret"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":     user.Name,
		"id":       user.ID,
		"isActive": user.IsActive,
		"isAdmin":  user.IsAdmin,
		"roles":    user.Roles,
		"type":     user.Type,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func CreateRefreshToken(user entity.User) (string, error) {
	exp := viper.GetInt("jwt.refreshExp")
	secretKey := []byte(viper.GetString("jwt.secret"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":     user.Name,
		"id":       user.ID,
		"isActive": user.IsActive,
		"isAdmin":  user.IsAdmin,
		"type":     user.Type,
		"roles":    user.Roles,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func VerifyToken(signKey, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected kwt signing method: %v", token.Header["alg"])
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	tok := r.Header.Get("Authorization")

	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:]
	}

	return ""
}

func ExtractTokenMetadata(r *http.Request) (*entity.User, error) {
	signKey := viper.GetString("jwt.secret")
	token, err := VerifyToken(signKey, ExtractToken(r))
	if err != nil {
		return nil, err
	}

	var user entity.User

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		js, err := json.Marshal(claims)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(js, &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, errors.New(constant.TOKEN_INVALID)
}
