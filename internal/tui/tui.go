package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func RunTUI() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}

	defer g.Close()

	g.SetManagerFunc(mainLayout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err = g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func mainLayout(g *gocui.Gui) error {
	if err := headerView(g); err != nil {
		return err
	}

	if err := footerView(g); err != nil {
		return err
	}

	return nil
}

func headerView(g *gocui.Gui) error {
	maxX, _ := g.Size()
	if v, err := g.SetView("header", -1, -1, maxX, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FgColor = gocui.ColorWhite
		v.BgColor = gocui.ColorCyan

		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false
		v.Overwrite = true

		fmt.Fprintln(v, " klingo")
	}

	return nil
}

func footerView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("footer", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FgColor = gocui.ColorWhite
		v.BgColor = gocui.ColorCyan

		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false
		v.Overwrite = true

		// TODO: Implement context configuration
		fmt.Fprintln(v, " Context: default")
	}

	return nil
}
