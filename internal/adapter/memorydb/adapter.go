package memorydb

import (
	"errors"
	"sync"

	dmsession "github.com/fbriansyah/bank-ina-test/internal/application/domain/session"
	"github.com/google/uuid"
)

var data *map[string]dmsession.Session
var once sync.Once

type MemoryDatabase struct {
}

// NewMemoryDatabase instantiate data in memory using singleton pattern
func NewMemoryDatabase() *MemoryDatabase {
	once.Do(func() {
		dt := make(map[string]dmsession.Session)

		data = &dt
	})

	return &MemoryDatabase{}
}

func (m *MemoryDatabase) SetSession(id uuid.UUID, sess dmsession.Session) {
	(*data)[id.String()] = sess
}

func (m *MemoryDatabase) GetSession(id uuid.UUID) (dmsession.Session, error) {
	sess, ok := (*data)[id.String()]
	if !ok {
		return dmsession.Session{}, errors.New("invalid session id")
	}

	return sess, nil
}
