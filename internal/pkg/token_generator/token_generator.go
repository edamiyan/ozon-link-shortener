package token_generator

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateToken() string {
	token := bytes.Buffer{}
	uniqueTime := time.Now().Unix()
	_, _ = fmt.Fprintf(&token, "%s", convert(uniqueTime, int64(len(alphabet))))

	for len(token.String()) < 10 {
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(len(alphabet))
		_, _ = fmt.Fprintf(&token, "%c", alphabet[int64(number)])

	}
	return token.String()
}

func convert(decimalNumber, n int64) string {
	buf := bytes.Buffer{}
	for decimalNumber > 0 {
		curNumber := decimalNumber % n
		decimalNumber /= n
		_, _ = fmt.Fprintf(&buf, "%c", alphabet[curNumber])
	}
	return buf.String()
}
