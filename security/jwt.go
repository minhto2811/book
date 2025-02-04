package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createMapClaims(userId string, exp int64) jwt.MapClaims {
	return jwt.MapClaims{
		"sub": userId,                // Subject (ID người dùng)
		"iss": os.Getenv("APP_NAME"), // Issuer (ứng dụng phát hành)
		"aud": "your-audience",       // Audience (người nhận hợp lệ)
		"exp": exp,                   // Thời gian hết hạn (1 giờ)
		"iat": time.Now().Unix(),     // Thời điểm phát hành
		"nbf": time.Now().Unix(),     // Token có hiệu lực ngay lập tức
	}
}

func NewAccessToken(userId string) (string, error) {
	claims := createMapClaims(userId, time.Now().Add(time.Hour).Unix())
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func NewRefreshToken(userId string) (string, error) {
	refreshClaims := createMapClaims(userId, time.Now().Add(time.Hour).Unix())
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	return refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) (*jwt.MapClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedAccessToken.Claims.(*jwt.MapClaims), nil
}

func ParseRefreshToken(refreshToken string) (*jwt.MapClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedRefreshToken.Claims.(*jwt.MapClaims), nil
}
