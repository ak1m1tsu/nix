// Package jwt provides functions for creating and validating JSON Web Tokens (JWT).
//
// It uses the "github.com/golang-jwt/jwt/v4" and "github.com/google/uuid" packages.
//
// The package defines two main types: Claims and Details. Claims is used to represent the JWT claims,
// and Details represents the token details including the token itself, its payload, and expiration time.
//
// Example Usage:
//
//	// Create a new JWT token with the given payload, time-to-live (ttl), and RSA private key.
//	tokenPayload := map[string]interface{}{"user_id": 123, "role": "admin"}
//	ttl := time.Hour
//	privateKey := []byte("YOUR_RSA_PRIVATE_KEY")
//	tokenDetails, err := jwt.CreateToken(tokenPayload, ttl, privateKey)
//	if err != nil {
//		// Handle error
//	}
//
//	// Validate a JWT token with the given RSA public key.
//	publicKey := []byte("YOUR_RSA_PUBLIC_KEY")
//	decodedPayload, err := jwt.ValidateToken(tokenDetails.Token, publicKey)
//	if err != nil {
//		// Handle error
//	}
//	fmt.Println("Decoded Payload:", decodedPayload)
//
// Note: The package uses the RS256 signing method for JWTs.
package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// ErrInvalidClaims is the error returned when the JWT claims are invalid.
var ErrInvalidClaims = errors.New("invalid claims")

// Claims represents the custom JWT claims along with standard registered claims.
type Claims struct {
	Payload interface{}
	jwt.RegisteredClaims
}

// Details represents the JWT token details, including the token itself, its payload, and expiration time.
type Details struct {
	ID      uuid.UUID
	Token   string
	Payload interface{}
	Expires time.Time
}

// CreateToken creates a new JWT token with the given payload, time-to-live (ttl), and RSA private key.
// The function generates a new UUID as the token ID and uses the RS256 signing method for JWTs.
// The token is signed with the RSA private key, and its expiration time is set according to the ttl value.
// Returns the token details on success, or an error if any operation fails.
func CreateToken(payload interface{}, ttl time.Duration, key *rsa.PrivateKey) (*Details, error) {
	now := time.Now().UTC()

	details := &Details{
		ID:      uuid.New(),
		Payload: payload,
		Expires: now.Add(ttl),
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

	var err error
	details.Token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, &claims).SignedString(key)
	if err != nil {
		return nil, err
	}

	return details, nil
}

// ValidateToken validates a JWT token using the given RSA public key.
// The function parses the token and verifies its signature using the RSA public key.
// If the token is valid, it returns the payload stored in the token.
// If the token is invalid, expired, or tampered with, an error is returned.
func ValidateToken(token string, key *rsa.PublicKey) (interface{}, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return key, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return claims.Payload, nil
}
