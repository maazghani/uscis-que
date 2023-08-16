package main

import (
  "fmt"
  "html/template"
  "net/http"
  "strings"
  "time"
)

func get_case_status(receipt_number string) string {
  """Gets the case status for a given case number."""

  url := "https://egov.uscis.gov/casestatus/landing.do?receipt_number={}".format(
      receipt_number)

  response, err := http.Get(url)
  if err != nil {
    panic(err)
  }

  if response.StatusCode != 200 {
    panic(fmt.Sprintf("Failed to get case status. Code: %d", response.StatusCode))
  }

  body := response.Body
  defer body.Close()

  text := string(body)

  // Find the case status text.
  start := strings.Index(text, "<td class=\"status\">")
  end := strings.Index(text, "</td>")

  if start == -1 || end == -1 {
    return ""
  }

  case_status := text[start + len("<td class=\"status\">"):end]

  return case_status
}

func main() {
  // Create a web server.
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // Get the receipt number from the query string.
    receipt_number := r.URL.Query().Get("receipt_number")

    // Get the case status.
    case_status := get_case_status(receipt_number)

    // Create a template.
    tmpl := template.Must(template.ParseFiles("index.html"))

    // Render the template with the case status.
    err := tmpl.Execute(w, case_status)
    if err != nil {
      panic(err)
    }
  })

  // Start the web server.
  fmt.Println("Starting web server...")
  http.ListenAndServe(":8080", nil)
}
