// Copyright (c) 2011-2013 Simon Schmidt
// Permission to use, copy, modify, and distribute this software and its
// documentation, for any purpose with or without fee is hereby granted, provided
// that the above copyright notice appears in all copies of the source code.
// It is provided "as is" without express or implied warranty.

package purechaos1

type Cipher struct{
	buffer []byte
	key []byte
	kp int
	bp int
	decrypt bool
	lk,lb int
}

func NewEncrypter(key []byte) *Cipher {
	return newCipher(key,false)
}

func NewDecrypter(key []byte) *Cipher {
	return newCipher(key,true)
}

func newCipher(key []byte,decrypt bool) *Cipher {
	lk := len(key)
	lb := lk+1
	c := new(Cipher)
	c.buffer = make([]byte,lb)
	c.key = make([]byte,lk)
	c.bp = 0
	c.kp = 0
	c.decrypt = decrypt
	c.lk,c.lb = lk,lb
	sum := byte(0)
	for i,b := range key {
		c.key[i] = b
		c.buffer[i] = b^0xff
		sum^=b
	}
	c.buffer[lk] = sum
	return  c
}

func (c *Cipher) XORKeyStream(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		if c.decrypt {
			o := c.buffer[c.bp]
			c.buffer[c.bp] = src[i]
			t := decodeByte(c.buffer[c.bp],c.key[c.kp])
			dst[i] = decodeByte(t,o)
		} else {
			t := encodeByte(src[i],c.buffer[c.bp])
			c.buffer[c.bp] = encodeByte(t,c.key[c.kp])
			dst[i] = c.buffer[c.bp]
		}
		c.kp = (c.kp+1)%c.lk
		c.bp = (c.bp+1)%c.lb
	}
}

