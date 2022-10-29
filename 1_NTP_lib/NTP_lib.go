package NTP_lib

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func GetTime() *time.Time {
	time, err := ntp.Time("0.pool.ntp.org") //("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(time)
		fmt.Printf("%T", time)
	}
	return &time
}
