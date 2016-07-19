package persistence

import (
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func setup() {
	fmt.Println("Migrating adventurous_mind_test...")
	url := "postgres://localhost/adventurous_mind_test?sslmode=disable"
	allErrors, ok := migrate.UpSync(url, "./migrations")
	if !ok {
		fmt.Println("Oh no ...")
		log.Fatal(allErrors)
	}
}

func teardown() {
	fmt.Println("Demigrating adventurous_mind_test...")
	url := "postgres://localhost/adventurous_mind_test?sslmode=disable"
	allErrors, ok := migrate.DownSync(url, "./migrations")
	if !ok {
		fmt.Println("Oh no ...")
		log.Fatal(allErrors)
	}
}

func TestCreateUser(t *testing.T) {
	err := CreateUser("foo", "foo@example.com", "foopass")
	if err != nil {
		t.Error("Couldn`t create a testuser", err)
	}
}

func TestUserAll(t *testing.T) {
	us := All()

	if len(us) < 2 {
		t.Error("No users found where there should have been two")
	}

	if us[0].Name != "foo" {
		t.Error(us[0].Name, "!= foo")
	}
}

func TestUserFindOne(t *testing.T) {
	err, u := FindOne("foo")
	if err != nil {
		t.Error(err)
	}

	if u.Name != "foo" {
		t.Error(u.Name, "!= foo")
	}
}

func TestUserFindOneNoMatch(t *testing.T) {
	err, u := FindOne("notexisting")

	if err == nil {
		t.Error(u, "mustn`t exist")
	}

	if fmt.Sprintf("%v", err) != "Not found" {
		t.Error(err, "must be \"Not found\"")
	}
}

func TestUserFindMany(t *testing.T) {
	us := FindMany("foo")

	if len(us) == 0 {
		t.Error("No users found where there should have been one")
	}

	if us[0].Name != "foo" {
		t.Error(us[0].Name, "!= foo")
	}
}

func TestUserFindManyNoMatch(t *testing.T) {
	us := FindMany("notexisting")

	if len(us) != 0 {
		t.Error("Users found where there should have been none")
	}
}
