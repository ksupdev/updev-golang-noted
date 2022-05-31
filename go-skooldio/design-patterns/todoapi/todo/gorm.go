package todo

import "gorm.io/gorm"

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db}
}

func NewGormStoreFake() *GormStore {
	return &GormStore{db: nil}
}

func (s *GormStore) New(todo *Todo) error {
	if s.db == nil {
		return nil
	}
	return s.db.Create(todo).Error
}

func (s *GormStore) Find(todos *[]Todo) error {
	if s.db == nil {
		return nil
	}
	return s.db.Find(todos).Error
}

func (s *GormStore) Delete(todo *Todo, id int) error {
	if s.db == nil {
		return nil
	}
	return s.db.Delete(todo, id).Error
}
