package engine

import (
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func testResults(t *testing.T, engine Engine) {
	viper.SetDefault("selenium-url", os.Getenv("GOPHIE_SELENIUM_URL"))

	counter := map[string]int{}
	var result SearchResult
	var searchTerm string
	if !strings.HasPrefix(engine.String(), "TvSeries") {
		searchTerm = "jumanji"
	} else {
		// search for the flash for movie series
		searchTerm = "devs"
	}
	result = engine.Search(searchTerm)

	if len(result.Movies) < 1 {
		t.Errorf("No movies returned from %v", engine.String())
	} else {
		for _, movie := range result.Movies {
			if _, ok := counter[movie.DownloadLink.String()]; ok {
				t.Errorf("Duplicated Link")
			} else {
				counter[movie.DownloadLink.String()] = 1
			}
			if movie.IsSeries == false {
				downloadlink := movie.DownloadLink.String()
				if !(strings.HasSuffix(downloadlink, "1") || strings.HasSuffix(downloadlink, ".mp4") || strings.Contains(downloadlink, ".mkv") || strings.Contains(downloadlink, ".avi") || strings.Contains(downloadlink, ".webm") || strings.Contains(downloadlink, "freeload") || strings.Contains(downloadlink, "download_token=") || strings.Contains(downloadlink, "mycoolmoviez")) {
					t.Errorf("Could not obtain link for single movie, linked returned is %v", downloadlink)
				}
			}
		}
	}
}

func TestEngines(t *testing.T) {
	engines := GetEngines()
	for _, engine := range engines {
		testResults(t, engine)
	}
}
