package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/lucasefe/go-nv/ui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	g.Cursor = true
	g.Mouse = false

	g.SetManagerFunc(ui.RenderLayout)

	if err := g.SetKeybinding(ui.SearchView, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding(ui.SearchView, gocui.KeyEnter, gocui.ModNone, search); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	g.Close()
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func search(g *gocui.Gui, v *gocui.View) error {
	query := v.Buffer()
	v.Clear()
	v.SetCursor(0, 0)

	g.Execute(func(g *gocui.Gui) error {
		status, err := g.View(ui.StatusView)
		if err != nil {
			return err
		}

		status.Clear()
		status.Write([]byte(fmt.Sprintf("=> %+v\n", query)))
		return nil
	})

	return nil
}
