package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"

	"github/pacosw1/chat-server/network"
	pb "github/pacosw1/chat-server/proto"

	"google.golang.org/grpc"
)

const (
	port = ":1290"
)

func main() {

	// start grpc server listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create new grpc server and register its dependencies
	s := grpc.NewServer()
	pb.RegisterChatServer(s, &Server{
		Clients: make(map[string]*network.Client, 0),
		Rooms:   make(map[string][]chan *pb.Message, 0),
	})

	log.Printf("Running grpc service on port: " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

//Server main server stuff
type Server struct {
	pb.UnimplementedChatServer
	Clients map[string]*network.Client
	Rooms   map[string][]chan *pb.Message
}

//CreateSession create session and store user in server clients
func (s *Server) CreateSession(ctx context.Context, user *pb.SessionRequest) (*pb.UserData, error) {

	client, err := network.NewClient(user)

	if err != nil {
		return nil, err
	}

	s.Clients[client.ID] = client

	return &pb.UserData{
		UserID:   client.ID,
		Username: client.Username,
	}, nil
}

//JoinRoom create session and store user in server clients
func (s *Server) JoinRoom(room *pb.RoomID, stream pb.Chat_JoinRoomServer) error {

	if s.Rooms[room.Id] == nil {
		return errors.New("Invalid Room ID")
	}

	msgChannel := make(chan *pb.Message)
	s.Rooms[room.Id] = append(s.Rooms[room.Id], msgChannel)

	for {
		select {
		//close the stream if user disconnects
		case <-stream.Context().Done():
			return nil
		case msg := <-msgChannel:
			//send new messages to client
			stream.Send(msg)
		}
	}
}

//SendMessage create session and store user in server clients
func (s *Server) SendMessage(stream pb.Chat_SendMessageServer) error {

	//receive message
	msg, err := stream.Recv()

	//we simply close connection
	if err == io.EOF {
		return nil
	}

	//we have an err
	if err != nil {
		return err
	}

	//send back sent message ack and close stream
	stream.SendAndClose(&pb.MessageSent{
		MessageID: "123",
	})

	//broadcast the message to all users in room
	go func() {

		members := s.Rooms[msg.RoomID]

		for _, member := range members {
			member <- msg
		}

	}()

	return nil

}

// switch msg.MessageName().Name() {
// case "Typing":
// 	return nil
// case "ChatMessage":
// 	return nil
// case "MessageSent":
// 	return nil
// case "MessageDelivered":
// 	return nil
// case "MessageRead":
// 	return nil
// case "UserJoined":
// 	return nil
// case "UserLeft":
// 	return nil
// }
