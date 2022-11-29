package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/an1l4/go-usermgmt-grpc-part2/usermgmt"
	"google.golang.org/grpc"
)

const (
	address="localhost:50051"
)

func main() {
	conn,err:=grpc.Dial(address,grpc.WithInsecure(),grpc.WithBlock())

	if err!=nil{
		log.Fatalf("failed to connect %v",err)
	}
	defer conn.Close()

	c:=pb.NewUserManagementClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)

	defer cancel()

	var new_user=make(map[string]int32)

	new_user["Alice"]=43
	new_user["Bob"]= 30

	for name,age:=range new_user{
		r,err:=c.CreateNewUser(ctx,&pb.NewUser{Name: name,Age: age})

		if err!=nil{
			log.Fatalf("could not create user %v",err)
		}
		log.Printf(`User Details:
NAME:%s
AGE:%d
ID:%d`,r.GetName(),r.GetAge(),r.GetId())
	}
	params:=&pb.GetUsersParams{}
	r,err:=c.GetUsers(ctx,params)

	if err!=nil{
		log.Fatalf("could not retrieve users: %v",err)
	}
	log.Print("\n User List: \n")
	fmt.Printf("r.GetUsers():%v\n",r.GetUsers())
}