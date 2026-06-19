package fred

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type releases struct {
	Releases []release `xml:"release"`
}

type release struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	Link string `xml:"link,attr"`
	Desc string `xml:"notes,attr"`
}

func ViewReleases(baseURL string, apiKey string) string {
	url := baseURL + "releases?api_key=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "error"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("error reading response: %v", err)
	}

	var releasesBin releases

	err = xml.Unmarshal(body, &releasesBin)
	if err != nil {
		return err.Error()
	}
	for _, r := range releasesBin.Releases {
		fmt.Printf("%s - %s\n", r.Name, r.ID)
		fmt.Printf("    %s\n", r.Link)
	}

	return url
}
