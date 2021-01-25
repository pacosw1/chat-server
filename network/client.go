package network

import (
	pb "github/pacosw1/chat-server/proto"

	"github.com/google/uuid"
)

//Client t
type Client struct {
	ID       string
	IPAddr   string
	Username string
	Location string
}

//NewClient creates new client struct instance
func NewClient(user *pb.SessionRequest) (*Client, error) {

	id, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	return &Client{
		ID:       id.String(),
		Username: user.Username,
		IPAddr:   user.IpAddr,
		Location: user.RelativeLocation,
	}, nil
}
