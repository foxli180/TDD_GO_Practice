package httpserver

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	Database io.Writer
	League   League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)
	return &FileSystemPlayerStore{
		Database: &tape{file},
		League:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.League
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.League.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.League.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.League = append(f.League, Player{name, 1})
	}
	json.NewEncoder(f.Database).Encode(f.League)
}
