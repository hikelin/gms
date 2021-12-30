package gms

const (
	DEFAULT_DELIVERY_MODE  int   = PERSISTENT
	DEFAULT_PRIORITY       int   = 4
	DEFAULT_TIME_TO_LIVE   int64 = 0
	DEFAULT_DELIVERY_DELAY int64 = 0
)

type Message interface {
	getGMSMessageID() string

	setGMSMessageID(id string)

	getGMSTimestamp()

	setGMSTimestamp(timestamp int64)

	getGMSCorrelationIDAsBytes() []byte

	setGMSCorrelationIDAsBytes(correlationID []byte)

	setGMSCorrelationID(correlationID string)

	getGMSCorrelationID() string

	getGMSReplyTo() Destination

	setGMSReplyTo(replyTo Destination)

	getGMSDestination() Destination

	setGMSDestination(destination Destination)

	getGMSDeliveryMode() int

	setGMSDeliveryMode(deliveryMode int)

	getGMSRedelivered() bool

	setGMSRedelivered(redelivered bool)

	getGMSType() string

	setGMSType(gmsType string)

	getGMSExpiration() int64

	setGMSExpiration(expiration int64)

	getGMSPriority() int

	setGMSPriority(priority int)

	clearProperties()

	propertyExists(name string) bool

	getboolProperty(name string) bool

	getByteProperty(name string) byte

	getInt32Property(name string) int32

	getIntProperty(name string) int

	getInt64Property(name string) int64

	getFloat32Property(name string) float32

	getFloat64Property(name string) float64

	getStringProperty(name string) string

	setboolProperty(name string) bool

	setByteProperty(name string) byte

	setInt32Property(name string) int32

	setIntProperty(name string) int

	setInt64Property(name string) int64

	setFloat32Property(name string) float32

	setFloat64Property(name string) float64

	setStringProperty(name string) string

	acknowledge()

	clearBody()

	getGMSDeliveryTime() int64

	setGMSDeliveryTime(deliveryTime int64)

	getBody(interface{})

	isBodyAssignableTo(interface{}) bool
}

type TextMessage interface {
	Message
	setText(text string)
	getText() string
}

type BytesMessage interface {
	Message
	setBytes(bytes []byte)
	getBytes() []byte
}

type MapMessage interface {
	Message

	getbool(name string) bool

	getByte(name string) byte

	getInt32(name string) int32

	getInt(name string) int

	getInt64(name string) int64

	getFloat32(name string) float32

	getFloat64(name string) float64

	getString(name string) string

	setbool(name string) bool

	setByte(name string) byte

	setInt32(name string) int32

	setInt(name string) int

	setInt64(name string) int64

	setFloat32(name string) float32

	setFloat64(name string) float64

	setString(name string) string

	getMapNames() []string

	itemExists(name string) bool
}
