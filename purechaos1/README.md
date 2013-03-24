# purechaos1
--
    import "purechaos1"


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
