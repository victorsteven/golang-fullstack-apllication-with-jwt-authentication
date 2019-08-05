package routes

import (
	"golang_api_fullstack/api/controllers"
	"net/http"
)

var postsRoutes = []Route{
	Route{
		Uri:     "/posts",
		Method:  http.MethodGet,
		Handler: controllers.GetPosts,
	},
	Route{
		Uri:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
	},
	Route{
		Uri:     "/posts/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetPost,
	},
	Route{
		Uri:     "/posts/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdatePost,
	},
	Route{
		Uri:     "/posts/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeletePost,
	},
}
