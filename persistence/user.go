package persistence

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Id    string
	Name  string
	Email string
}

func AuthUser(name string, pass string) error {
	q := fmt.Sprintf("SELECT pwhash = crypt('%s', pwhash) FROM users WHERE name = '%s'", name, pass)
	_, res := fmt.Println(q)
	if res == nil {
		return nil
	}
	return errors.New("Wrong Username or Password")
}

func CreateUser(name string, email string, pass string) error {
	q := fmt.Sprintf("INSERT INTO users (name,email,pwhash) VALUES ('%s', '%s', crypt('%s', gen_salt('md5')));", name, email, pass)
	log.Printf(q)
	return nil
}

func (u User) UpdatePassword(old string, new string) error {
	err := AuthUser(u.Name, old)
	if err != nil {
		return err
	}
	q := fmt.Sprintf("UPDATE users SET pwhash = crypt('%s', gen_salt('md5')) WHERE name = '%s';", new, u.Name)
	log.Printf(q)
	return nil
}

// --- SELECT pwhash = crypt('password', pwhash) FROM users WHERE name = 'slaxor';
// --- SELECT pwhash = crypt('password', pwhash) FROM users;

func All() []User {
	var us []User
	return us
}

func FindOne(criteria string) (error, User) {
	return errors.New("Not found"), User{}
}

func FindMany(criteria string) []User {
	var us []User
	return us
}
