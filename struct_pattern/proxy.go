package struct_pattern

type DBHandler interface {
	Create(key string, o interface{}) error
	Update(key string, o interface{}) error
	Delete(key string, o interface{}) error
}

type levelDB struct {}

func (db *levelDB) Create(key string, o interface{}) error {
	return nil
}

func (db *levelDB) Update(key string, o interface{}) error {
	return nil
}

func (db *levelDB) Delete(key string, o interface{}) error {
	return nil
}

type StoreProxy struct {
	cache map[string]interface{}
	handler DBHandler
}

func (p *StoreProxy) Create(key string, o interface{}) error {
	// check cache
	_, ok := p.cache[key]
	if ok {
		return nil
	}
	// add to db
	err := p.handler.Create(key, o)
	if err != nil {
		return err
	}
	// add to cache
	p.cache[key] = o
	return nil
}

func (p *StoreProxy) Update(key string, o interface{}) error {
	// update db
	err := p.handler.Update(key, o)
	if err != nil {
		return err
	}
	// update cache
	p.cache[key] = o
	return nil
}

func (p *StoreProxy) Delete(key string, o interface{}) error {
	// delete from db
	err := p.handler.Delete(key, o)
	if err != nil {
		return err
	}
	// delete from cache
	delete(p.cache, key)
	return nil
}

