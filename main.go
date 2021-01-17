package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// load HTML
	ui.Load("data:text/html," + url.PathEscape(`
	<html>
		<body>
			<p>hello</p>
		</body>
	</html>
	`))

	// wait until the interrupt signal arrives or browser widow is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	log.Println("existing...")
}
