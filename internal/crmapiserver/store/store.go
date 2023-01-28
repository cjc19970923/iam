package store

var client Factory

type Factory interface {
	CrmApplet() CrmAppletStore
	Admin() AdminStore
	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}
