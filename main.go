package main

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		ws := c.Query("w")
		hs := c.Query("h")

		width, err := strconv.Atoi(ws)
		height, err := strconv.Atoi(hs)
		if err != nil {
			log.Fatal(err)
		}
		wt := uint(width)
		ht := uint(height)

		c.Stream(func(w io.Writer) bool {
			img, err := loadImg("1.png")
			if err != nil {
				log.Fatal("出错了！")
			}
			m := resize.Resize(wt, ht, img, resize.Lanczos3)
			png.Encode(w, m)

			return false
		})
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
}

func loadImg(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}
