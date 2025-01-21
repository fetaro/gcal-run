package gcal_run

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"io/ioutil"
	"os"
	"strings"
)

const (
	Delimiter = "\n"
)

type EventIDStore struct {
	filePath string
}

func NewEventIDStore(filePath string) *EventIDStore {
	return &EventIDStore{filePath}
}

func (e *EventIDStore) SaveID(id string) error {
	if id == "" {
		return fmt.Errorf("id is empty")
	}
	file, err := os.OpenFile(e.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprint(file, id+Delimiter)
	return nil
}

func (e *EventIDStore) LoadIDList() ([]string, error) {
	bytes, err := ioutil.ReadFile(e.filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSuffix(string(bytes), Delimiter), Delimiter), nil
}

func (e *EventIDStore) IsInclude(id string) (bool, error) {
	if !common.FileExists(e.filePath) {
		return false, nil
	}
	loadedIDList, err := e.LoadIDList()
	if err != nil {
		return false, err
	}
	for _, loadedID := range loadedIDList {
		if id == loadedID {
			return true, nil
		}
	}
	return false, nil
}

func (e *EventIDStore) Clear() error {
	return os.WriteFile(e.filePath, []byte(""), 0666)
}
