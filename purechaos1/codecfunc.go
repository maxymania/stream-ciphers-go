// Copyright (c) 2011-2013 Simon Schmidt
// Permission to use, copy, modify, and distribute this software and its
// documentation, for any purpose with or without fee is hereby granted, provided
// that the above copyright notice appears in all copies of the source code.
// It is provided "as is" without express or implied warranty.

package purechaos1

func encodeByte(data, key byte) byte {
	return encBox[ (int(key)<<8) | int(data) ]
}

func decodeByte(data, key byte) byte {
	return decBox[ (int(key)<<8) | int(data) ]
}

func encode(dst, src, key []byte) {
	for i,b := range src {
		dst[i] = encBox[ (int(key[i])<<8) | int(b) ]
	}
}

func decode(dst, src, key []byte) {
	for i,b := range src {
		dst[i] = decBox[ (int(key[i])<<8) | int(b) ]
	}
}
