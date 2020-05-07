package main

import (
	"github.com/NovikovRoman/gorough"
	"github.com/ajstarks/svgo"
	"os"
)

func curves() (err error) {
	curve := gorough.NewCurve([]gorough.Point{
		{X: 10, Y: 10},
		{X: 200, Y: 10},
		{X: 100, Y: 100},
		{X: 300, Y: 100},
		{X: 60, Y: 200},
	}, nil)

	curve2 := gorough.NewCurve([]gorough.Point{
		{X: 50, Y: 20},
		{X: 30, Y: 200},
		{X: 300, Y: 180},
		{X: 280, Y: 30},
		{X: 170, Y: 180},
	}, &gorough.EllipseOptions{
		Styles: &gorough.Styles{
			Stroke: "#ff0000",
		},
	})

	var f *os.File
	if f, err = os.Create("curves.svg"); err != nil {
		return
	}
	defer func() {
		if derr := f.Close(); derr != nil {
			err = derr
		}
	}()

	width := 320
	height := 220
	canvas := svg.New(f)
	canvas.Start(width, height)
	gorough.DrawSVG(canvas, curve, nil)
	gorough.DrawSVG(canvas, curve2, nil)
	canvas.End()
	return
}
