package vf

import (
	"crypto/cipher"
	"io"
)

type CipherBox struct {
	Upstream io.Reader
	Key      []byte
	Stream   cipher.Stream
}

func (c *CipherBox) Read(p []byte) (n int, err error) {
	n, err = c.Upstream.Read(p)
	if err != nil && err != io.EOF {
		return
	}
	c.Stream.XORKeyStream(p[:n], p[:n])
	return
}
