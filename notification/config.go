package notification

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/dtcookie/dynatrace/rest"
)

// Config TODO: documentation
type Config struct {
	ListenPort  int
	Credentials *rest.Credentials
}

// NewConfig TODO: documentation
func NewConfig(listenPort int, credentials *rest.Credentials) *Config {
	return &Config{ListenPort: listenPort, Credentials: credentials}
}

// ParseConfig TODO: documentation
func ParseConfig() (*Config, error) {
	args := os.Args
	var config Config
	var err error

	config = Config{ListenPort: 0, Credentials: &rest.Credentials{}}

	readConfigFromEnv(&config)
	if err = readConfigFromFlags(&config, args); err != nil {
		return nil, err
	}

	if config.ListenPort == 0 {
		return nil, errors.New("no listen port specified")
	}

	if config.Credentials.APIToken == "" {
		return nil, errors.New("no api token specified")
	}

	if config.Credentials.EnvironmentID == "" {
		return nil, errors.New("no environment id specified")
	}

	return &config, nil
}

func readConfigFromEnv(target *Config) {
	var config Config
	var sListenPort string
	var listenPort int
	var environmentID string
	var apiToken string
	var cluster string
	var err error

	environmentID = os.Getenv("DT_NOTIFICATION_ENVIRONMENT")
	if environmentID != "" {
		config.Credentials.EnvironmentID = environmentID
	}
	apiToken = os.Getenv("DT_NOTIFICATION_API_TOKEN")
	if apiToken != "" {
		config.Credentials.APIToken = apiToken
	}
	cluster = os.Getenv("DT_NOTIFICATION_CLUSTER")
	if cluster != "" {
		config.Credentials.Cluster = cluster
	}
	sListenPort = os.Getenv("DT_NOTIFICATION_LISTEN_PORT")
	if sListenPort != "" {
		if listenPort, err = strconv.Atoi(sListenPort); err != nil {
			fmt.Println("the value environment variable '" + "DT_NOTIFICATION_LISTEN_PORT" + "' is not a valid listen port")
		} else {
			config.ListenPort = listenPort
		}
	}

	adoptConfig(target, &config)
}

func readConfigFromFlags(target *Config, args []string) error {
	var config Config
	var configFromFile Config
	var configFromFlags Config
	var configFileName string
	var err error

	configFromFile = Config{ListenPort: 0, Credentials: &rest.Credentials{}}
	configFromFlags = Config{ListenPort: 0, Credentials: &rest.Credentials{}}

	flagSet := flag.NewFlagSet("dynatrace", flag.ContinueOnError)
	flagSet.StringVar(&configFileName, "config", "", "")
	flagSet.IntVar(&configFromFlags.ListenPort, "listen", 0, "")
	flagSet.StringVar(&configFromFlags.Credentials.EnvironmentID, "environment", "", "")
	flagSet.StringVar(&configFromFlags.Credentials.Cluster, "cluster", "", "")
	flagSet.StringVar(&configFromFlags.Credentials.APIToken, "apiToken", "", "")
	if err = flagSet.Parse(args); err != nil {
		return err
	}

	if configFileName != "" {
		readConfigFromFile(&configFromFile, configFileName)
		adoptConfig(&configFromFile, &configFromFlags)
	}

	adoptConfig(&config, &configFromFlags)

	return nil
}

func adoptConfig(target *Config, source *Config) {
	if source.ListenPort != 0 {
		target.ListenPort = source.ListenPort
	}
	if source.Credentials.EnvironmentID != "" {
		target.Credentials.EnvironmentID = source.Credentials.EnvironmentID
	}
	if source.Credentials.APIToken != "" {
		target.Credentials.APIToken = source.Credentials.APIToken
	}
	if source.Credentials.Cluster != "" {
		target.Credentials.Cluster = source.Credentials.Cluster
	}
}

func fromJSON(config *Config, configFile *os.File) {
	var bytes []byte
	var err error

	if bytes, err = ioutil.ReadAll(configFile); err != nil {
		fmt.Println("[WARNING] [ioutil.ReadAll]" + err.Error())
		return
	}

	if err = json.Unmarshal(bytes, config); err != nil {
		fmt.Println("[WARNING] " + err.Error())
		return
	}

	return
}

func readConfigFromFile(config *Config, configFileName string) {
	var err error
	var configFile *os.File

	if configFileName != "" {
		if _, err = os.Stat(configFileName); err == nil {
			if configFile, err = os.Open(configFileName); err != nil {
				fmt.Println("[WARNING] [os.Open]" + err.Error())
				return
			}
			defer configFile.Close()
			fromJSON(config, configFile)
		}
	}
}
