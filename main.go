package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"os"
	"sync"

	"github.com/leonlarsson/go-image-gen/engine"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

var iterations int

// init is called before main
func init() {

	// Parse command line arguments
	flag.IntVar(&iterations, "iterations", 1, "Number of iterations")
	flag.Parse()

	// Clean up old renders
	resetRendersFolder()
}

// generateImage generates an image and saves it to disk
func generateImage(width, height, identifier int, wg *sync.WaitGroup) {
	defer wg.Done()
	scene := engine.NewScene(width, height)
	scene.EachPixel(func(x, y int) color.RGBA {
		return color.RGBA{
			uint8(x * 255 / width),
			uint8(y * 255 / height),
			100,
			255,
		}
	})
	fileName := fmt.Sprintf("%d.png", identifier+1)
	scene.Save("./renders/" + fileName)
}

// resetRendersFolder removes the renders folder and creates a new one
func resetRendersFolder() {
	os.RemoveAll("./renders")
	os.MkdirAll("./renders", os.ModePerm)
}

func main() {

	c := canvas.New(1200, 750)

	ctx := canvas.NewContext(c)

	bgFile, err := os.Open("assets/images/BF2042/BF2042_IMAGE_BG_0.png")
	if err != nil {
		panic(err)
	}

	gridFile, err := os.Open("assets/images/Skeleton_BGs/Regular.png")
	if err != nil {
		panic(err)
	}

	gameLogoFile, err := os.Open("assets/images/BF2042/BF2042_LOGO_BG.png")
	if err != nil {
		panic(err)
	}

	bgImg, err := png.Decode(bgFile)
	if err != nil {
		panic(err)
	}

	gridImg, err := png.Decode(gridFile)
	if err != nil {
		panic(err)
	}

	gameLogoImg, err := png.Decode(gameLogoFile)
	if err != nil {
		panic(err)
	}

	ctx.DrawImage(0, 0, bgImg, canvas.DPMM(1))
	ctx.DrawImage(0, 0, gridImg, canvas.DPMM(1))
	ctx.DrawImage(0, 0, gameLogoImg, canvas.DPMM(1))

	font := canvas.NewFontFamily("Roboto")
	font.LoadFontFile("assets/fonts/Roboto-Medium.ttf", canvas.FontRegular)
	face := font.Face(35*10, canvas.White, canvas.FontRegular, canvas.FontNormal)

	ctx.DrawText(57, 180, canvas.NewTextLine(face, "Lorem", canvas.Left))

	if err := renderers.Write("renders/test.png", c, canvas.DPMM(1)); err != nil {
		panic(err)
	}

	// startTime := time.Now()

	// var wg sync.WaitGroup

	// // Run one goroutine per iteration
	// for i := 0; i < iterations; i++ {
	// 	wg.Add(1)
	// 	go generateImage(rand.IntN(500)+1, rand.IntN(500)+1, i, &wg)
	// }

	// // Wait for all goroutines to finish
	// wg.Wait()

	// fmt.Printf("Generated %d images in %s\n", iterations, time.Since(startTime))
}
