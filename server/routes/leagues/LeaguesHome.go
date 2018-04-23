package leagues

import (
	"cdm/server/muxrouter"
)

func LeaguesRoutes(mounter string) {
	r := &muxrouter.Routes{
		muxrouter.Route{
			Name:        "All",
			Method:      "GET",
			Path:        mounter,
			HandlerFunc: GetAllLeagues,
			Protected:   false,
		},
		muxrouter.Route{
			Name:        "OneById",
			Method:      "GET",
			Path:        mounter + "/:id",
			HandlerFunc: GetOneById,
			Protected:   false,
		},
		muxrouter.Route{
			Name:   "New",
			Method: "POST",
			Path : mounter,
			HandlerFunc: CreateLeague,
			Protected: true,
		},
		muxrouter.Route{
			Name:        "Update",
			Method:      "PUT",
			Path:        mounter + "/:id",
			HandlerFunc: UpdateOne,
			Protected:   false,
		},
	}
	muxrouter.GetRouter().AddRoute(r)
}
