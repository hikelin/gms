package gms

type Consumer interface {
	close()

	getMessageListener() MessageListener

	getMessageSelector() string

	receive() Message

	receiveWithTimeOut(timeout int64) Message

	setMessageListener(listener MessageListener)
}
