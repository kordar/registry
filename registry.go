type Registry interface {
	Register() error
	Listener() error
	Get() error
}
