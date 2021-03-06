package authentication

import (
	"cdm/server/muxrouter"
	"net/http"
)

// all routes with match /api/auth

/*
Authentication function
*/

func route(name string, method string, path string, f http.HandlerFunc) muxrouter.Route {
	return muxrouter.Route{
		Name:        name,
		Method:      method,
		Path:        path,
		HandlerFunc: f,
		Protected:   false,
	}
}

/*
Authentication function
*/
func Routes(mounter string) {
	r := &muxrouter.Routes{
		muxrouter.Route{
			Name:        "login",
			Method:      "POST",
			Path:        mounter,
			HandlerFunc: Authentication,
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
		muxrouter.Route{
			Name:        "resetPasswordrequest",
			Method:      "POST",
			Path:        mounter + "/reset_password_request",
			HandlerFunc: ResetPasswordRequest,
			Protected:   false,
		},
		muxrouter.Route{
			Name:        "validateToken",
			Method:      "POST",
			Path:        mounter + "/validate_token",
			HandlerFunc: ValidateToken,
			Protected:   false,
		},
		muxrouter.Route{
			Name:        "resetpassword",
			Method:      "POST",
			Path:        mounter + "/reset_password",
			HandlerFunc: ResetPassword,
			Protected:   false,
		},
	}
	muxrouter.GetRouter().AddRoute(r)
}
