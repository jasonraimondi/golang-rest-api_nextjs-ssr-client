package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"time"
)

var (
	EnableDebugging = kingpin.
		Flag("debug", "Enable Debug.").
		Short('d').
		Bool()
	JwtSecureKey = kingpin.
		Flag("jwt-secure-key", "Secure JWT Key, changing this logs everyone out.").
		Short('k').
		String()
)

func init() {
	kingpin.Parse()
}

func main() {
	e := echo.New()
	e.Debug = *EnableDebugging

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	authRoute := middleware.JWT([]byte(*JwtSecureKey))

	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(authRoute)
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}






func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" && password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
