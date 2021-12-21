package cf

import (
	"github.com/cchaiyatad/seestern/internal/dataformat"
	"github.com/cchaiyatad/seestern/internal/file"
)

type ConfigFileReader struct {
	filepath string
	vendor   string
}

func NewConfigFileReader(filepath string, vendor string) *ConfigFileReader {
	return &ConfigFileReader{filepath: filepath, vendor: vendor}
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

	ssConfig.vendor = c.vendor
	return &ssConfig, nil
}

func (c *ConfigFileReader) getDecodeOpts() []dataformat.DecodeOption {
	var opts []dataformat.DecodeOption

	fileType, err := file.GetFileType(c.filepath)
	if err != nil {
		return opts
	}

	if fileType == "toml" {
		if fun, err := getParseAliasFunc(c.filepath); err != nil && fun != nil {
			opts = append(opts, fun)
		}
	}

	return opts
}
