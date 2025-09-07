package main

//Сделано на основе gio и учебника на jonegil.github.io

import (
	"errors"
	"log"
	"os"

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
	var secondButton widget.Clickable

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
				layout.Rigid(
					func(gtx C) D {
						theme := material.NewTheme()
						text := material.H5(theme, "some walue")
						return text.Layout(gtx)
					},
				),
				layout.Rigid(
					func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}

						return margins.Layout(
							gtx,
							func(gtx C) D {
								fbtn := material.Button(th, &firstButton, "Ring")

								return fbtn.Layout(gtx)
							},
						)
					},
				),
				layout.Rigid(
					func(gtx C) D {
						theme := material.NewTheme()
						text1 := material.H5(theme, "some walue")
						text1.Alignment = text.Middle
						return text1.Layout(gtx)
					},
				),
				layout.Rigid(
					func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(25),
							Left:   unit.Dp(25),
						}

						return margins.Layout(
							gtx,
							func(gtx C) D {
								sbtn := material.Button(th, &secondButton, "Vint")

								return sbtn.Layout(gtx)
							},
						)
					},
				),
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
		w.Option(app.Title("Path"))
		w.Option(app.Size(unit.Dp(500), unit.Dp(800)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
