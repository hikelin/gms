package gms

type MessageListener interface {
	onMessage(message Message)
}
