package rsacrypt_test

import (
	"go-rsacrypt/rsacrypt"
	"log"
	"testing"
)

const (
	valid_user   = "sdaryani"
	invalid_user = "dummy"
)

func TestRSAPublicKey_Success(t *testing.T) {
	pubkeys, err := rsacrypt.RSAPublicKey(valid_user)
	if err != nil {
		log.Fatalf("Fetch public key error : %s", err)
	}
	pubkey := pubkeys[0].Key // getting only the first public key
	if len(pubkey) == 0 {
		t.Fail()
	}
}

func TestRSAPublicKey_InvalidUser(t *testing.T) {
	_, err := rsacrypt.RSAPublicKey(invalid_user)
	if err == nil {
		t.Fail()
	}
}
