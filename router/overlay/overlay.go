package overlay

import (
	"fmt"
)


type Overlay interface {
	String() string

	Connect() (bool, error)
	Disconnect() (bool, error)

	Peers() []*Peer

}





