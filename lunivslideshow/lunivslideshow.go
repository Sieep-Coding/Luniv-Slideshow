package lunivslideshow

import "fmt"

type SlideShow struct {
	Slides       []Slide
	CurrentIndex int
}

type Slide struct {
	Title           string
	BackgroundColor string
	BackgroundImage string
	Content         string
}

type BackgroundColors struct {
	Blue  string
	Green string
	Red   string
	Black string
	White string
}

func ConvertColorsToHex() BackgroundColors {
	return BackgroundColors{
		Blue:  "#6060e3ff",
		Green: "#71c171ff",
		Red:   "#e06666ff",
		Black: "#242424ff",
		White: "#e7e7e7ff",
	}
}

func ReturnDefaultSlideShow() SlideShow {
	return SlideShow{
		Slides: []Slide{
			{
				Title:           "Welcome to Luniv Slideshow",
				BackgroundColor: BackgroundColors{}.Blue,
				Content:         "This is the first slide.",
			},
			{
				Title:           "Features",
				BackgroundColor: BackgroundColors{}.Green,
				Content:         "Luniv Slideshow supports multiple slides.",
			},
		},
		CurrentIndex: 0,
	}
}

func AdvanceSlide(show *SlideShow) {
	if show.CurrentIndex < len(show.Slides)-1 {
		show.CurrentIndex++
	}
}

func RenderCurrentSlide(show *SlideShow) {
	if show.CurrentIndex < len(show.Slides) {
		slide := show.Slides[show.CurrentIndex]
		if slide.Title == "" && slide.Content == "" {
			fmt.Println("Slide or Slide Title is empty.")
			return
		}
		fmt.Printf("Title: %s\nBackground Color: %s\nContent: %s\n",
			slide.Title, slide.BackgroundColor, slide.Content)
	} else {
		fmt.Println("No slide to display.")
	}
}
