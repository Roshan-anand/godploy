package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func DockerRoutes(mux *http.ServeMux, client *client.Client) {
	mux.HandleFunc("GET /container/list/all", func(w http.ResponseWriter, r *http.Request) {
		cntrs, err := client.ContainerList(context.Background(), container.ListOptions{All: true})
		if err != nil {
			http.Error(w, "failed to list containers", http.StatusInternalServerError)
			return
		}

		info := "List of all containers:\n"
		for _, c := range cntrs {
			info += fmt.Sprintf("%s %s %s \n", c.Image, c.ID, c.Status)
		}
		w.Write([]byte(info))
	})

	mux.HandleFunc("GET /image/list/", func(w http.ResponseWriter, r *http.Request) {
		imgs, err := client.ImageList(context.Background(), image.ListOptions{All: true})
		if err != nil {
			http.Error(w, "failed to list images", http.StatusInternalServerError)
			return
		}

		info := "List of all images:\n"
		for _, img := range imgs {
			info += fmt.Sprintf("%s %s \n", img.ID, img.RepoTags)
		}
		w.Write([]byte(info))
	})
}
