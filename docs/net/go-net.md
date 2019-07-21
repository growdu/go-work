# go-net

## IP

````go
type IP []byte
````

## 掩码

```go
//掩码定义
type IPMask []byte
//
func (ip IP) Mask(mask IPMask) IP 
````

## TCP

````go
type TCPAddr struct{
	IP IP
	PORT int
}

//创建tcp地址
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
//全双工通信
func (c *TCPConn) Write(b []byte) (n int, err os.Error)
func (c *TCPConn) Read(b []byte) (n int, err os.Error) 
````