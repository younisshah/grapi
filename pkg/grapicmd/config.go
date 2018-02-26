package grapicmd

import (
	"io"

	"github.com/spf13/afero"
	"github.com/spf13/viper"

	"github.com/izumin5210/grapi/pkg/grapicmd/protoc"
)

// Config stores general setting params and provides accessors for them.
type Config interface {
	Init(cfgFile string)
	Fs() afero.Fs
	RootDir() string
	AppName() string
	Version() string
	Revision() string
	InReader() io.Reader
	OutWriter() io.Writer
	ErrWriter() io.Writer
	ProtocConfig() *protoc.Config
}

// NewConfig creates new Config object.
func NewConfig(
	rootDir string,
	appName, version, revision string,
	in io.Reader,
	out, err io.Writer,
) Config {
	return &config{
		v:        viper.New(),
		fs:       afero.NewOsFs(),
		rootDir:  rootDir,
		appName:  appName,
		version:  version,
		revision: revision,
		in:       in,
		out:      out,
		err:      err,
	}
}

type config struct {
	cfgFile                    string
	v                          *viper.Viper
	fs                         afero.Fs
	rootDir                    string
	appName, version, revision string
	in                         io.Reader
	out, err                   io.Writer
	readConfigErr              error
}

func (c *config) Init(cfgFile string) {
	c.cfgFile = cfgFile
	c.v.SetConfigFile(c.cfgFile)
	c.readConfigErr = c.v.ReadInConfig()
}

func (c *config) Fs() afero.Fs {
	return c.fs
}

func (c *config) RootDir() string {
	return c.rootDir
}

func (c *config) AppName() string {
	return c.appName
}

func (c *config) Version() string {
	return c.version
}

func (c *config) Revision() string {
	return c.revision
}

func (c *config) InReader() io.Reader {
	return c.in
}

func (c *config) OutWriter() io.Writer {
	return c.out
}

func (c *config) ErrWriter() io.Writer {
	return c.err
}

func (c *config) ProtocConfig() *protoc.Config {
	cfg := &protoc.Config{}
	c.v.UnmarshalKey("protoc", cfg)
	return cfg
}