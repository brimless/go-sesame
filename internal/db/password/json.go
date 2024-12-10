package password

import (
	"encoding/json"
	"os"
	"strconv"
)

// this is a super bootleg implementation just to test out the theory lol
// your repo will be a json file stored at the specified location
type PasswordRepoJson struct {
	Location string
	Repo     map[string]PasswordEntry
}

func NewPasswordRepoJson() *PasswordRepoJson {
	location := "./tmp/password.json" // for now

	repoFile, err := os.Open(location)
	if err != nil {
		if os.IsNotExist(err) {
			err = createFile(location)
			if err != nil {
				return nil
			}
			repoFile, err = os.Open(location)
			if err != nil {
				return nil
			}
		} else {
			return nil
		}
	}
	defer repoFile.Close()

	var repo map[string]PasswordEntry
	decoder := json.NewDecoder(repoFile)
	if err := decoder.Decode(&repo); err != nil {
		return nil
	}

	return &PasswordRepoJson{
		Location: location,
		Repo:     repo,
	}
}

func (r *PasswordRepoJson) Create(entry *PasswordEntry) error {
	// for now the index will just be an integer
	newIdx := strconv.Itoa(len(r.Repo) + 1)
	entry.Id = newIdx
	r.Repo[newIdx] = *entry
	return overwriteContent(r.Repo, r.Location)
}

func (r *PasswordRepoJson) Read(id string) PasswordEntry {
	entry := r.Repo[id]
	return entry
}

func (r *PasswordRepoJson) ReadAll() []PasswordEntry {
	var entries []PasswordEntry
	for _, entry := range r.Repo {
		entries = append(entries, entry)
	}
	return entries
}

func (r *PasswordRepoJson) Update(id string, newEntry *PasswordEntry) error {
	r.Repo[id] = *newEntry
	return overwriteContent(r.Repo, r.Location)
}

func (r *PasswordRepoJson) Delete(id string) error {
	delete(r.Repo, id)
	return overwriteContent(r.Repo, r.Location)
}

// helpers
func createFile(location string) error {
	if err := os.MkdirAll(getDir(location), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(location)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte("{}"))

	return err
}

func getDir(path string) string {
	var fileName string
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			fileName = path[i+1:]
		}
	}
	if len(fileName) == 0 {
		fileName = path
	}
	return path[:len(path)-len("/"+fileName)]
}

func overwriteContent(newContent map[string]PasswordEntry, location string) error {
	repoFile, err := os.OpenFile(location, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer repoFile.Close()

	encoder := json.NewEncoder(repoFile)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(newContent)
	return err
}
