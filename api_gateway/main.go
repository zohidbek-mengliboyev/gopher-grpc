package main

import (
	"log"

	pbu "tutorial_grpc/api_gateway/genproto/book_service"
	pb "tutorial_grpc/api_gateway/genproto/user_service"

	"tutorial_grpc/api_gateway/controller"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	user_service_address = "localhost:50051"
	book_service_address = "localhost:40041"
)

var client pb.UserManagementClient
var bookClient pbu.BookManagementClient

func main() {
	// set up a connection to the server
	conn, err := grpc.Dial(user_service_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client = pb.NewUserManagementClient(conn)

	// set up a connection to the server
	conn, err = grpc.Dial(book_service_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	bookClient = pbu.NewBookManagementClient(conn)

	// ct, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// set up a http server
	g := gin.Default()
	g.POST("/CreateNewUser", controller.CreateNewUser)
	g.GET("/GetUsers", controller.GetUsers)
	g.GET("/GetById/path/:id", controller.GetById)
	g.PUT("/UpdateUser/path/:id", controller.UpdateUser)
	g.DELETE("/DeleteUser/path/:id", controller.DeleteUser)

	g.POST("/Create", controller.Create)
	g.GET("/GetAll", controller.GetAll)
	g.GET("/Get/path/:book_id", controller.Get)
	g.PUT("/Update/path/:book_id", controller.Update)
	g.DELETE("/Delete/path/:book_id", controller.Delete)
	g.GET("/GetUserBooks/path/:user_id", controller.GetUserBooks)

	// Run http server
	if err := g.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
