package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/awesome-gocui/gocui"
	"golang.org/x/net/websocket"
)

type Message struct {
	Date    string `json:"date"`
	Message string `json:"message"`
}

const address string = "localhost:8080"

func main() {
	initWebsocketClient()
}

func initWebsocketClient() {
	gui, err := gocui.NewGui(gocui.OutputNormal, false)
	
	if err != nil {
		log.Fatalf("Failed to initialize GUI: %v", err)
	}
	defer gui.Close()

	gui.SetManagerFunc(layout)

	// koble til websocket
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/", address), "", fmt.Sprintf("http://%s/", address))
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	incomingMessages := make(chan string)
	go readClientMessages(ws, gui, incomingMessages)


	// legge til keybinds sånn at det er mulig å quitte appen
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalf("Failed to set quit key combination: %v", err)
	}
	// input field
	if err := gui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, sendMessageHandler(ws, gui)); err != nil {
		log.Fatalf("Failed to set send message key combination: %v", err)
	}

	// main loop til app gui
	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalf("Failed to start GUI main loop: %v", err)
	}
}

func readClientMessages(ws *websocket.Conn, gui *gocui.Gui, incomingMessages chan string) {
    for {
        var raw string
		var jsonString Message
        err := websocket.Message.Receive(ws, &raw)
        if err != nil {
            log.Printf("Error: %v", err)
            return
        }

		json.Unmarshal([]byte(raw), &jsonString)

        parsedTime, err := time.Parse(time.RFC3339, jsonString.Date)
        if err != nil {
            log.Printf("Error parsing date: %v", err)
            continue
        }


        formattedDate := parsedTime.Format("2006-01-02 15:04:05")

        message := fmt.Sprintf("%s : %s", formattedDate, jsonString.Message)
        // oppdater guien sånn at de nye meldingene kommer opp
        gui.Update(func(g *gocui.Gui) error {
            v, err := g.View("messages")
            if err != nil {
                return err
            }
            fmt.Fprintln(v, message)
            return nil
        })
    }
}


func sendMessage(ws *websocket.Conn, message Message) error {
	data, err := json.Marshal(message)
    if err != nil {
        return fmt.Errorf("Error marshaling message to JSON: %s", err)
    }

	err = websocket.Message.Send(ws, data)
    if err != nil {
        return fmt.Errorf("Error sending message: %s", err)
    }
    return nil
}

// bruk senere om GUI blir tatt bort

// func inputLoop(ws *websocket.Conn) {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		text, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("Error reading input: ", err)
// 			continue
// 		}
// 		text = strings.TrimSpace(text)

// 		if text == "" {
// 			continue
// 		}

// 		msg := Message{
// 			Date:    time.Now().Format(time.RFC3339),
// 			Message: text,
// 		}

// 		err = sendMessage(ws, msg)
// 		if err != nil {
// 			fmt.Printf("Failed to send message: %s\n", err.Error())
// 		}

// 	}
// }

func sendMessageHandler(ws *websocket.Conn, gui *gocui.Gui) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		message := strings.TrimSpace(v.Buffer())
		if message == "" {
			return nil
		}

		msg := Message{
			Date:    time.Now().Format(time.RFC3339),
			Message: message,
		}

		err := sendMessage(ws, msg)
		if err != nil {
			log.Printf("Failed to send message: %s\n", err.Error())
		}

		// rens input field etter melding sent
		v.Clear()
		return nil
	}
}