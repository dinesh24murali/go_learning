package p2p

// Message represents any or arbitrary data that is being sent over
// each transport between two nodes in the network
type Message struct {
	Payload []byte
}
