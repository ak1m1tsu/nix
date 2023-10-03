package jwt

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testPayload struct {
	UserID   int
	Username string
}

func loadRSAKeys() ([]byte, []byte) {
	pvPem, _ := os.ReadFile("testdata/private.pem")
	pubPem, _ := os.ReadFile("testdata/public.pem")
	return pvPem, pubPem
}

func Test_ValidateToken(t *testing.T) {
	type testCase struct {
		name        string
		token       string
		expectedErr error
		key         []byte
	}

	_, pub := loadRSAKeys()

	testCases := []testCase{
		{
			name:        "valid token",
			token:       "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVzZXJJRCI6MSwiVXNlcm5hbWUiOiJ0ZXN0In0sImV4cCI6NTI5MDk4MDc0OSwibmJmIjoxNjkwOTg0MzQ5LCJpYXQiOjE2OTA5ODQzNDksImp0aSI6IjJhOTFlNDdiLWIwZTYtNDI3MS05NWU3LWY3YjdmYjY4MWJiZiJ9.BiudJ6fSKuZX1Onw48Oki8WA5csUltnPIv7EGZgeHYV1oRumYooM6Bwrn7UNYAL1_JCUP6C8VyFsQE-M2NAeOKauQquaNXmtIPSyQcHkFN2FDj4pqL9V0Ci0XCoxtqcuPN027j9mIEj5JfXL_Yj1-RmYlK-JBQuf4EE5loMMpzQfaDqU77hYFcar-F-8Hhc51bv-sSZ9f3En87nlGPmhtWWr7uREhnAjA9e0SW-HZaqU5tUBqgyLCR5xJxEIbTP-NKZBBk6e2wMjAamendVdx8uDob7Z9zmDWFFczOZIFSuWpXQdChO3-HAJobaqeq-wziSXBDl4Q0cRNQrmNp9wJGi0ZrSjrPOW0P2Qt7zQOtqfux65lEfEEjJM6DpY6O4Njx4XzfMK6bQuwtUOsF8BA1MsBkGVLsRcne874rX8pFpQYJDZwfJB4LjoRQVISaJeSArw8Bt9sQ9InTUDY1wmYlnlj7k4h4F_ylv0J3MqEFRVT0NJYcOvRE6jQU9aM5JToDFW9C4WLMaSUxB6195pMgArqvAVVvI-34gKYBdrxGl2gDTYG05uvdZJnbfsHITpRH2c5x-WAveWCXGY-bchHxlUKyqwXq55cw0e3mt9RE2ivZb_0gC8VgH1VzX48aTIrRMLmpYvdgTEViaKvObzJf5vR1Te-zU38pbUTbqHBhU",
			expectedErr: nil,
			key:         pub,
		},
		{
			name:        "invalid token",
			token:       "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVzZXJJRCI6MSwiVXNlcm5hbWUiOiJ0ZXN0In0sImV4cCI6NTI5MDk4MTEzNywibmJmIjoxNjkwOTg0NzM3LCJpYXQiOjE2OTA5ODQ3MzcsImp0aSI6ImZjOWFkZjIxLTg2YTAtNDNjYy1iYjc4LWQxZDhiMWM0M2ExMSJ9.W-fsbiLg1U_CeClbtIgZNCIvAs6lu3Qm2W3Qm5l0Vx52YYQJCxXOUsLZsHiUY52kW37NOmHrfn85gMUFvNV39GNq1y4wcooK2jVCrWFCXRpgBzlWHj-fTISUMr9H1OYUMQdG9Us3RsNvRh4Bg-s9kTWlzzXlNgKIEvb4iZX5v2L9LacrIf9UKb2wcqAWd-qyqcFrSaA6lpRj7cLht4k2PsXztsGf_qXSHa_JEthGFIf4tdDTUf6YSL2J7ouiU0vDbZULKY2gHdCblS7xg_-JZmeyvB7egTouO98GQI2LP8Ua9AzMbzwGCIHqEjtMZ8B7cae7tWZv7OgFFcMITNKUAEYsaVivSTw4sK1Vhd6F1FF6wvx4DfCxjIRPRaYVzT5fW8NK-RBCQXOoLvSJPpalICa1Ild3YZM9rM7xLBQJwHSX-jy6ubHdcQQLxkE5kvw2VAu7Fnuns-hFnC2eDR-wvOri2kdOLr_GOF4AWhGoCBNmplouiSuIyIv0pGj_ZW8P2TtCGHPTuOhfFVoc8-VMGF7-5NBnTK_93JGwrueHZVy4CIfvoSKQrJbd6rM4DMhQkmvfG4s8Y1u_35RLq5eUSeQzj8jqfdGKX2ufqrg1e6-sA5tybIS-9KOqXWrPIEqOCWGEz8-MGqUNusivhGKAo6twZDsl3DQp9GH8OsYS8x0",
			expectedErr: &jwt.ValidationError{},
			key:         pub,
		},
		{
			name:        "invalid claims",
			token:       "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVzZXJJRCI6MSwiVXNlcm5hbWUiOiJ0ZXN0In0sIlRlbXAiOiJ0ZW1wIiwiZXhwIjo1MjkwOTgxNzU4LCJuYmYiOjE2OTA5ODUzNTgsImlhdCI6MTY5MDk4NTM1OCwianRpIjoiNDQ0OWRlMmMtMDMwMS00MGE0LWExNGQtZDNiMTcwY2ExNDcyIn0.eO7tX9cCKSOoYEdEHGW5tPdxx4nZckl021s48CseLaVMP1-yTAHdtD0s7c-61oILC0lQnKXWRu75s2BLEcXnRHHhh1QpSUW4NiHWcWmcR4siHgONOBtclv-bI312vtJ5Oj-wcKHCi3odDs4ZsXEyzsaCjhEoJFvogbWSzh1NfCrhaUsy7lQfTQ1JfBDNstY1aYn9_L1-xhd_hSM0XRYeAVRl0h8JSY643o953TGV80k-4ya0GI7dhNwi6D2OlwFol0uSqWgx8pL-HVjB2h3nU9ZSz_TrBVPxxxLoRqP8eFdjnJBy9wAwV9DqEj57PkaYKkGh0mgYisCsy60-23p__NGXF4lj36wcU6EbnORRpXVnC-pRgQausHrNpE-_qEqi5GGzdpMmgSsjCUfIGSbq8GbnVTm0yS1vk2KNzPNILkPM4BRrM01WFrRXhRB1XLBsCqpkg5yodZgdGmnLMOb7zXdrDDCl2YTn6iv32B6IEEsKzrUi6KDHxXnV1d2xN1BTwL1wwyzIYV-D68mr82AulnMFBcbcIrPIgFJnV4Ami5QvhUGEbjvjdaoQMHDkizIMtTwivzjKutxJhPc0wPJQBkDIUJQE1SaVBOcn41XOsVIa82pyPDPLxJm83PmqQ7A1gw3OOivjKjZY3TEXZwvMOL5-UgewVb5afoOVCORlEOM",
			expectedErr: nil,
			key:         pub,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			key, err := jwt.ParseRSAPublicKeyFromPEM(tc.key)
			require.NoError(t, err)

			payload, err := ValidateToken(tc.token, key)

			assert.IsType(t, tc.expectedErr, err)

			if err == nil {
				t.Log(payload)
				require.NotNil(t, payload)
				require.IsType(t, map[string]interface{}{}, payload)
			}
		})
	}
}

func Test_CreateToken(t *testing.T) {
	type testCase struct {
		name        string
		payload     testPayload
		key         []byte
		expectedErr error
	}

	pk, _ := loadRSAKeys()

	testCases := []testCase{
		{
			name: "valid details",
			payload: testPayload{
				UserID:   1,
				Username: "test",
			},
			key: pk,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			key, err := jwt.ParseRSAPrivateKeyFromPEM(tc.key)
			require.NoError(t, err)

			details, err := CreateToken(tc.payload, time.Hour*999999, key)

			require.Equal(t, tc.expectedErr, err)

			if err == nil {
				require.NotNil(t, details.Payload)

				assert.NotEmpty(t, details.Token)
				assert.NotEmpty(t, details.Expires)
				assert.NotEmpty(t, details.ID)

				assert.Equal(t, tc.payload, details.Payload)
			}
		})
	}
}
