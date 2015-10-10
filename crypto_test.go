package gp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	plain := "These pretzels are making me thirsty."
	cipher := "b0804ec967f48520697662a204f5fe72"

	assert.Equal(t, cipher, MD5(plain))
}

func TestSHA1(t *testing.T) {
	plain := "This page intentionally left blank."
	cipher := "af064923bbf2301596aac4c273ba32178ebc4a96"

	assert.Equal(t, cipher, SHA1(plain))
}
