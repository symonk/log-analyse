package notify

// Actionable is the signature of something that can
// alert or notify a third party system
type Actionable interface {
	Notify() error
}
