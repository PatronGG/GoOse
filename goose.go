package goose

// Goose is the main entry point of the program
type Goose struct {
	Config Configuration
}

// New returns a new instance of the article extractor
func New(args ...string) Goose {
	return Goose{
		Config: GetDefaultConfiguration(args...),
	}
}

// ExtractFromURL follows the URL, fetches the HTML page and returns an article object
func (g Goose) ExtractFromURL(url string) (*Article, error) {
	cc := NewCrawler(g.Config, url, "")
	return cc.Crawl()
}

// ExtractFromRawHTML returns an article object from the raw HTML content
func (g Goose) ExtractFromRawHTML(url string, RawHTML string) (*Article, error) {
	cc := NewCrawler(g.Config, url, RawHTML)
	return cc.Crawl()
}
