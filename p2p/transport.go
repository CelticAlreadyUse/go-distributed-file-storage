package p2p

//peer is a interface that represents the remote node
type Peer interface{}

// Transport that anything can handle the communication
//between the nodes in the network. This can be of the
//from (TCP,UDP,Websockets...)   //communication protocol
type Transport interface{}