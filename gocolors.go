package gocolors

//Simply returns a slice of colors in RGB.
//The basic approach: blur, resize, kmeans cluster.
func ExtractRGB(filename string, k int, alpha float64) {
	img := blurAndResize(filename)
	imgSlice := imageToSlice(img)
	clustered := extractByKmeans(mSlice, k, alpha)
	return clustered
}

//More will be available.
