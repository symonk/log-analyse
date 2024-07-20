package notify

// Notifier is the signature of something that can
// alert or notify a third party system
type Notifier interface {
	Notify() error
}
