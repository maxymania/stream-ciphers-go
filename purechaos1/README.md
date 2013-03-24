# purechaos1
--
    import "github.com/maxymania/stream-ciphers-go/purechaos1"


## Usage

#### type Cipher

```go
type Cipher struct {
}
```


#### func  NewDecrypter

```go
func NewDecrypter(key []byte) *Cipher
```

#### func  NewEncrypter

```go
func NewEncrypter(key []byte) *Cipher
```

#### func (*Cipher) XORKeyStream

```go
func (c *Cipher) XORKeyStream(dst, src []byte)
```
Note that PureChaos1 does not use XOR. It's just called XOR to implement chipher.Stream.
