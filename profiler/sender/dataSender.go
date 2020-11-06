package sender

type DataSender interface {
	Send(T interface{}) bool
	Stop()
}
