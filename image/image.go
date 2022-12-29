package image

import (
	"log"
	"os"

	"github.com/fogleman/gg"
	"github.com/wcharczuk/go-chart/v2"
)

var directory = os.Getenv("DIRECTORY")

const (
	chartWidth  = 512
	chartHeight = 512

	chartX = 50
	chartY = 0
)

func GenerateChart(username string) {
	pie := chart.PieChart{
		Width:  chartWidth,
		Height: chartHeight,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	f, _ := os.Create(username + ".png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}

func CreateUser(username string) error {
	graph, err := gg.LoadPNG(username + ".png")
	if err != nil {
		return err
	}

	dc := gg.NewContext(512, 512+100)

	// background colour
	dc.SetHexColor("#FFFFFF")
	dc.Clear()

	// draw graph
	dc.DrawImage(graph, 0, 99)

	// draw text
	dc.SetHexColor("#000000")
	dc.LoadFontFace(directory+"fonts/NotoSansDisplay-Regular.ttf", 24)
	dc.DrawStringAnchored(username, 512/2, 30, 0.5, 0)

	err = dc.SavePNG(username + "-x.png")
	if err != nil {
		return err
	}

	return nil
}

func Start() {
	username := "output"

	GenerateChart(username)
	err := CreateUser(username)
	log.Print(err)
}
