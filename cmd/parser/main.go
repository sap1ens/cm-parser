package main

import (
    "fmt"
    "log"
    "strings"
    "time"
    "os"
    "net/http"
    "bytes"

    "github.com/PuerkitoBio/goquery"
    "io/ioutil"
)

const CityToFind = "ванкувер"
const LastDateKnown = "28.03.17"
const DateLayout = "02.01.06"

func main() {
    document := fetchDocument("http://canada.mid.ru/novosti")

    var newEvent string

    document.Find(".asset-title a").EachWithBreak(func(i int, s *goquery.Selection) bool {
        title := strings.TrimSpace(s.Text())

        blockParent := s.ParentsFiltered(".asset-abstract")

        dateString := strings.TrimSpace(blockParent.Find(".metadata-publish-date").Text())
        date, err := time.Parse(DateLayout, dateString)

        if err != nil {
            return true // next
        }

        if isNewEvent(title, date) {
            newEvent = title
            return false // break
        }

        return true // next
    })

    if newEvent != "" {
        log.Printf("New event found: %s", newEvent)

        sendNotification(newEvent)
    } else {
        log.Print("Nothing found")
    }
}

func fetchDocument(url string) (doc *goquery.Document) {
    log.Print("Fetching page")

    doc, err := goquery.NewDocument(url)
    if err != nil {
        log.Fatal(err)
    }

    return
}

func isNewEvent(rawTitle string, date time.Time) bool {
    title := strings.ToLower(rawTitle)
    lastDate, _ := time.Parse(DateLayout, LastDateKnown)

    return strings.Contains(title, CityToFind) && date.After(lastDate)
}

func sendNotification(message string) {
    slackURL := os.Getenv("SLACK_URL")
    data := []byte(fmt.Sprintf(`{"channel": "#notifications", "username": "canada mid crawler", "text": "%s", "icon_emoji": ":robot_face:"}`, message))
    req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(data))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}