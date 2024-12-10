package password

import "time"

// for now, i'm going to ignore users because i'll just have an sqlite implementation, which should be local to each user anyways
type PasswordEntry struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Hashed    string    `json:"hashed"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPasswordEntry(name, url, hashed string) *PasswordEntry {
	return &PasswordEntry{
		Name:      name,
		Url:       url,
		Hashed:    hashed,
		CreatedAt: time.Now(),
	}
}

// basic interface to allow for different database implementations
type PasswordRepo interface {
	Create(entry *PasswordEntry) error
	Read(id string) PasswordEntry
	ReadAll() []PasswordEntry
	Update(id string, newEntry *PasswordEntry) error
	Delete(id string) error
}
