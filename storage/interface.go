package storage

type Storage interface {
	Create(word, explain string) error
}
