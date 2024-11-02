// jwt.go
package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Khóa bí mật để mã hóa token (nên lưu ở nơi an toàn, ví dụ biến môi trường)
var jwtSecretKey = []byte("your_secret_key")

// Cấu trúc chứa thông tin về các claims của token
type JWTClaims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT - Tạo JWT với userId và role
func GenerateJWT(userId string, role string) (string, error) {
	// Đặt các claims cho token
	claims := JWTClaims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Thời hạn 24 giờ
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Tạo token với phương thức mã hóa HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token với khóa bí mật
	return token.SignedString(jwtSecretKey)
}
