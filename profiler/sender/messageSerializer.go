package sender

type MessageSerializer interface {
	serializer(Object interface{}) []byte
}
