> I meant to ask you if you know or can find on the internet - bots which basically can keep checking USCIS for these open slots.

# introducing uscis-que like que?

Sure, here is the entire README.md file in Markdown format, in a codeblock:


## How it works

The app works by first making a request to the USCIS website. The request includes the receipt number for the case that you want to check. The USCIS website returns the case status in HTML format.

The app then parses the HTML response to extract the case status. The case status is returned as a string.

## To-do

* Add support for more USCIS case types.
* Add support for storing the case status in a database.
* Add a web interface for the app.

## License

This app is licensed under the MIT License.

## More information

* USCIS website: https://www.uscis.gov/
* USCIS Case Status Online: https://egov.uscis.gov/casestatus/landing.do

## Code breakdown

```
```
The `main.go` file in the `cmd` folder contains the code for the USCIS Case Status App. The file starts by importing the following packages:

```
import (
    "fmt"
    "net/http"
    "strings"
    "time"
)
```

The `main()` function in the `main.go` file starts by creating a new `http.Client` object. This object will be used to make HTTP requests to the USCIS website.

```
func main() {
    client := &http.Client{}

    // Get the case status for the receipt number that was passed as an argument.
    caseStatus, err := getCaseStatus(client, os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the case status to the console.
    fmt.Println(caseStatus)
}
```

The `getCaseStatus()` function works by first creating a new `url.URL` object with the URL for the USCIS Case Status Online page. The URL includes the receipt number for the case that you want to check.

```
func getCaseStatus(client *http.Client, receiptNumber string) (string, error) {
    url := &url.URL{
        Scheme: "https",
        Host:   "egov.uscis.gov",
        Path:   "/casestatus/landing.do",
        Query:  fmt.Sprintf("receiptNumber=%s", receiptNumber),
    }

    // Make the HTTP request.
    response, err := client.Get(url.String())
    if err != nil {
        return "", err
    }

    // Check the status code of the response.
    if response.StatusCode != 200 {
        return "", fmt.Errorf("HTTP status code: %d", response.StatusCode)
    }

    // Parse the response body to extract the case status.
    caseStatus, err := parseCaseStatus(response.Body)
    if err != nil {
        return "", err
    }

    // Return the case status.
    return caseStatus, nil
}
```

The `parseCaseStatus()` function parses the response body to extract the case status. The response body is in HTML format, so the `parseCaseStatus()` function uses the `strings` package to extract the case status from the HTML.

```
func parseCaseStatus(body []byte) (string, error) {
    // Find the case status text.
    statusStart := strings.Index(body, "<td class=\"status\">")
    if statusStart == -1 {
        return "", fmt.Errorf("Could not find case status in HTML")
    }

    