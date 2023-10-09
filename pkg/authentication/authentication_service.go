package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/golang-jwt/jwt"
)

var client = &http.Client{}

type AuthenticationService struct {
	AuthUrl string
}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{
		AuthUrl: os.Getenv("AUTH_URL"),
	}
}

func (as AuthenticationService) CreateConsumer(email string) error {
	consumerUrl := as.AuthUrl + "/consumers"
	body, err := json.Marshal(dto.ConsumerDto{
		Username: email,
	})
	if err != nil {
		return err
	}

	payload := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPost, consumerUrl, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (as AuthenticationService) GenerateJwt(email string) (string, error) {
	jwtUrl := as.AuthUrl + "/consumers/" + email + "/jwt"
	req, err := http.NewRequest(http.MethodPost, jwtUrl, nil)
	if err != nil {
		return "", err
	}
	fmt.Println()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result := dto.JwtResponseDto{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return "", err
	}
	fmt.Println("result", result)
	defer res.Body.Close()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Unix(result.CreateAt, 0).Add(time.Hour * 24).Unix()
	claims["iss"] = result.Key

	tokenString, err := token.SignedString([]byte(result.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
