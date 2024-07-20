package access

import "testing"

var maximumTries = 10000000

func Test_GenerateAccessToken(t *testing.T) {
	Test_Randomness := func(t *testing.T) {
		tokenTable := map[AccessToken]string{}

		t.Run("Generate random tokens", func(t *testing.T) {
			for i := 0; i < maximumTries; i++ {
				accessToken := GenerateAccessToken()

				if _, exists := tokenTable[accessToken]; exists {
					t.Errorf("Test_Randomness() generated a duplicated '%s' token", accessToken)
				}
				tokenTable[accessToken] = ""
			}
		})

	}

	Test_Randomness(t)
}
