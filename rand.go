package gp

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

const (
	alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var (
	defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandDigits(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += strconv.Itoa(defaultRand.Intn(10))
	}
	return s
}

func RandString(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(alphanum[defaultRand.Intn(len(alphanum))])
	}
	return buf.String()
}
