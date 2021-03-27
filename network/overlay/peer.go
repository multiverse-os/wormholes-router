package overlay

type Peer struct {
	ID muiltiverse.id
	Username string
	Keyring []*wormholes.Key
	SessionKey *wormholes.Key
}


