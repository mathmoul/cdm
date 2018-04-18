package authentication

import (
	"cdm/server/muxrouter"
)

// all routes with match /api/auth

/*
Authentication function
*/
func Authentication(mounter string) {
	r := &muxrouter.Routes{
		muxrouter.Route{
			Name:        "login",
			Method:      "POST",
			Path:        mounter,
			HandlerFunc: Auth,
			Protected:   false,
		},
		muxrouter.Route{
			Name:        "signup",
			Method:      "POST",
			Path:        "/signup",
			HandlerFunc: Signup,
			Protected:   false,
		},
		muxrouter.Route{
			Name:        "confirm",
			Method:      "POST",
			Path:        mounter + "/confirmation",
			HandlerFunc: Confirmation,
			Protected:   false,
		},
	}
	muxrouter.GetRouter().AddRoute(r)
}
