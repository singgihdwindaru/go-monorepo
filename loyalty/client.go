package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	resources "github.com/singgihdwindaru/go-monorepo/internal/connection_manager/redis"
)

func main() {
	db, err := resources.GetRedisClient()
	if err != nil {
		fmt.Println("Error getting DB in pkgA:", err)
		return
	}

	fmt.Println("Loyalty is using the shared resource:", db)
	fmt.Println("Loyalty is using the shared resource:", db)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
