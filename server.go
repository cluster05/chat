package main

import (
	"fmt"
	"log"
	"web-chat/api"
	"web-chat/websocket"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	if err := setupEnviroment(); err != nil {
		return err
	}
	log.Println("[env setup][init]")

	router, err := api.InitRouter()
	log.Println("[http][init]")
	if err != nil {
		return err
	}
	socket := websocket.InitSocketIO(router)
	log.Println("[socket][init]")
	defer socket.Close()

	PORT := fmt.Sprintf(":%s", "8000")
	err = router.Run(PORT)
	if err != nil {
		return err
	}
	return nil
}

func setupEnviroment() error {
	return nil
}
