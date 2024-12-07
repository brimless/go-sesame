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

// basic interface to allow for different database implementations
type PasswordRepo interface {
	Create(entry *PasswordEntry) error
	Read(id string) PasswordEntry
	Update(id string, newEntry *PasswordEntry) error
	Delete(id string) error
}
