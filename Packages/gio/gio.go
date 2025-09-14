package main

//Сделано на основе gio и учебника на jonegil.github.io

import (
	"errors"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

func draw(w *app.Window) error {
	var ops op.Ops

	var firstButton widget.Clickable
	var fieldValueRing widget.Editor

	var secondButton widget.Clickable
	var fieldValueVint widget.Editor
	var fieldHeight widget.Editor
	var fieldDiametr widget.Editor

	type Vint struct {
		Height  float64
		Diametr float64
		Value   float64
	}
	var Vi Vint

	var form func(f float64) string = func(f float64) string { return strconv.FormatFloat(f, 'f', 2, 64) }

	var valout float64

	th := material.NewTheme()

	for {
		e := w.Event()

		switch typ := e.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)

			layout.Flex{
				Axis: layout.Vertical,

				Spacing: layout.SpaceBetween,
			}.Layout(
				gtx,
				layout.Rigid(func(gtx C) D {
					theme := material.NewTheme()
					valv := (valout * 3.14159) / 3.2
					text1 := material.H5(theme, "Кольцо "+form(valv))
					text1.Alignment = text.Middle
					return text1.Layout(gtx)
				}),

				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(350),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
					}

					ed := material.Editor(th, &fieldValueRing, "Значение")

					fieldValueRing.SingleLine = true
					fieldValueRing.Alignment = text.Middle

					val := fieldValueRing.Text()
					val = strings.TrimSpace(val)
					valout, _ = strconv.ParseFloat(val, 64)

					border := widget.Border{
						Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
						CornerRadius: unit.Dp(3),
						Width:        unit.Dp(2),
					}
					return margins.Layout(gtx,
						func(gtx C) D {
							return border.Layout(gtx, ed.Layout)
						},
					)
				}),
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),
					}

					return margins.Layout(
						gtx,
						func(gtx C) D {
							fbtn := material.Button(th, &firstButton, "Кольцо")

							return fbtn.Layout(gtx)
						},
					)
				}),

				//Второй блок

				layout.Rigid(func(gtx C) D {
					theme := material.NewTheme()
					vall := (3.14159 * Vi.Diametr)
					valc := vall / Vi.Value
					valo := (60 / (vall / Vi.Value)) * 100 / 32
					valt := (valc / 1.6) * Vi.Height
					valg := valt / 4
					valf := valt / 3
					str := "Винт: l - " + form(vall) +
						" c " + form(valc) +
						" o " + form(valo) +
						" g " + form(valg) +
						" f " + form(valf)
					text1 := material.H5(theme, str)
					text1.Alignment = text.Middle
					return text1.Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(350),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
					}

					ed := material.Editor(th, &fieldHeight, "Высота")

					fieldHeight.SingleLine = true
					fieldHeight.Alignment = text.Middle

					val := fieldHeight.Text()
					val = strings.TrimSpace(val)
					Vi.Height, _ = strconv.ParseFloat(val, 64)

					border := widget.Border{
						Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
						CornerRadius: unit.Dp(3),
						Width:        unit.Dp(2),
					}
					return margins.Layout(gtx,
						func(gtx C) D {
							return border.Layout(gtx, ed.Layout)
						},
					)
				}),
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(350),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
					}

					ed := material.Editor(th, &fieldDiametr, "Диаметр")

					fieldDiametr.SingleLine = true
					fieldDiametr.Alignment = text.Middle

					val := fieldDiametr.Text()
					val = strings.TrimSpace(val)
					Vi.Diametr, _ = strconv.ParseFloat(val, 64)

					border := widget.Border{
						Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
						CornerRadius: unit.Dp(3),
						Width:        unit.Dp(2),
					}

					return margins.Layout(gtx,
						func(gtx C) D {
							return border.Layout(gtx, ed.Layout)
						},
					)
				}),
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(10),
						Right:  unit.Dp(350),
						Bottom: unit.Dp(10),
						Left:   unit.Dp(10),
					}

					ed := material.Editor(th, &fieldValueVint, "Значение")

					fieldValueVint.SingleLine = true
					fieldValueVint.Alignment = text.Middle

					val := fieldValueVint.Text()
					val = strings.TrimSpace(val)
					Vi.Value, _ = strconv.ParseFloat(val, 64)

					border := widget.Border{
						Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
						CornerRadius: unit.Dp(3),
						Width:        unit.Dp(2),
					}

					return margins.Layout(gtx,
						func(gtx C) D {
							return border.Layout(gtx, ed.Layout)
						},
					)
				}),
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(25),
						Left:   unit.Dp(25),
					}

					return margins.Layout(
						gtx,
						func(gtx C) D {
							sbtn := material.Button(th, &secondButton, "Винт")

							return sbtn.Layout(gtx)
						},
					)
				}),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(50)}.Layout,
				),
			)
			typ.Frame(gtx.Ops)
		case app.DestroyEvent:
			return errors.New("Window is destroy")
		}
	}
}

func main() {
	//Надо будет пропробовать не запускать в горутине и посмотреть как оно заблокирует вывод.
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Программа"))
		w.Option(app.Size(unit.Dp(500), unit.Dp(800)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
