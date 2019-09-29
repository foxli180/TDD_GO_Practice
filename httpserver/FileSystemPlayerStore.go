package httpserver

import (
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
	league, _ := NewLeague(f.database)
	return league
}
