package main

import (
	"context"
	"fmt"
	"time"
)

const ReqId = "reqId"

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ReqId, "123")
	go fun1(1, context.WithValue(ctx, ReqId, 1))
	go fun1(2, context.WithValue(ctx, ReqId, 2))
	go fun1(3, context.WithValue(ctx, ReqId, 3))
	time.Sleep(10 * time.Second)
}

func fun1(id int, cont context.Context) {
	fmt.Println(id, ":", cont.Value(ReqId))
}
