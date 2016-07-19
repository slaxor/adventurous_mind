package api

import "github.com/kataras/iris"

type UserAPI struct {
	*iris.Context
}

// GET /users
func (u UserAPI) Get() {
	// u.Write("a list of all users")
	u.MustRender("userlist.pug", u)
}

// GET /users/:param1
func (u UserAPI) GetBy(id string) {
	// u.JSON(iris.StatusOK, iris.Map{"ID": id, "Username": "kataras"})
}

// PUT /users/:param1
func (u UserAPI) PutBy(id string) {
	newUsername := string(u.FormValue("username"))
	// myDb.InsertUser(newUsername)
	println("User with id " + id + " has changes his/her username to: " + newUsername)
	println(u.Session())
	u.JSON(iris.StatusOK, iris.Map{"Status": "Success"})
}

// POST /users
func (u UserAPI) Post() {
	newUsername := string(u.FormValue("username"))
	// myDb.UpdateUser(id,{Username:newUsername})
	println(newUsername + " has been inserted to database")
}

// DELETE /users/:param1
func (u UserAPI) DeleteBy(id string) {
	// myDb.DeleteUser(id)
	println("User with id: " + id + " has been removed from database")
}
