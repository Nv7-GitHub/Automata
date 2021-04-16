package main

func initAutomata() {
	var w, h int
	if isHidpi {
		w = width * 2
		h = height * 2
	} else {
		w = width
		h = height
	}

	frame = make([][]Cell, h)
	for i := range frame {
		frame[i] = make([]Cell, w)
	}
}
