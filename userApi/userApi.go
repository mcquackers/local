package userApi

import (
	"github.com/evilwire/go-env"
	"github.com/labstack/echo"
)

type UserApi struct {
}

type user struct {
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	DayOfBirth string      `json:"dob"`
	Gender     interface{} `json:"gender"`
	ZipCode    uint32      `json:"zip_code"`
	Interests  []string    `json:"interests"`
	Email      string      `json:"email"`
}

func NewUserApi(envReader *goenv.OsEnvReader) (*UserApi, error) {
	return &UserApi{}, nil
}

func (api *UserApi) SignUpUser(c echo.Context) error {
	newUser := &user{}
	c.Bind(newUser)

	//sign up user and stuff

	return c.JSON(200, &struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
