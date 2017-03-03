package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/lucasefe/go-nv/ui"
)

const (
	// SearchView is the name of the Search view
	SearchView = "SearchInput"
	// StatusView is the name of the Status view
	StatusView = "Status"
)

var vbuf, buf string

func quit(g *gocui.Gui, v *gocui.View) error {
	vbuf = v.ViewBuffer()
	buf = v.Buffer()
	return gocui.ErrQuit
}

func search(g *gocui.Gui, v *gocui.View) error {
	query := v.Buffer()
	v.Clear()
	v.SetCursor(0, 0)

	g.Execute(func(g *gocui.Gui) error {
		status, err := g.View(StatusView)
		if err != nil {
			return err
		}

		status.Clear()
		status.Write([]byte(fmt.Sprintf("=> %+v\n", query)))
		return nil
	})

	return nil
}

func layout(g *gocui.Gui) error {
	if err := ui.RenderSearch(g, SearchView); err != nil {
		return err
	}

	if err := ui.RenderStatus(g, StatusView); err != nil {
		return err
	}

	return nil
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding(SearchView, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding(SearchView, gocui.KeyEnter, gocui.ModNone, search); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	g.Close()
}
