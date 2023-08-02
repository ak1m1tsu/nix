package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	Payload interface{}
	jwt.RegisteredClaims
}

type Details struct {
	ID      uuid.UUID
	Token   string
	Payload interface{}
	Expires time.Time
}

// CreateToken creates a new JWT token details with the given payload, ttl, and RSA private key.
func CreateToken(payload interface{}, ttl time.Duration, key []byte) (*Details, error) {
	now := time.Now().UTC()

	details := &Details{
		ID:      uuid.New(),
		Payload: payload,
		Expires: now.Add(ttl),
	}

	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return nil, err
	}

	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        details.ID.String(),
		},
	}

	details.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, &claims).SignedString(parsedKey)
	if err != nil {
		return nil, err
	}

	return details, nil
}

// ValidateToken validates a JWT token with the given RSA public key.
func ValidateToken(token string, key []byte) (interface{}, error) {
	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return parsedKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims.Payload, nil
}
