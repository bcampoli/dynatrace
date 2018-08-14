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
	ListenPort  int              `json:"listenPort,omitempty"`
	Credentials rest.Credentials `json:"credentials,omitempty"`
	verbose     bool
}

// NewConfig TODO: documentation
func NewConfig(listenPort int, credentials rest.Credentials) *Config {
	return &Config{ListenPort: listenPort, Credentials: credentials}
}

// ParseConfig TODO: documentation
func ParseConfig(flagset *flag.FlagSet) (*Config, error) {
	args := os.Args
	var config Config
	var err error

	config = Config{ListenPort: 0, Credentials: rest.Credentials{}}

	readConfigFromEnv(&config)
	if err = readConfigFromFlags(&config, flagset, args); err != nil {
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

	adoptConfig(target, &config, "ENV")
}

func readConfigFromFlags(target *Config, parentFlags *flag.FlagSet, args []string) error {
	var configFromFile Config
	var configFromFlags Config
	var configFileName string
	var err error

	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)

	if parentFlags != nil {
		parentFlags.VisitAll(func(flag *flag.Flag) {
			flagSet.Var(flag.Value, flag.Name, flag.Usage)
		})
	}

	flagSet.BoolVar(&configFromFlags.verbose, "v", false, "")
	flagSet.StringVar(&configFileName, "config", "", "")
	flagSet.IntVar(&configFromFlags.ListenPort, "listen", 0, "")
	flagSet.StringVar(&configFromFlags.Credentials.EnvironmentID, "environment", "", "")
	flagSet.StringVar(&configFromFlags.Credentials.Cluster, "cluster", "", "")
	flagSet.StringVar(&configFromFlags.Credentials.APIToken, "api-token", "", "")

	flagSet.Usage = func() {}
	if err = flagSet.Parse(args[1:]); err != nil {
		return err
	}

	if configFileName != "" {
		readConfigFromFile(&configFromFile, configFileName)
		adoptConfig(target, &configFromFile, "FILE")
	}

	adoptConfig(target, &configFromFlags, "FLAGS")

	return nil
}

func adoptConfig(target *Config, source *Config, sourceName string) {
	if source.verbose {
		// fmt.Println(".. adopting " + "verbose" + " from config " + sourceName)
		target.verbose = source.verbose
	}
	if source.ListenPort != 0 {
		// fmt.Println(".. adopting " + "ListenPort" + " from config " + sourceName)
		target.ListenPort = source.ListenPort
	}
	if source.Credentials.EnvironmentID != "" {
		// fmt.Println(".. adopting " + "EnvironmentID" + " from config " + sourceName)
		target.Credentials.EnvironmentID = source.Credentials.EnvironmentID
	}
	if source.Credentials.APIToken != "" {
		// fmt.Println(".. adopting " + "APIToken" + " from config " + sourceName)
		target.Credentials.APIToken = source.Credentials.APIToken
	}
	if source.Credentials.Cluster != "" {
		// fmt.Println(".. adopting " + "Cluster" + " from config " + sourceName)
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

	fmt.Println(string(bytes))

	if err = json.Unmarshal(bytes, config); err != nil {
		fmt.Println("[WARNING] " + err.Error())
		return
	}
}

func readConfigFromFile(config *Config, configFileName string) {
	var err error
	var configFile *os.File

	if configFileName != "" {
		fmt.Println("reading config file " + configFileName)
		if _, err = os.Stat(configFileName); err == nil {
			if configFile, err = os.Open(configFileName); err != nil {
				fmt.Println("[WARNING] [os.Open]" + err.Error())
				return
			}
			fromJSON(config, configFile)
			configFile.Close()
		} else {
			fmt.Println("some error happened: ", err.Error())
		}
	}
}
