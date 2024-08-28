package authentication

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yazmeyaa/virtual_ideas_book/internal/config"
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	config *config.Config
}

func (a authenticationService) CreateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(a.config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// RefreshToken implements AuthenticationService.
func (a authenticationService) RefreshToken(oldToken string) (string, error) {
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.config.JWT.Secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := newToken.SignedString([]byte(a.config.JWT.Secret))
		if err != nil {
			return "", err
		}

		return tokenString, nil
	}

	return "", errors.New("invalid token")
}

// ValidateToken implements AuthenticationService.
func (a authenticationService) ValidateToken(token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.config.JWT.Secret), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return false, errors.New("token expired")
			}
		}
		return true, nil
	}

	return false, errors.New("invalid token")
}

// HashPassword implements AuthenticationService.
func (a authenticationService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// ValidatePassword implements AuthenticationService.
func (a authenticationService) ValidatePassword(user models.User, password string) bool {
	hash := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func New(config *config.Config) AuthenticationService {
	return authenticationService{
		config: config,
	}
}
