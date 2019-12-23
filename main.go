package main

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
)

var Driver *agouti.WebDriver
var Page *agouti.Page

func init() {
	// driver := agouti.PhantomJS()
	//driver := agouti.ChromeDriver()
	Driver = agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}),
	)
}

func Start() {
	err := Driver.Start()
	if err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	Page, err = Driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}
}

func Stop() {
	if err := Driver.Stop(); err != nil {
		log.Fatal("Failed to close pages and stop WebDriver:", err)
	}
}

func Open(url string) {
	if err := Page.Navigate(url); err != nil {
		log.Fatal("Failed to navigate:", err)
	}
}

func Click(selector string) {
	err := Page.Find(selector).Click()
	if err != nil {
		Page.Screenshot("test.png")
		log.Fatal("Element Not Found:", err)
	}
}

func See(selector string) {
	ok, _ := Page.Find(selector).Visible()
	if !ok {
		Page.Screenshot("test.png")
		log.Fatal("Element Not Found")
	}
}

func NotSee(selector string) {
	ok, _ := Page.Find(selector).Visible()
	if ok {
		Page.Screenshot("test.png")
		log.Fatal("Element Found!")
	}
}

func Fill(selector string, value string) {
	err := Page.Find(selector).Fill(value)
	if err != nil {
		Page.Screenshot("test.png")
		log.Fatal("Fill error:", err)
	}
}

func Select(selector string, value string) {
	err := Page.Find(selector).Select(value)
	if err != nil {
		Page.Screenshot("test.png")
		log.Fatal("Select error:", err)
	}
}

func AddFile(selector string, filename string) {
	err := Page.Find(selector).UploadFile(filename)
	if err != nil {
		Page.Screenshot("test.png")
		log.Fatal("File upload error:", err)
	}
}

func Get(selector string) {
	sectionTitle, err := Page.Find(selector).Text()
	if err != nil {
		Page.Screenshot("test.png")
		log.Fatal("Element Not Found:", err)
	}
	log.Println(sectionTitle)
}

func Wait(selector string, timeout int) {
	for i := 1; i < timeout; i++ {
		ok, _ := Page.Find(selector).Visible()
		if ok {
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	Start()
	Open("http://localhost/prozorrosale2/auctions/public")
	See("#loginbtn")
	//Get("#getting-agouti")
	Stop()
}
