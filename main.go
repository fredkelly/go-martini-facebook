package main

import (
	"log"
	"os"
	//"fmt"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/encoder"

	fb "github.com/huandu/facebook"

	"./models"
)

func main() {
	// TODO move to own handler
	// create a global App var to hold your app id and secret.
	globalApp := fb.New(os.Getenv("FACEBOOK_APP_ID"), os.Getenv("FACEBOOK_SECRET"))
	globalApp.RedirectUri = "http://localhost:3000/auth/facebook/callback" // TODO

	// https://developers.facebook.com/docs/graph-api/securing-requests
	globalApp.EnableAppsecretProof = true

	// instantiate Martini
	m := martini.Classic()
	m.Use(martini.Logger())

	// CORS
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Authentication
	m.Use(func(c martini.Context, res http.ResponseWriter, req *http.Request) {
		accessToken := req.Header.Get("Authorization") // TODO expect "Bearer: {TOKEN}" format?
		session := globalApp.Session(accessToken)

		user := &models.User{}
		attrs, err := session.Get("/me", nil)

		if err != nil {
			// err can be an facebook API error.
			// if so, the Error struct contains error details.
			if e, ok := err.(*fb.Error); ok {
				log.Printf("Facebook error. [message:%v] [type:%v] [code:%v] [subcode:%v]", e.Message, e.Type, e.Code, e.ErrorSubcode)
			}

			res.WriteHeader(http.StatusUnauthorized)
		} else {
			err := attrs.Decode(user)
			if err != nil {
				log.Printf("Error decoding user from Facebook: %s", err)
			} else {
				models.FindOrCreateUser(user)
			}
			c.Map(user)
		}
	})

	// Encoding
	m.Use(func(c martini.Context, w http.ResponseWriter, r *http.Request) {
		// Use indentations. &pretty=1
		pretty, _ := strconv.ParseBool(r.FormValue("pretty"))
		// Use null instead of empty object for json &null=1
		null, _ := strconv.ParseBool(r.FormValue("null"))
		// JSON no matter what
		c.MapTo(encoder.JsonEncoder{PrettyPrint: pretty, PrintNull: null}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Get("/", func(user *models.User) string {
		return "Hello " + user.FirstName
	})

	m.Run()
}
