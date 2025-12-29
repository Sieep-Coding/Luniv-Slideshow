package main

import (
	"Luniv-Slideshow/lunivslideshow"
)

func main() {
	show := lunivslideshow.ReturnDefaultSlideShow()
	lunivslideshow.RenderCurrentSlide(&show)
	lunivslideshow.AdvanceSlide(&show)
	lunivslideshow.RenderCurrentSlide(&show)
}
