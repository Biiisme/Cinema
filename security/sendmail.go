package security

import (
	"crypto/rand"
	"math/big"
	"net/smtp"
)

func GeneratorOTP(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[randomIndex.Int64()]
	}

	return string(otp), nil
}

func SendSecretCodeToEmail(email, secretCode, secretHashCode string) error {
	// Cấu hình SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUser := "phanthang21091902@gmail.com"
	smtpPass := "kkpo iygq zuje xinh"
	clientDomain := "https://www.meteor-restaurant.com"

	subject := "Subject: Meteor Restaurant Send OTP\n"
	body := "<body><p>Mã OTP của bạn là: <strong>" + secretCode + "</strong></p>" +
		"<p><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"" + clientDomain + "/forgot-change-password?e=" + email + "&token=" + secretHashCode + "\">Change Password</a></p></body>"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + body

	msg := []byte(subject + mime)

	to := []string{email}

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	// Thiết lập chứng thực SMTP
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, to, msg)
	if err != nil {
		return err
	}

	return nil
}
