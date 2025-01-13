package checkeol

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	logToFile "python-updater/pkg/logtofile"
	"text/tabwriter"
)

type VersionInfo struct {
	Cycle             string `json:"cycle"`
	ReleaseDate       string `json:"releaseDate"`
	EOL               string `json:"eol"`
	Latest            string `json:"latest"`
	LatestReleaseDate string `json:"latestReleaseDate"`
	Support           string `json:"support"`
}

func CheckEOL() {
	res, err := http.Get("https://endoflife.date/api/python.json")

	if err != nil {
		logToFile.Log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logToFile.Log.Fatal(err)
	}

	var versions []VersionInfo
	json.Unmarshal(body, &versions)

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(writer, "Cycle\tReleaseDate\tEOL\tLatest\tLatestReleaseDate\tSupport")
	fmt.Fprintln(writer, "-----\t-----------\t----------\t--------\t-----------------\t----------")

	for _, version := range versions[:5] {
		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\t%s\n", version.Cycle, version.ReleaseDate, version.EOL, version.Latest, version.LatestReleaseDate, version.Support)
	}

	writer.Flush()
}
