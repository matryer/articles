package main

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/machinebox/sdk-go/facebox"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("usage: anon <image>")
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	log.Println("Detecting faces...")
	fb := facebox.New("http://localhost:8080")
	faces, err := fb.Check(f)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		log.Fatalln(err)
	}
	srcImage, _, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	dstImage := anonymise(srcImage, faces)
	filename = filepath.Base(filename)
	ext := filepath.Ext(filename)
	dstFilename := filename[:len(filename)-len(ext)] + "-anon" + ext
	dstFile, err := os.Create(dstFilename)
	if err != nil {
		log.Fatalln(err)
	}
	defer dstFile.Close()
	log.Println("Saving image to " + dstFilename + "...")
	err = jpeg.Encode(dstFile, dstImage, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Done.")
}

func anonymise(src image.Image, faces []facebox.Face) image.Image {
	dstImage := image.NewRGBA(src.Bounds())
	draw.Draw(dstImage, src.Bounds(), src, image.ZP, draw.Src)
	for _, face := range faces {
		faceRect := image.Rect(face.Rect.Left, face.Rect.Top, face.Rect.Left+face.Rect.Width, face.Rect.Top+face.Rect.Height)
		facePos := image.Pt(face.Rect.Left, face.Rect.Top)
		draw.Draw(dstImage, faceRect, &image.Uniform{color.Black}, facePos, draw.Src)
	}
	return dstImage
}
