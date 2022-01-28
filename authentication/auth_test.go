package authentication_test

import (
	"testing"

	"github.com/KalbiProject/kalbi/authentication"
	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	Username, Realm, Password, URI, Nonce, Cnonce, Nc, Qop, Method string
}

// MD5Challenge returns computed challenge
func TestMD5Challenge(t *testing.T) {
	expected := "6880451ac59baaffe5aa28eca8afa28c"
	a := TestUser{
		Username: "test",
		Realm:    "test",
		Password: "test",
		URI:      "test",
		Nonce:    "test",
		Cnonce:   "test",
		Nc:       "test",
		Qop:      "test",
		Method:   "test",
	}
	output := authentication.MD5Challenge(a.Username, a.Realm, a.Password, a.URI, a.Nonce, a.Cnonce, a.Nc, a.Qop, a.Method)
	assert.EqualValues(t, expected, output)

}
