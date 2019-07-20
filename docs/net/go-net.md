# go-net

## IP

``
type IP []byte
``

## 掩码

```go
//掩码定义
type IPMask []byte
//
func (ip IP) Mask(mask IPMask) IP 
````
