package login_grpc

import (
	"github.com/ahmedsharyo/SystemContorellerService/login_grpc/pb"
	"google.golang.org/grpc"

	"log"

	"golang.org/x/net/context"
)

func Login(email string, password string) bool {
	log.Println("hey")
	ctx := context.Background()
	request := pb.LoginRequest{Email: email, Password: password}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewLogingInClient(conn)

	response, err := c.Login(ctx, &request)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return false
	}
	log.Printf("Response from server: %s", response)
	return true
}
