package gms

type Connection interface {
	createSession(transacted bool, acknowledgeMode int) Session
}
