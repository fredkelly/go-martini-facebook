# go-martini-facebook
Boilerplate API server for applications authenticating via Facebook.

## Motivation
To create simple boilerplate code for building a golang/martini API using Facebook for authentication.

## Basic usage
- User connects to your app (e.g. SPA or mobile)
- You get an access token from Facebook via the standard [OAuth flow](https://developers.facebook.com/docs/reference/dialogs/oauth)
- That token is then passed to subsequent requests to YOUR API (e.g. `Authorization: {MY_TOKEN}`).
- The server then stores a corresponding user record (see `models/user.go`).

## Built with:
- [go-martini/martini](https://github.com/go-martini/martini)
- [huandu/facebook](https://github.com/huandu/facebook) 
- [coopernurse/gorp](https://github.com/coopernurse/gorp)
- other great stuff.

There's also a bunch of other cruft for receiving/rendering JSON - feel free to add/remove as required!

## TODO:
Make it better; right now everything is bunch up, would be nice to add some modularity - I'm a n00b Gopher so could use all the help I can get!
