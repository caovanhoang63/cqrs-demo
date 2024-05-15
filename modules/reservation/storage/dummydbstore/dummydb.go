package dummydbstore

import "context"

type DummyDb struct {
}

func (db *DummyDb) Create(ctx context.Context, v interface{}) error { return nil }
func (db *DummyDb) Delete(ctx context.Context, v interface{}) error { return nil }
func (db *DummyDb) Update(ctx context.Context, v interface{}) error { return nil }
func (db *DummyDb) Select(ctx context.Context, condition map[string]interface{}) (interface{}, error) {
	return nil, nil
}
