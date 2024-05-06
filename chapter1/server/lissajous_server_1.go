package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette []color.Color

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.Atoi(r.FormValue("cycles"))
	if err != nil {
		cycles = 5
	}
	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		size = 100
	}
	lissajous(w, cycles, size)
}

func lissajous(out io.Writer, cycles, size int) {
	const (
		res     = 0.001
		nframes = 64
		delay   = 8
	)

	palette = append(palette, color.Black)
	palette = append(palette, color.RGBA{0x00, 0xaa, 0x00, 0xff})

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
