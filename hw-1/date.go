package date

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

func GetCurrentTime() {
	time, err := ntp.Time("pool.ntp.org")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Network time: %v\n", time)
	os.Exit(1)
}
