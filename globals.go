package artty

import (
	"regexp"

	"gitlab.com/mjwhitta/artty/cache"
)

var bsRegex = regexp.MustCompile(`<p class\=\"fact\">(.+)</p>`)

// Cache related vars
var Cache *cache.ArtCache = cache.New(Version)

var devRegex = regexp.MustCompile(`<a href.+>(.+)</a>`)

// Version is the package version
const Version = "1.3.1"
