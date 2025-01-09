package store

import (
	"context"
	"database/sql"
	"time"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockuUserStore{},
	}
}

type MockuUserStore struct{}

func (m *MockuUserStore) Create(ctx context.Context, tx *sql.Tx, u *User) error {
	return nil
}

func (m *MockuUserStore) GetByID(ctx context.Context, id int64) (*User, error) {
	return &User{}, nil
}

func (m *MockuUserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	return &User{}, nil
}

func (m *MockuUserStore) CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error {
	return nil
}

func (m *MockuUserStore) Activate(ctx context.Context, token string) error {
	return nil
}

func (m *MockuUserStore) Delete(ctx context.Context, userID int64) error {
	return nil
}

