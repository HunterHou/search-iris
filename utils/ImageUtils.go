package utils

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

func ImageToString(path string) string {
	if !ExistsFiles(path) {
		return ""
	}
	file, _ := os.Open(path)
	defer file.Close()
	sourceBuffer := make([]byte, 500000)
	n, _ := file.Read(sourceBuffer)
	return base64.StdEncoding.EncodeToString(sourceBuffer[:n])
}

func ImageToPng(src string) error {
	des := GetPng(src, "png")
	fin, _ := os.Open(src)
	fin2, _ := os.Open(src)
	defer fin.Close()
	defer fin2.Close()
	fout, _ := os.Create(des)
	defer fout.Close()
	config, _, _ := image.DecodeConfig(fin2)
	srcImage, fm, err := image.Decode(fin)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	height := config.Height
	width := config.Width
	left := int(float64(0.53) * float64(width))
	switch fm {
	case "jpeg":
		rgbImg := srcImage.(*image.YCbCr)
		subImg := rgbImg.SubImage(image.Rect(left, 0, width, height)).(*image.YCbCr)
		png.Encode(fout, subImg)
	case "png":
		switch srcImage.(type) {
		case *image.NRGBA:
			img := srcImage.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(left, 0, width, height)).(*image.NRGBA)
			return png.Encode(fout, subImg)
		case *image.RGBA:
			img := srcImage.(*image.RGBA)
			subImg := img.SubImage(image.Rect(left, 0, width, height)).(*image.RGBA)
			return png.Encode(fout, subImg)
		}
	}
	return nil

}
