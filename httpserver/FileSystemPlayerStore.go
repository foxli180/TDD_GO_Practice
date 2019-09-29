package httpserver

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f FileSystemPlayerStore) GetPlayerScore(name string) int {
	panic("implement me")
}

func (f FileSystemPlayerStore) RecordWin(name string) {
	panic("implement me")
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
