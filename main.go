package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

const (
	searchView = "SearchInput"
	statusView = "Status"
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
		status, err := g.View(statusView)
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
	if err := renderSearch(g); err != nil {
		return err
	}

	if err := renderStatus(g); err != nil {
		return err
	}

	return nil
}

func renderSearch(g *gocui.Gui) error {
	maxX, _ := g.Size()

	if v, err := g.SetView(searchView, 0, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Wrap = false
		if _, err := g.SetCurrentView(searchView); err != nil {
			return err
		}
	}

	return nil
}

func renderStatus(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(statusView, 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = false
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

	if err := g.SetKeybinding(searchView, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding(searchView, gocui.KeyEnter, gocui.ModNone, search); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	g.Close()
}
