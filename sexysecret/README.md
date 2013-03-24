# sexysecret
--
    import "github.com/maxymania/stream-ciphers-go/sexysecret"


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
Note that SexySecret does not use XOR. It's just called XOR to implement chipher.Stream.
