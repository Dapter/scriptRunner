package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Dapter/scriptRunner/internal/interfaces"
	"image/color"
)

type desktop struct {
	app interfaces.App
}

func New(app interfaces.App) *desktop {
	return &desktop{app}
}

func (d desktop) Start() {
	a := app.New()
	win := a.NewWindow("Server Mon")
	c1 := canvas.NewText("Canvas Object 1", color.White)
	c2 := canvas.NewText("Canvas Object 2                                      2", color.White)
	n1 := canvas.NewText("Canvas Object 3", color.White)
	n2 := canvas.NewText("Canvas Object 4                                   2", color.White)

	row := container.New(layout.NewGridLayoutWithColumns(2), c1, c2)
	row.Resize(fyne.NewSize(20, 20))

	topLevelLayOutContent := container.New(layout.NewVBoxLayout(),
		container.New(layout.NewMaxLayout(), row),
		container.New(layout.NewMaxLayout(),
			container.New(layout.NewGridLayoutWithColumns(2), n1, n2)),
	)

	win.SetContent(topLevelLayOutContent)
	win.Resize(fyne.NewSize(float32(400), float32(30)))
	win.ShowAndRun()

	//commands, err := d.app.GetList()

	//if err != nil {
	//	panic(err)
	//}

	//a := app.New()
	//w := a.NewWindow("Command runner")
	//
	//w.Resize(fyne.NewSize(460, 460))
	//
	//w.SetContent(
	//	makeTable(
	//		[]string{"Name", "Commands", "Action"},
	//		[][]string{{"1", "21111111111111111111111111111111111111111111111111111", "3"}, {"4", "5", "6"}},
	//	),
	//)

	//for label, command := range commands {
	//	grid.Add(widget.NewLabel(label))
	//	grid.Add(widget.NewLabel(command.String()))
	//	(func(label string) {
	//		grid.Add(widget.NewButton("Run", func() {
	//			fmt.Println(label)
	//		}))
	//	})(label)
	//}
	//
	//grid.Resize(fyne.NewSize(20, 20))
	//
	//w.SetContent(grid)

	//w.ShowAndRun()
}

func makeTable(headings []string, rows [][]string) *fyne.Container {

	columns := rowsToColumns(headings, rows)

	objects := make([]fyne.CanvasObject, len(columns))
	for k, col := range columns {
		box := container.NewVBox(widget.NewLabelWithStyle(headings[k], fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
		//container.NewMax
		box.Resize(fyne.NewSize(20, 20))
		for _, val := range col {
			box.Add(widget.NewLabel(val))
		}
		objects[k] = box
	}
	return container.New(layout.NewGridLayoutWithColumns(len(columns)), objects...)
}

func rowsToColumns(headings []string, rows [][]string) [][]string {
	columns := make([][]string, len(headings))
	for _, row := range rows {
		for colK := range row {
			columns[colK] = append(columns[colK], row[colK])
		}
	}
	return columns
}
