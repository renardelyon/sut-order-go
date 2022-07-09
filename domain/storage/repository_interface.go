package storage

type StorageRepoInterface interface {
	AddFile(string, string) error
	GetFileByUserId(userId string) ([]byte, error)
}
