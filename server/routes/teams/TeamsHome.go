package teams

import "cdm/server/muxrouter"

func Routes(mounter string) {
	r := &muxrouter.Routes{
		muxrouter.Route{
			Name:        "All",
			Method:      "GET",
			Path:        mounter,
			HandlerFunc: GetAllTeams,
			Protected:   false,
		},
	}
	muxrouter.GetRouter().AddRoute(r)
}
