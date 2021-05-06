package main

import (
	"context"
	"log"
	"os"
	"time"

	client "github.com/eben/sector-memory/client"
	server "github.com/eben/sector-memory/server"
)

func main() {
	os.Setenv("STORAGE_LISTEN", "127.0.0.1:1357")
	// os.Setenv("SECTOR_COUNTER", "0")
	go server.Run("test")
	time.Sleep(time.Second * 6)
	for i := 0; i < 6; i++ {
		sectorid := 123 + i
		sid, err := client.NewClient().ReportSectorID(context.Background(), uint64(sectorid),345,true,"dahdhadlhoa",7)
		if err != nil {
			return
		}
		log.Println("curn sector id: ", sid)
	}
}
