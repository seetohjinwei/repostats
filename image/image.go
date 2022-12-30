package image

import (
	"errors"
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/seetohjinwei/repostats/image/circle"
	"github.com/seetohjinwei/repostats/models"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

const (
	backgroundRound  = 6
	backgroundColour = "#171214"

	imageWidth  = 380
	imageHeight = 180

	marginX = 10
	marginY = 10

	textAlign  = "start"
	textColour = "#eeeeee"

	titleSize      = 20
	titleGap       = 16
	titleMaxLength = 20
	titleX         = marginX + 10
	titleY         = marginY + titleSize

	listSize      = 16
	listGap       = 16
	listMaxLength = 16
	listX         = titleX + 4
	listY         = titleY + titleSize*2

	pieRadius float64 = imageHeight/2 - marginY
	pieX      float64 = imageWidth - pieRadius - marginX
	pieY      float64 = pieRadius + marginY

	watermarkSize = 12
	watermark     = "by RepoStats"
	watermarkX    = marginX + 14
	watermarkY    = imageHeight - marginY
)

type item struct {
	name   string
	ratio  float64
	colour string
}

// Colour scheme from: https://www.schemecolor.com/orange-green-blue-pie-chart.php
var colours = []string{"#F47A1F", "#FDBB2F", "#377B2B", "#7AC142", "#007CC3", "#00529B"}

func createAllValues(typeData map[string]models.TypeData) []item {
	data := maps.Values(typeData)
	slices.SortFunc(data, models.MoreTypeData)

	var totalFiles float64 = 0
	for _, d := range data {
		totalFiles += float64(d.FileCount)
	}

	items := make([]item, len(data))

	for i, d := range data {
		ratio := float64(d.FileCount) / totalFiles * 100
		var colour string
		if i >= 6 {
			colour = "white"
		} else {
			colour = colours[i]
		}
		item := item{
			name:   d.Type,
			ratio:  ratio,
			colour: colour,
		}
		items[i] = item
	}

	return items
}

func createValues(typeData map[string]models.TypeData) []item {
	items := make([]item, 6)

	all := createAllValues(typeData)
	if len(typeData) <= 6 {
		return all
	}

	others := 100.0
	for i := 0; i < 5; i++ {
		items[i] = all[i]
		others -= all[i].ratio
	}

	items[5] = item{
		name:   "others",
		ratio:  others,
		colour: colours[5],
	}

	return items
}

func CreateUserSvg(w io.Writer, username string, typeData map[string]models.TypeData) error {
	values := createValues(typeData)

	return createUserSvg(w, username, values)
}

var ErrNot100 = errors.New("values do not add to 100")

func createUserSvg(w io.Writer, username string, values []item) error {
	canvas := svg.New(w)
	canvas.Start(imageWidth, imageHeight)

	// background
	s := fmt.Sprintf("fill: %s", backgroundColour)
	canvas.Roundrect(0, 0, imageWidth, imageHeight, backgroundRound, backgroundRound, s)

	// username
	wrappedUsername := wrapText(username)
	titleYPad := 0
	if len(wrappedUsername) == 1 {
		titleYPad = titleSize / 2
	}
	canvas.Textlines(titleX, titleY+titleYPad, wrappedUsername, titleSize, titleGap, textColour, textAlign)

	// languages list
	canvas.Textlines(listX, listY, languagesList(values), listSize, listGap, textColour, textAlign)

	// pie chart
	circle := circle.New(pieX, pieY, pieRadius)
	acc := 0.0
	for _, value := range values {
		start := acc / 100
		acc += value.ratio
		end := acc / 100

		s := fmt.Sprintf("fill: %s", value.colour)
		canvas.Path(circle.NewSlice(start, end), s)
	}

	// watermark
	s = fmt.Sprintf("fill: %s; font-size: %dpx", textColour, watermarkSize)
	canvas.Text(watermarkX, watermarkY, watermark, s)

	canvas.End()

	// signature
	fmt.Fprintf(w, "<!-- Stats by RepoStats -->\n")
	fmt.Fprintf(w, "<!-- https://repostats.jinwei.dev -->\n")

	const margin = 1
	if 100 < acc-margin || 100 > acc+margin {
		return ErrNot100
	}
	return nil
}

func wrapText(text string) []string {
	l := len(text)

	lines := []string{}

	for i := 0; i < l; i = i + titleMaxLength {
		sub := text[i:min(i+titleMaxLength, l)]
		lines = append(lines, sub)
	}

	return lines
}

func languagesList(items []item) []string {
	truncate := func(x string) string {
		if len(x) <= listMaxLength {
			return x
		}

		sub := x[:listMaxLength-2] + ".."
		return sub
	}

	const format = "%d: %s (%.1f%%)"

	list := make([]string, len(items))

	for i, item := range items {
		list[i] = fmt.Sprintf(format, i+1, truncate(item.name), item.ratio)
	}

	return list
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
