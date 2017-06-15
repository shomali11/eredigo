package main

import (
	"fmt"
	"github.com/shomali11/xredis"
)

func main() {
	client := xredis.DefaultClient()
	defer client.Close()

	fmt.Println(client.Set("number", "10"))
	fmt.Println(client.Get("number"))
	fmt.Println(client.Incr("number"))
	fmt.Println(client.Get("number"))
	fmt.Println(client.Del("number"))
}
