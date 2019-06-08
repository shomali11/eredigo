package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Ping())
	fmt.Println(client.Echo("Hello"))
	fmt.Println(client.FlushDb())
	fmt.Println(client.FlushAll())
	fmt.Println(client.Info())
}
