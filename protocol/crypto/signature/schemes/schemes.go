package schemes

type SignatureScheme uint8

const (
	XMSS SignatureScheme = iota
	Multisignature SignatureScheme
	//SchnorrSignature SignatureAlgorithm
)




