package json_server

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
	League   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		Database: &Tape{database},
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

	f.Database.Seek(0, io.SeekStart)
	json.NewEncoder(f.Database).Encode(f.League)
}
