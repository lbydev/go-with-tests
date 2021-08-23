package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	palyer := f.GetLeague().Find(name)
	if palyer != nil {
		wins = palyer.Wins
	}
	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}