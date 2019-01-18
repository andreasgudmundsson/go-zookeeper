package main

import (
	"fmt"
	"github.com/andreasgudmundsson/go-zookeeper/zk"
	"os"
	"time"
)

func main() {
	conn_s := "127.0.0.1"
	if len(os.Args) >= 2 {
		conn_s = os.Args[1]
	}
	servers, chroot := zk.ParseConnectionString(conn_s)
	fmt.Printf("%+v, chroot=%s\n", servers, chroot)
	c, _, err := zk.Connect(servers, time.Second, zk.WithChroot(chroot))
	if err != nil {
		panic(err)
	}
	children, _, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	for _, c := range children {
		fmt.Printf("%v\n", c)
	}
	for e := range ch {
		fmt.Printf("%v\n", e)
	}
}
