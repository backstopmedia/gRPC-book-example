package main

import (
  "fmt"

  "golang.org/x/net/context"
  "google.golang.org/grpc"

  "./proto"
)

func main() {
  conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
  if err != nil {
    panic(err)
  }
  stub := proto.NewStarfriendsClient(conn)

  // Now we can use the stub to make RPCs
  req := &proto.GetFilmRequest{Id: "4"}
  resp, err := stub.GetFilm(context.Background(), req)
  if err != nil {
    panic(err)
  }

  fmt.Println(resp)
}
