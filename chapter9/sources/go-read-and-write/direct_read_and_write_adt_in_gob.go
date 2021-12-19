package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

type Player struct {
	Name   string
	Age    int
	Gender string
}

func directWriteADTToFile(path string, players []Player) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("open file error:", err)
		return err
	}
	defer func() {
		f.Sync()
		f.Close()
	}()

	enc := gob.NewEncoder(f)

	for _, player := range players {
		err = enc.Encode(player)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var players = []Player{
		{"Tommy", 18, "male"},
		{"Lucy", 17, "female"},
		{"George", 19, "male"},
	}
	err := directWriteADTToFile("players.dat", players)
	if err != nil {
		fmt.Println("write file error: ", err)
		return
	}

	f, err := os.Open("players.dat")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	var player Player
	dec := gob.NewDecoder(f)
	for {
		err := dec.Decode(&player)
		if err == io.EOF {
			fmt.Println("read meet EOF")
			return
		}
		if err != nil {
			fmt.Println("read file error: ", err)
			return
		}
		fmt.Printf("%v\n", player)
	}
}
