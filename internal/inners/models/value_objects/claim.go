package value_objects

import "gopkg.in/square/go-jose.v2/jwt"

type Claims struct {
	Subject  string           `json:"sub"`
	IssuedAt *jwt.NumericDate `json:"iat"`
	Expiry   *jwt.NumericDate `json:"exp"`
	Issuer   string           `json:"iss"`
	Audience *jwt.Audience    `json:"aud"`
	Scope    string           `json:"scope"`
}
