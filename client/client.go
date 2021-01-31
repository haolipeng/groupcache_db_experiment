package client

// Client for dbserver/slowdb

import (
	"fmt"
	"groupcache-db-experiment/api"
	"net/rpc"
)

type Client struct{}

func (c *Client) Get(key string) string {
	client, err := rpc.DialHTTP("tcp", "localhost:9090")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	realKey := &api.Load{key}
	var reply api.ValueResult
	err = client.Call("Server.Get", realKey, &reply)
	if err != nil {
		fmt.Printf("error %s", err)
	}
	return string(reply.Value)
}

func (c *Client) Set(key string, value string) {
	client, err := rpc.DialHTTP("tcp", "localhost:9090")
	if err != nil {
		fmt.Printf("error %s", err)
	}
	realKey := &api.Store{key, value}
	var reply int
	err = client.Call("Server.Set", realKey, &reply)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}
