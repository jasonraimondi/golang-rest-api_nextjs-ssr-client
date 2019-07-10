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

	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	authRoute := middleware.JWTWithConfig(config)

	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(authRoute)
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// diminish brute force attempts
	time.Sleep(500 * time.Millisecond)

	// Throws unauthorized error
	if username != "jon" && password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		username,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(*JwtSecureKey))
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
