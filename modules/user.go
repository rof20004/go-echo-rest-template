package modules

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
	"github.com/rof20004/go-echo-rest-template/model"
)

// User model
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// UserController - database reference controller to user
type UserController struct {
	DB *sql.DB
}

// NewUserController - function to use database
func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

// Create - insert an user
func (u *UserController) Create(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	password, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	strsql := "INSERT INTO users(username, password, name, email, phone) VALUES (?, ?, ?, ?, ?)"

	_, err = u.DB.Exec(strsql, user.Username, password, user.Name, user.Email, user.Phone)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, &model.CustomResponse{Code: http.StatusOK, Message: "User saved successfully"})
}

// Read - search an user
func (u *UserController) Read(c echo.Context) error {
	return c.JSON(http.StatusOK, &model.CustomResponse{Code: http.StatusOK, Message: "User found in this database" )
}

// Update - update an usuario
func (u *UserController) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, &model.CustomResponse{Code: http.StatusOK, Message: "User updated successfully"})
}

// Delete - delete an usuario
func (u *UserController) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, &model.CustomResponse{Code: http.StatusOK, Message: "User deleted successfully"})
}

// List - list all users
func (u *UserController) List(c echo.Context) error {
	rows, err := u.DB.Query("SELECT id, username, name, email, phone FROM users")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer rows.Close()

	var (
		users                        = []*User{}
		id                           sql.NullInt64
		username, name, email, phone sql.NullString
	)

	for rows.Next() {
		if err = rows.Scan(&id, &username, &name, &email, &phone); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		users = append(users, &User{ID: id.Int64, Username: username.String, Name: name.String, Email: email.String, Phone: phone.String})
	}

	return c.JSON(http.StatusOK, users)
}

// HashPassword - crypt password string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - decrypt password string
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
