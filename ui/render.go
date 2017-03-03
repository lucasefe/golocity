package ui

import "github.com/jroimartin/gocui"

func RenderSearch(g *gocui.Gui, name string) error {
	maxX, _ := g.Size()

	if v, err := g.SetView(name, 0, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Wrap = false
		if _, err := g.SetCurrentView(name); err != nil {
			return err
		}
	}

	return nil
}

func RenderStatus(g *gocui.Gui, name string) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(name, 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = false
	}

	return nil
}
