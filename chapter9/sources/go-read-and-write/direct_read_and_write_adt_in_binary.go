package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Player struct {
	Name   [20]byte
	Age    int16
	Gender [6]byte
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

	for _, player := range players {
		err = binary.Write(f, binary.BigEndian, &player)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var players [3]Player

	copy(players[0].Name[:], []byte("Tommy"))
	players[0].Age = 18
	copy(players[0].Gender[:], []byte("male"))

	copy(players[1].Name[:], []byte("Lucy"))
	players[1].Age = 17
	copy(players[1].Gender[:], []byte("female"))

	copy(players[2].Name[:], []byte("George"))
	players[2].Age = 19
	copy(players[2].Gender[:], []byte("male"))

	err := directWriteADTToFile("players.dat", players[:])
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
	for {
		err = binary.Read(f, binary.BigEndian, &player)
		if err == io.EOF {
			fmt.Println("read meet EOF")
			return
		}
		if err != nil {
			fmt.Println("read file error: ", err)
			return
		}
		fmt.Printf("%s %d %s\n", player.Name, player.Age, player.Gender)
	}
}
