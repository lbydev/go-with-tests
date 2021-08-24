package main

import (
	"encoding/json"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0,0)
	league,_ := NewLeague(file)
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&Tape{file}),
		league: league,
	}
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	palyer := f.league.Find(name)
	if palyer != nil {
		wins = palyer.Wins
	}
	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	f.database.Encode(f.league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}
