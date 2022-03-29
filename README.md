# hwid

Quick hack to determine the hardware id of a machine.  This uses the MAC address of the primary network interface and expresses it as a string containing the address in base36.  This isn't particularly robust, but it comes in handly in some cases.

To install the command line utility just to a

```shell
go install github.com/borud/hwid/cmd/hwid@latest
```

If you want to use it in your application as a library, the following is the source of `cmd/hwid` (just because I know you are too lazy to click 😃)

```go

package main

import (
    "flag"
    "fmt"
    "log"
    
    "github.com/borud/hwid"
)

var intf string

func init() {
    flag.StringVar(&intf, "i", "", "network interface")
    flag.Parse()
}

func main() {
     id, err := hwid.ID(intf)
     if err != nil {
         log.Fatal(err)
     }
     fmt.Println(id)
}

```
