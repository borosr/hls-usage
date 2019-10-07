package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"strconv"
)

const (
	mediaRoot     = "assets/media"
	basicM3U8Name = "index.m3u8"
)

func main() {
	log.Println("Server starting on 8000")
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	r.Get("/media/{mId:[0-9]+}/stream/", handleStream)
	r.Get("/media/{mId:[0-9]+}/stream/{segName:index[0-9]+.ts}", handleStream)
	_ = http.ListenAndServe(":8000", r)
}

func handleStream(w http.ResponseWriter, r *http.Request) {
	mid := chi.URLParam(r, "mId")
	mID, err := strconv.Atoi(mid)
	if err != nil {
		log.Println("mid not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	segName := chi.URLParam(r, "segName")
	if segName == "" {
		useM3U8(w, r, mID)
	} else {
		useHlsTs(w, r, segName, mID)
	}
}

func useHlsTs(w http.ResponseWriter, r *http.Request, segname string, mId int) {
	mediaFile := fmt.Sprintf("%s/hls/%s", getMediaBase(mId), segname)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "video/MP2T")
}

func useM3U8(w http.ResponseWriter, r *http.Request, mId int) {
	mediaFile := fmt.Sprintf("%s/hls/%s", getMediaBase(mId), basicM3U8Name)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "application/x-mpegURL")

}

func getMediaBase(mId int) string {
	return fmt.Sprintf("%s/%d", mediaRoot, mId)
}
