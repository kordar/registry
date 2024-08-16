package registry

type Registry interface {
	Listener()
	Register() error
	Get() (interface{}, error)
	Remove() error
}
