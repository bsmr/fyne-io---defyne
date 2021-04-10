package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xWidget "fyne.io/x/fyne/widget"
	"github.com/fyne-io/terminal"
)

func (d *defyne) makeEditorPanel() fyne.CanvasObject {
	welcome := widget.NewLabel("Welcome to DEFyne, the Fyne IDE.\n\nChoose a project file from the list.")
	d.fileTabs = container.NewDocTabs(
		container.NewTabItemWithIcon("Welcome", theme.HomeIcon(),
			container.NewCenter(welcome)))

	return container.NewMax(d.fileTabs)
}

func (d *defyne) makeFilesPanel() fyne.CanvasObject {
	files := xWidget.NewFileTree(d.projectRoot)
	files.Filter = filterHidden()
	files.Sorter = func(u1, u2 fyne.URI) bool {
		return u1.String() < u2.String() // Sort alphabetically
	}

	files.OnSelected = func(uid widget.TreeNodeID) {
		u, err := storage.ParseURI(uid)
		if err != nil {
			dialog.ShowError(err, d.win)
		}

		d.openEditor(u)
	}
	return files
}

func (d *defyne) makeTerminalPanel() fyne.CanvasObject {
	term := terminal.New()
	term.SetStartDir(d.projectRoot.Path())
	go term.RunLocalShell()

	return term
}

type filter struct{}

func (f *filter) Matches(u fyne.URI) bool {
	return u.Name()[0] != '.'
}

func filterHidden() storage.FileFilter {
	return &filter{}
}