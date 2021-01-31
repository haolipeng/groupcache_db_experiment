package main

import (
	"flag"
	"fmt"
	"groupcache-db-experiment/api"
	"groupcache-db-experiment/client"
	"net/rpc"
)

func main() {

	var port = flag.String("port", "9001", "frontend port")
	var set = flag.Bool("set", false, "doing a set?")
	var get = flag.Bool("get", false, "doing a get?")
	var cacheGet = flag.Bool("cget", true, "doing a cache get?")
	var key = flag.String("key", "foo", "key to get")
	var value = flag.String("value", "bar", "value to set")
	flag.Parse()

	client := new(client.Client)
	//cacheGet是通过rpc获取cache缓存内容
	if *cacheGet {
		client, err := rpc.DialHTTP("tcp", "localhost:"+*port)
		if err != nil {
			fmt.Printf("error %s", err)
		}
		realKey := &api.Load{*key}
		var reply api.ValueResult
		err = client.Call("Frontend.Get", realKey, &reply)
		if err != nil {
			fmt.Printf("error %s", err)
		}
		fmt.Println(string(reply.Value))
		return
	}

	//get为数据库读操作
	if *get {
		var reply = client.Get(*key)
		fmt.Println(reply)
		return
	}

	//set为数据库写操作
	if *set {
		client.Set(*key, *value)
		return
	}

	flag.PrintDefaults()
	return

}
