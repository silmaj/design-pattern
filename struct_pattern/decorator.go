package struct_pattern

type DBHandle interface {
	Read(key string) (interface{}, error)
	Create(key string, o interface{}) error
	Update(key string, o interface{}) error
	Delete(key string) error
}

type BasicDBHandle struct {

}

func (h *BasicDBHandle) Read(key string) (interface{}, error) {
	// Load form store
	return nil, nil
}

func (h *BasicDBHandle) Create(key string, o interface{}) error {
	// Save to store
	return nil
}

func (h *BasicDBHandle) Update(key string, o interface{}) error {
	// Load form store
	return nil
}

func (h *BasicDBHandle) Delete(key string, o interface{}) error {
	// Delete form store
	return nil
}

type TransactionDecorator struct {
	DBHandle
	fn func()
}

func WarpTransactionDecorator(h DBHandle, fn func()) DBHandle {
	return &TransactionDecorator{
		DBHandle: h,
		fn: fn,
	}
}

func (h *TransactionDecorator) Create(key string, o interface{}) error {
	// Save to store
	// Begin Transaction
	// End Transaction
	return nil
}

func (h *TransactionDecorator) Update(key string, o interface{}) error {
	// Load form store
	// Begin Transaction
	// End Transaction
	return nil
}

func (h *TransactionDecorator) Delete(key string) error {
	// Delete form store
	// Begin Transaction
	// End Transaction
	return nil
}
