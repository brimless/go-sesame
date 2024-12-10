package password

// probably contains db stuff for sqlite
type PasswordRepoSqlite struct{}

func NewPasswordRepoSqlite() *PasswordRepoSqlite {
	return &PasswordRepoSqlite{}
}

func (r *PasswordRepoSqlite) Create(entry *PasswordEntry) error {
	return nil
}

func (r *PasswordRepoSqlite) Read(id string) *PasswordEntry {
	return nil
}

func (r *PasswordRepoSqlite) ReadAll() []PasswordEntry {
	return nil
}

func (r *PasswordRepoSqlite) Update(id string, newEntry *PasswordEntry) error {
	return nil
}

func (r *PasswordRepoSqlite) Delete(id string) error {
	return nil
}
