package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	pErr "github.com/pkg/errors"
)

var (
	ErrNoConfigFile = errors.New("can not find config file")
	ErrFileType     = errors.New("the file must be JSON")
)

type Config interface {
	Load() error
	Scan(interface{}) error
}

type config struct {
	path string
	data []byte
}

func New(path string) *config {
	return &config{path: path}
}

func (c *config) Load() error {
	dat, err := ioutil.ReadFile(c.path)
	if err != nil {
		return pErr.WithMessage(ErrNoConfigFile, c.path)
	}

	c.data = dat
	return nil
}

func (c *config) Scan(v interface{}) error {
	if m, ok := v.(proto.Message); ok {
		return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(c.data, m)
	}
	return json.Unmarshal(c.data, v)
}
