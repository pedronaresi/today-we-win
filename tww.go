package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
    "time"
    "os/exec"
    "os"
)

func main() {
    for {
    // Create the client profile.
    client := &http.Client{
        Timeout: 30 * time.Second,
    }

    // Make HTTP request
    // Place the URL of the sneaker on the line below.
    req, err := http.NewRequest("GET", "https://www.nike.com.br/calendario-lancamentos/sportswear/rise-react-flyknit-diffused-taupe?icid=109434", nil)
    req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
    req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
    response, err := client.Do(req)

    defer response.Body.Close()

    // Load the HTML file on the memory.
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal("Error reading HTTP body. ", err)
    }

    // Create the regular expression to find if the "Avise-me" tag is still online.
    re := regexp.MustCompile("Avise-me")
    comments := re.FindAllString(string(body), -1)
    if comments == nil {
      //If the "Avise-me" dissapear, you are ready to go!
      exec.Command("open", "https://www.nike.com.br/calendario-lancamentos/sportswear/the-ten-blazer-mid-serena?icid=109423").Start()
      fmt.Println("Lets cop!")
      os.Exit(1)
    } else {
      fmt.Println("No matches yet.")
    }
  }
  time.Sleep(500)
}
