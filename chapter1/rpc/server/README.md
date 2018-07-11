unlike net/http, RPC need to do a little more manual work

```go
type Listener interface {
    Accep() (Conn, error)

    Close() error

    Addr() Addr
}
```