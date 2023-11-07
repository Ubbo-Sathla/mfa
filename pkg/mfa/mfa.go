package mfa

import (
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

func (m *Mfa) GenerateCode(secret string, period uint, digits otp.Digits, algorithm otp.Algorithm) string {
	code, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    period,
		Skew:      1,
		Digits:    digits,
		Algorithm: algorithm,
	})
	if err != nil {
		panic(err)
	}
	return code
}
func (m *Mfa) Display() {
	key, err := otp.NewKeyFromURL(m.Url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s Issuer: %s Account: %s Code: %s\n", m.Name, key.Issuer(), key.AccountName(), m.GenerateCode(key.Secret(), uint(key.Period()), key.Digits(), key.Algorithm()))
}
