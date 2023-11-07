package mfa

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

func (m *Mfa) GenerateCode() string {
	code, err := totp.GenerateCodeCustom(m.Secret, time.Now(), totp.ValidateOpts{
		Period:    m.Period,
		Skew:      1,
		Digits:    m.Digits,
		Algorithm: m.Algorithm,
	})
	if err != nil {
		panic(err)
	}
	return code
}
func (m *Mfa) Load() {
	key, err := otp.NewKeyFromURL(m.Url)
	if err != nil {
		panic(err)
	}
	m.Period = uint(key.Period())
	m.Digits = key.Digits()
	m.Algorithm = key.Algorithm()
	m.Secret = key.Secret()

	m.Issuer = key.Issuer()
	m.AccountName = key.AccountName()
}
