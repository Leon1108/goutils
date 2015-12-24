package goutils

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestEncryptAES(t *testing.T) {
	plaintext := []byte("1*1*%23*%23*%23*%23*427632662*%23*23425*3*12*0*18*10*1431193733989")
	key := []byte{0x08, 0x08, 0x04, 0x0b, 0x02, 0x0f, 0x0b, 0x0c, 0x01, 0x03, 0x09, 0x07, 0x0c, 0x03, 0x07, 0x0a}
	iv := []byte{0x02, 0x0c, 0x01, 0x0b, 0x01, 0x0b, 0x0b, 0x0f, 0x03, 0x02, 0x05, 0x02, 0x0c, 0x03, 0x0e, 0x0a}

	ciphertext, err := EncryptAES(plaintext, key, iv)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%x", ciphertext)
	if !strings.EqualFold(fmt.Sprintf("%x", ciphertext), "98fa70c3bf1456878babbbba98fa0da4b935085b9130820ee2d337f88fbb30a5acfd318412047c8f5a4f9ed54235874f69c545ac2bddfb2e991c679cfef94015fe64a3921ba28910ad6277109d26295e") {
		t.Fatalf("Not equal")
	}
}

func TestDecryptAES(t *testing.T) {
	cipherhex := "98fa70c3bf1456878babbbba98fa0da4b935085b9130820ee2d337f88fbb30a5acfd318412047c8f5a4f9ed54235874f69c545ac2bddfb2e991c679cfef94015fe64a3921ba28910ad6277109d26295e"
	key := []byte{0x08, 0x08, 0x04, 0x0b, 0x02, 0x0f, 0x0b, 0x0c, 0x01, 0x03, 0x09, 0x07, 0x0c, 0x03, 0x07, 0x0a}
	iv := []byte{0x02, 0x0c, 0x01, 0x0b, 0x01, 0x0b, 0x0b, 0x0f, 0x03, 0x02, 0x05, 0x02, 0x0c, 0x03, 0x0e, 0x0a}

	if cipherbytes, err := hex.DecodeString(cipherhex); err != nil {
		t.Fatal(err)
	} else {
		if plaintbytes, err := DecryptAES(cipherbytes, key, iv); err != nil {
			t.Fatal(err)
		} else if !strings.EqualFold(string(plaintbytes), "1*1*%23*%23*%23*%23*427632662*%23*23425*3*12*0*18*10*1431193733989") {
			t.Fatalf("Not equal")
		} else {
			t.Log("Success")
		}
	}
}
