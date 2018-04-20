package main

import (
	"fmt"
	"log"

	syslog "github.com/RackSec/srslog"
)

func main() {
	w, err := syslog.Dial("udp", "localhost:20129", syslog.LOG_EMERG|syslog.LOG_KERN|syslog.LOG_DEBUG, "myclient")
	if err != nil {
		log.Fatal(err)
	}

	defer w.Close()

	fmt.Print("Log successfully opened")
	w.Emerg("This is emerg log")
	w.Debug("This is debug log")

}
