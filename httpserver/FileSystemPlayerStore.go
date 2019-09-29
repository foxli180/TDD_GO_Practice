package httpserver

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f FileSystemPlayerStore) GetPlayerScore(name string) int {
	panic("implement me")
}

func (f FileSystemPlayerStore) RecordWin(name string) {
	panic("implement me")
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(f.database).Decode(&league)
	return league
}
