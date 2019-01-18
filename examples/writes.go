package main

import (
	"fmt"
	"github.com/andreasgudmundsson/go-zookeeper/zk"
	"math/rand"
	"os"
	"time"
)

func main() {
	conn_s := "127.0.0.1"
	path := ""
	if len(os.Args) >= 2 {
		conn_s = os.Args[1]
	}
	if len(os.Args) >= 3 {
		path = os.Args[2]
	}

	servers, chroot := zk.ParseConnectionString(conn_s)
	fmt.Printf("%+v, chroot=%s\n", servers, chroot)
	c, _, err := zk.Connect(servers, time.Second, zk.WithChroot(chroot))
	if err != nil {
		panic(err)
	}
	for {
		name, data := genNode(10, 10)
		_, err := c.Create(path+name, data, 0, zk.WorldACL(zk.PermAll))
		fmt.Printf("%s: %s\n", path+name, data)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}

func genNode(n, m int) (string, []byte) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	path := make([]rune, n)
	for i := range path {
		path[i] = letters[rand.Intn(len(letters))]
	}

	data := make([]rune, m)
	for i := range data {
		data[i] = letters[rand.Intn(len(letters))]
	}

	return "/" + string(path), []byte(string(data))

}
