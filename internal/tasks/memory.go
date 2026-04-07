package tasks

import "sync"

type MemoryStorage struct {
	mu    sync.RWMutex
	tasks map[int64]Task
	nextID int64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make(map[int64]Task),
		nextID: 1,
	}
}

func (s *MemoryStorage) Create(title string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := Task{
		ID:    s.nextID,
		Title: title,
		Done:  false,
	}

	s.tasks[s.nextID] = task
	s.nextID++

	return task
}

func (s *MemoryStorage) List() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Task, 0, len(s.tasks))

	for _, t := range s.tasks {
		result = append(result, t)
	}

	return result
}

func (s *MemoryStorage) Get(id int64) (Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, ok := s.tasks[id]
	return t, ok
}

func (s *MemoryStorage) Delete(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tasks[id]; !ok {
		return false
	}

	delete(s.tasks, id)
	return true
}