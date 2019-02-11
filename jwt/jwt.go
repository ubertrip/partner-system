package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"github.com/json-iterator/go"
	"golang.org/x/crypto/scrypt"
	"hash"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	salt = []byte("UWqbM9r35bG6JLgda5NtEuXt")
	json = jsoniter.ConfigFastest
)

type JWT struct {
	secret  []byte
	shaPool sync.Pool
}

func New(secretStr string) *JWT {
	secret := []byte(secretStr)
	if len(secret) < 24 {
		log.Fatalln("jwt len(secret) < 24")
	}
	newSecret, err := scrypt.Key(secret, salt, 8192, 5, 1, 32)
	if err != nil {
		log.Fatalln("jwt scrypt.Key:", err)
	}
	return &JWT{secret: newSecret}
}

//

func (rcv *JWT) Encode(s *Session) (token string) {
	je := getEncoder()

	je.jsonenc.Encode(s)
	if je.json.Len() > 0 {
		je.json.Truncate(je.json.Len() - 1)
	}

	je.encoder.Write(je.json.Bytes())
	je.encoder.Close()

	data := rcv.hash(je.body.Bytes(), je.sign[:])
	je.body.WriteString(".")

	je.encoder.Write(data)
	je.encoder.Close()

	token = je.body.String()

	putEncoder(je)
	return
}

func (rcv *JWT) Decode(str string) (session Session, success bool) {
	if index := strings.Index(str, "."); index >= 0 {
		jsonBytes := stringToBytes(str[:index])
		hmacBytes := stringToBytes(str[index+1:])

		jd := getDecoder()

		if bytes.Equal(jd.decode(hmacBytes), rcv.hash(jsonBytes, jd.sign[:])) {
			if json.Unmarshal(jd.decode(jsonBytes), &session) == nil {
				success = session.Expired > time.Now().Unix()
			}
		}

		putDecoder(jd)
	}
	return
}

func (rcv *JWT) hash(message, data []byte) []byte {
	h := rcv.getHMAC()
	h.Write(message)
	data = h.Sum(data[:0])
	rcv.putHMAC(h)
	return data
}

func (rcv *JWT) getHMAC() (h hash.Hash) {
	v := rcv.shaPool.Get()
	if v == nil {
		return hmac.New(sha256.New, rcv.secret)
	}
	h = v.(hash.Hash)
	h.Reset()
	return
}

func (rcv *JWT) putHMAC(h hash.Hash) {
	rcv.shaPool.Put(h)
}
