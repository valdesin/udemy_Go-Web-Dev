package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "uID", 014)
	ctx = context.WithValue(ctx, "fName", "Vlad")

	result, err := dbAcces(ctx)
	if err != nil {
		http.Error(res, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(res, result)
}

func dbAcces(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, (1 * time.Second))
	defer cancel()

	ch := make(chan int)

	go func() {
		uId := ctx.Value("uID").(int)
		//time.Sleep(5 * time.Second)

		if ctx.Err() != nil {
			return
		}

		ch <- uId
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(res, ctx)
}
