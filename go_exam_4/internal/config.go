package internal

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Configs struct {
	vn         *viper.Viper
	ConfigPath string
	State      string

	Validator *validator.Validate
	TimeZone  string  `mapstructure:"time_zone"`
	MongoDB   MongoDB `mapstructure:"mongodb"`

	BangkokLocation *time.Location
}

type MongoDB struct {
	MongoUri string        `mapstructure:"mongodb_uri"`
	Timeout  time.Duration `mapstructure:"timeout"`
	Username string        `mapstructure:"username"`
	Password string        `mapstructure:"password"`
	Database string        `mapstructure:"database"`

	Client *mongo.Client
}

func (c *Configs) InitAllConfigurations(s string) error {
	if s == "" {
		s = "local"
	}

	name := fmt.Sprintf("config.%s", s)
	log.Infof("config file using : %s", name)

	if c.ConfigPath == "" {
		c.ConfigPath = "./go_exam_4/configs"
	}

	vn := viper.New()
	vn.AddConfigPath(c.ConfigPath)
	vn.SetConfigName(name)
	c.vn = vn
	c.State = s

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	if err := c.vn.Unmarshal(&c); err != nil {
		return err
	}

	if err := c.MongoDB.bindingClient(); err != nil {
		return errors.Wrap(err, "binding mongo error")
	}

	loc, err := time.LoadLocation(c.TimeZone)
	if err != nil {
		return errors.Wrapf(err, "load location %s error", c.TimeZone)
	}
	c.BangkokLocation = loc

	log.Infof("all config loaded : %#v", c)
	return nil
}

func (m *MongoDB) bindingClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn := fmt.Sprintf(m.MongoUri, m.Username, m.Password, m.Database)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return err
	}

	// check connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		return errors.Wrapf(err, "ping failed")
	}
	m.Client = client

	return nil
}