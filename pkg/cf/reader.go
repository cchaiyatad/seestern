package cf

import (
	"github.com/cchaiyatad/seestern/internal/dataformat"
	"github.com/cchaiyatad/seestern/internal/file"
)

type ConfigFileReader struct {
	filepath string
}

func NewConfigFileReader(filepath string) *ConfigFileReader {
	return &ConfigFileReader{filepath: filepath}
}

func (c *ConfigFileReader) GetSSConfig() (*SSConfig, error) {
	decoder, err := dataformat.NewDecoder(c.filepath)
	if err != nil {
		return nil, err
	}

	var ssConfig SSConfig
	if err = decoder.Decode(&ssConfig, c.getDecodeOpts()...); err != nil {
		return nil, err
	}

	return &ssConfig, nil
}

func (c *ConfigFileReader) getDecodeOpts() []dataformat.DecodeOption {
	var opts []dataformat.DecodeOption

	fileType, err := file.GetFileType(c.filepath)
	if err != nil {
		return opts
	}

	if fileType == "toml" {
		// TODO: get Alias function
	}

	return opts
}
