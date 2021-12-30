package gms

type Producer interface {
	send(message Message)
}
