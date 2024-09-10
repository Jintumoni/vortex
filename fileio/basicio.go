package fileio

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/Jintumoni/vortex/config"
	"github.com/Jintumoni/vortex/nodes"
	// "github.com/vmihailenco/msgpack/v5"
)

var (
	KeyNotFound = errors.New("The key does not exist")
)

func WriteToFile(path string, content io.Reader) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(content)
	if err != nil {
		return err
	}
	bytes = append(bytes, '\n')

	_, err = file.Write(bytes)
	return err
}

func ReadFromFile(path string) (io.Reader, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func FindSchemaById(key string) (*nodes.SchemaDefNode, error) {
	data, err := os.ReadFile(config.SchemaDefPath)
	if err != nil {
		return nil, err
	}

	log.Println(data)
	lines := bytes.Split(data, []byte{'\n'})
	log.Println(lines)
	for _, line := range lines {
		log.Println(string(line))
		var schemaDef *nodes.SchemaDefNode
		if err := json.Unmarshal(line, &schemaDef); err != nil {
			log.Println(string(line), "---------------")
			return nil, err
		}
		log.Println(schemaDef, "------")

		if schemaDef.SchemaName.Value == key {
			log.Println(*schemaDef)
			return schemaDef, nil
		}

	}

	return nil, KeyNotFound
}
