package main

import (
	"fmt"
	"os"

	"github.com/brimless/go-sesame/internal/db/password"
	"github.com/brimless/go-sesame/internal/encryption"
	"github.com/brimless/go-sesame/internal/generator"
)

func main() {
	// to replicate auth, we'll just hardcode a "master password" for now and check if the user input matches
	// if it matches, then we give them access to everything

	const MASTER_PASSWORD string = "testpassword"

	var masterPassword string
	for masterPassword != MASTER_PASSWORD {
		fmt.Printf("Enter your master password: ")
		fmt.Scanln(&masterPassword)
		if masterPassword == MASTER_PASSWORD {
			break
		}
		fmt.Println("Incorrect password. Please try again.")
	}

	// to encrypt/decrypt passwords
	encryptor := encryption.NewAES(masterPassword)

	// db (using json for simplicity)
	dbJson := password.NewPasswordRepoJson()

	var action int
	for {
		fmt.Println("[1] Create a password entry")
		fmt.Println("[2] List password entries")
		fmt.Println("[3] Update password entry")
		fmt.Println("[4] Delete password entry")
		fmt.Println("[5] Exit")
		fmt.Printf("Select an action: ")
		fmt.Scanln(&action)
		fmt.Println()

		var err error
		switch action {
		case 1:
			err = generatePassword(dbJson, encryptor)
		case 2:
			err = showAllPasswords(dbJson, encryptor)
		case 3:
			err = updatePassword(dbJson, encryptor)
		case 4:
			err = deletePassword(dbJson)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func generatePassword(db password.PasswordRepo, encryptor encryption.Encryptor) error {
	// get config params to generate password and password entry info
	var (
		name,
		url string
	)
	fmt.Printf("Enter a name for the password entry: ")
	fmt.Scanln(&name)

	fmt.Printf("Enter the url where this password is used: ")
	fmt.Scanln(&url)

	var length int
	fmt.Printf("Enter the length: ")
	fmt.Scanln(&length)
	fmt.Println()

	// for now, will set all booleans to true
	genConfig, err := generator.NewGeneratorConfig(length, true, true, true)
	if err != nil {
		return err
	}

	// generate the password
	gen := generator.NewNaiveGenerator()
	pw := gen.Generate(genConfig)
	fmt.Printf("Generated the following password: %s\n", pw)

	hashed, err := encryptor.Encrypt(pw)
	if err != nil {
		return err
	}

	// create new entry
	entry := password.NewPasswordEntry(name, url, hashed)
	err = db.Create(entry)
	if err != nil {
		return err
	}

	return nil
}

func updatePassword(db password.PasswordRepo, encryptor encryption.Encryptor) error {
	var id string
	fmt.Printf("Id of entry to update: ")
	fmt.Scanln(&id)

	existingEntry := db.Read(id)
	if existingEntry.Id == "" {
		return fmt.Errorf("Entry does not exist")
	}

	var (
		name,
		url,
		pw string
	)
	fmt.Printf("(previously %s) Enter new name: ", existingEntry.Name)
	fmt.Scanln(&name)

	if name != "" {
		existingEntry.Name = name
	}

	fmt.Printf("(previously %s) Enter new url: ", existingEntry.Url)
	fmt.Scanln(&url)

	if url != "" {
		existingEntry.Url = url
	}

	fmt.Printf("Enter new password: ")
	fmt.Scanln(&pw)

	if pw != "" {
		if len(pw) < 6 || len(pw) > 128 {
			return fmt.Errorf("Password must have at least 6 characters and at most 128 characters.")
		}
		newHashed, err := encryptor.Encrypt(pw)
		if err != nil {
			return err
		}
		existingEntry.Hashed = newHashed

		err = db.Update(id, &existingEntry)
		return err
	}
	return nil
}

func deletePassword(db password.PasswordRepo) error {
	var id string
	fmt.Printf("Id of entry to delete: ")
	fmt.Scanln(&id)

	err := db.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func showAllPasswords(db password.PasswordRepo, encryptor encryption.Encryptor) error {
	entries := db.ReadAll()
	for _, entry := range entries {
		decrypted, err := encryptor.Decrypt(entry.Hashed)
		if err != nil {
			return err
		}
		fmt.Printf("- [%s] %s on %s (%s): %s\n", entry.Id, entry.Name, entry.Url, entry.CreatedAt.Format("2006-01-02 15:04:05"), decrypted)
	}
	fmt.Println()
	return nil
}
