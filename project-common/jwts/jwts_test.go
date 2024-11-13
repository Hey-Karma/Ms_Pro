package jwts

import "testing"

func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzIwMTg3MzksInRva2VuIjoiMTAwNiJ9.0bhsqPh2xcGrHS6TxG2UFqHksw0_xzfyJc8yF76tsV0"
	ParseToken(tokenString, "msproject")
}
