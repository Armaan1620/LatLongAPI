package store

import (
	"errors"
	"latlongapi/backend/models"
	"sync"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

// MemoryStore is an in-memory implementation of UserStore
type MemoryStore struct {
	users  map[string]*models.User // email -> user
	usersByID map[int]*models.User // id -> user
	mu     sync.RWMutex
	nextID int
}

// NewMemoryStore creates a new in-memory user store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users:     make(map[string]*models.User),
		usersByID: make(map[int]*models.User),
		nextID:    1,
	}
}

// CreateUser creates a new user
func (s *MemoryStore) CreateUser(email, hashedPassword string) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[email]; exists {
		return nil, ErrUserExists
	}

	user := &models.User{
		ID:        s.nextID,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	s.users[email] = user
	s.usersByID[user.ID] = user
	s.nextID++

	return user, nil
}

// GetUserByEmail retrieves a user by email
func (s *MemoryStore) GetUserByEmail(email string) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[email]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *MemoryStore) GetUserByID(id int) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.usersByID[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

