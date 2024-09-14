package shapes

import (
	"design_patterns/chapter5/behavioralPatterns/strategy"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

type ImageSquare struct {
	strategy.PrintOutput
	DestinationFilePath string
}

func (t *ImageSquare) Print() error {
	width := 400
	height := 600
	origin := image.Point{0, 0}
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	// تنظیم رنگ پس‌زمینه (رنگ خاکستری روشن)
	bgColor := color.RGBA{R: 50, G: 60, B: 70, A: 250} // مقدار A باید 255 باشد تا غیرشفاف باشد
	draw.Draw(bgImage, bgImage.Bounds(), &image.Uniform{C: bgColor}, origin, draw.Src)
	// تنظیم ویژگی‌های مربع
	squareWidth := 200
	squareHeight := 200
	squareColor := color.RGBA{R: 255, G: 0, B: 0, A: 1} // قرمز رنگی و غیرشفاف
	// تعریف موقعیت مربع
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	// ایجاد تصویر مربع و رسم آن
	squareImg := image.NewRGBA(square)
	draw.Draw(squareImg, squareImg.Bounds(), &image.Uniform{C: squareColor}, image.Point{}, draw.Src)
	draw.Draw(bgImage, squareImg.Bounds(), squareImg, image.Point{}, draw.Over)

	// ایجاد و ذخیره تصویر
	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("خطا در باز کردن فایل تصویر برای نوشتن: %w", err)
	}
	defer w.Close()

	quality := &jpeg.Options{Quality: 75}
	if err = jpeg.Encode(w, bgImage, quality); err != nil {
		return fmt.Errorf("خطا در نوشتن تصویر به دیسک: %w", err)
	}

	return nil
}
