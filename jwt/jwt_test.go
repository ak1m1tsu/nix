package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testPayload struct {
	UserID   int
	Username string
}

func Test_CreateToken(t *testing.T) {
	type testCase struct {
		name    string
		payload testPayload
	}

	tc := testCase{
		name: "valid details",
		payload: testPayload{
			UserID:   1,
			Username: "test",
		},
	}

	t.Run(tc.name, func(t *testing.T) {
		rng := rand.Reader

		pk, _ := rsa.GenerateKey(rng, 4096)

		keyPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(pk),
		})

		details, err := CreateToken(tc.payload, time.Minute*15, keyPem)

		require.Nil(t, err)

		require.NotNil(t, details.Payload)

		assert.NotEmpty(t, details.Token)
		assert.NotEmpty(t, details.Expires)
		assert.NotEmpty(t, details.ID)

		assert.Equal(t, tc.payload, details.Payload)
	})
}
