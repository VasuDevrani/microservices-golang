package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// pb "grpc/protos"
)

const (
	port = ":8080"
)

func main() {
	con, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect: &v", err)
	}

	defer con.Close()

	// client := pb.NewGreetServiceClient(con)

	// names := &pb.NameList{
	// 	Names: []string{"Vasu", "Dev", "Com"},
	// }


}