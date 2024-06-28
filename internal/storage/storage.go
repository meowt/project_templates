package storage

import "project/internal/config"

type Storage struct {
}

func New(cfg *config.Config) (*Storage, error) {
	return &Storage{}, nil
}
