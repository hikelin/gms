package gms

const (
	AUTO_ACKNOWLEDGE int = 1

	CLIENT_ACKNOWLEDGE int = 2

	DUPS_OK_ACKNOWLEDGE int = 3

	SESSION_TRANSACTED int = 0
)

type Session interface {
	createBytesMessage() BytesMessage

	createMapMessage() MapMessage

	createMessage() Message

	createTextMessage() TextMessage

	createTextMessageWithText(text string) TextMessage

	getTransacted() bool

	getAcknowledgeMode() int

	commit()

	rollback()

	close()

	recover()

	getMessageListener() MessageListener

	setMessageListener(listener MessageListener)

	run()

	createProducer(destination Destination) Producer

	createConsumer(destination Destination) Consumer

	createConsumerWithSelector(destination Destination, messageSelector string) Consumer

	createConsumerWithSelectorLocal(destination Destination, messageSelector string, NoLocal bool)

	createQueue(queueName string) Queue

	createTopic(topicName string) Topic

	createDurableSubscriber(topic Topic, name string) TopicSubscriber

	createDurableSubscriberWithSelector(
		topic Topic,
		name string,
		messageSelector string,
		noLocal bool) TopicSubscriber

	createBrowser(queue Queue) QueueBrowser

	createBrowserWithSelector(queue Queue, messageSelector string) QueueBrowser

	createTemporaryQueue() TemporaryQueue

	createTemporaryTopic() TemporaryTopic

	unsubscribe(name string)

	createSharedConsumer(topic Topic, sharedSubscriptionName string) Consumer

	createSharedConsumerWithSelector(topic Topic, sharedSubscriptionName string, messageSelector string) Consumer

	createDurableConsumer(topic Topic, name string) Consumer

	createDurableConsumerWithSelectorLocal(topic Topic, name string, messageSelector string, noLocal bool) Consumer

	createSharedDurableConsumer(topic Topic, name string) Consumer

	createSharedDurableConsumerWithSelector(topic Topic, name string, messageSelector string) Consumer
}
