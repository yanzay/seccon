# seccon

Secure TCP connections via SSH for Go. Includes client and server implementation. Feels like plain `net.Dial` and `net.Listen`.

## Installation

```
go get github.com/yanzay/seccon
```

## Examples

### Client
```
// ...
client := seccon.NewClient("username")
conn, err := client.Dial("example.com:2022")
// ...
conn.Write(b)
```

### Server
```
// ...
listener, err := seccon.Listen(":2022", "")
// ...
conn, err := listener.Accept()
// ...
conn.Read(b)
```

See full examples in [examples](https://github.com/yanzay/seccon/tree/master/examples) dir.
