package main

import (
	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MovieResults struct {
	Results []Movie
}

type Movie struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
}

func LoadMovies() (MovieResults, error) {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		return MovieResults{}, err
	}
	var movieResults MovieResults
	err = json.Unmarshal(data, &movieResults)
	if err != nil {
		return MovieResults{}, err
	}
	return movieResults, nil
}

func main() {
	movieResults, err := LoadMovies()
	if err != nil {
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("Movie View")
	w.Resize(fyne.NewSize(1200, 800))

	listView := widget.NewList(
		func() int { return len(movieResults.Results) },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).Text = movieResults.Results[id].Title
		},
	)
	contentText := widget.NewLabel("Please select a movie title.")
	contentText.Wrapping = fyne.TextWrapWord

	listView.OnSelected = func(id widget.ListItemID) {
		contentText.Text = movieResults.Results[id].Overview
	}

	split := container.NewHSplit(
		listView,
		container.NewMax(contentText),
	)
	split.Offset = 0.3
	w.SetContent(split)

	w.ShowAndRun()
}
