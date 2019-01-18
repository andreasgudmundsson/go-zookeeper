package main

import (
	"fmt"
	"github.com/andreasgudmundsson/go-zookeeper/zk"
	"os"
	"path"
	"time"
)

func main() {
	conn_s := "127.0.0.1"
	watch_dir := "/"
	if len(os.Args) >= 2 {
		conn_s = os.Args[1]
	}
	if len(os.Args) >= 3 {
		watch_dir = os.Args[2]
	}

	servers, chroot := zk.ParseConnectionString(conn_s)
	fmt.Printf("%+v, chroot=%s\n", servers, chroot)
	c, _, err := zk.Connect(servers, time.Second, zk.WithChroot(chroot))
	if err != nil {
		panic(err)
	}
	seen := make(map[string]int)
	for {
		children, _, ch, err := c.ChildrenW(watch_dir)
		if err != nil {
			panic(err)
		}
		for _, child := range children {
			if _, ok := seen[child]; !ok {
				seen[child] = 1
				p := path.Join(watch_dir, child)
				data, _, err := c.Get(p)
				fmt.Printf("%s:", p)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(string(data))
				}
			}
		}
		for e := range ch {
			fmt.Printf("Type: %v\tState: %v\tPath: %v\tErr:%v\n",
				e.Type, e.State, e.Path, e.Err)
		}
	}
}
