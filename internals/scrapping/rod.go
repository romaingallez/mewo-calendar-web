package scraping

import (
	"bytes"
	"image"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/ysmood/gson"
)

// Display page
func RenderPage(name string, c *fiber.Ctx) error {
	return c.Render(name, nil)
}

func GetImage(html string) (img image.Image, format string, err error) {

	wsURL := os.Getenv("WSURL")
	var page *rod.Page

	if len(wsURL) == 0 {

		path, _ := launcher.LookPath()

		log.Println(path)

		// Devtools opens the tab in each new tab opened automatically
		l := launcher.New().
			Bin(path).
			Headless(true).
			Devtools(true)

		defer l.Cleanup() // remove launcher.FlagUserDataDir

		rodUrl := l.MustLaunch()

		// Trace shows verbose debug information for each action executed
		// SlowMotion is a debug related function that waits 2 seconds between
		// each action, making it easier to inspect what your code is doing.
		browser := rod.New().
			ControlURL(rodUrl).
			Trace(true).
			SlowMotion(1 * time.Second).
			MustConnect()

		defer browser.MustClose()

		// browser.NoDefaultDevice()
		browser.NoDefaultDevice()

		// page := browser.MustPage(url)

		page = browser.MustPage("about:blank")

	} else {
		log.Println(wsURL)
		page = rod.New().ControlURL(wsURL).MustConnect().MustPage("about:blank")
		log.Println("connected to ws")

	}

	page.SetDocumentContent(html)

	page.SetWindow(&proto.BrowserBounds{
		Width:  gson.Int(1920),
		Height: gson.Int(1080),
		// X:      proto.Int(0),
		// Y:      proto.Int(0),
	})

	err = page.Emulate(devices.Device{
		UserAgent: "Mozilla/5.0 (romain bot rod)",
	})

	if err != nil {
		log.Println(err)
	}

	page.MustWaitLoad()

	screenshot := page.MustElement("body").MustScreenshot()

	// convert screenshot []byte to image.Image
	img, format, err = image.Decode(bytes.NewReader(screenshot))
	log.Println(format)
	if err != nil {
		log.Println(err)
	}

	// time.Sleep(1 * time.Hour)

	// return "https://picsum.photos/200/300"

	// use go-rod to get screenshot of page

	return img, format, err
}
