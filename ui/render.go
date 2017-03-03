package ui

import "github.com/jroimartin/gocui"

const (
	// SearchView is the name of the Search view
	SearchView = "SearchInput"
	// StatusView is the name of the Status view
	StatusView = "Status"
)

// RenderLayout ....
func RenderLayout(g *gocui.Gui) error {
	if err := RenderSearch(g, SearchView); err != nil {
		return err
	}

	if err := RenderStatus(g, StatusView); err != nil {
		return err
	}

	return nil
}

// RenderSearch ..
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

// RenderStatus ...
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
