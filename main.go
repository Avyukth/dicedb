import (
	"flag"
	"log"
	"github.com/dicedb/dice/config"
	"github.com/dicedb/dice/server"
)

func main() {
	setupFlags()
	log.Println("Starting dice server ")
	server.RunSyncTCPServer()
}
