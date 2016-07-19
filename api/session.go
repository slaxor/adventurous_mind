package api

import (
	"database/sql"
	"log"

	"github.com/kataras/iris"
)

type SessionAPI struct {
	*iris.Context
	Db *sql.DB
}

// GET /sessions/:param1
func (s SessionAPI) GetBy(id string) {
	// u.JSON(iris.StatusOK, iris.Map{"ID": id, "Username": "kataras"})
}

// POST /sessions
func (s SessionAPI) Post() {
	username := string(s.FormValue("username"))
	password := string(s.FormValue("password"))
	// myDb.UpdateUser(id,{Username:newUsername})
	log.Printf("new session with %s:%s has been inserted to database", username, password)
	s.Text(iris.StatusCreated, "session created")
}

// DELETE /sessions/:param1
func (s SessionAPI) DeleteBy(id string) {
	// myDb.DeleteUser(id)
	println("Session " + id + " has been removed")
	s.SetFlash("info", "You are logged out")
	s.Redirect("/", iris.StatusTemporaryRedirect)
}
