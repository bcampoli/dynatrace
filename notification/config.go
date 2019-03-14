package notification

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/dtcookie/dynatrace/log"
	"github.com/dtcookie/dynatrace/rest"
)

// Config TODO: documentation
type Config struct {
	ListenPort  int              `json:"listenPort,omitempty"`
	Credentials rest.Credentials `json:"credentials,omitempty"`
	NoProxy     bool             `json:"noproxy,omitempty"`
	Insecure    bool             `json:"insecure,omitempty"`
	Verbose     bool             `json:"verbose,omitempty"`
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

	if config.Credentials.APIToken == "" || config.Credentials.APIBaseURL == "" {
		log.Info("API Token or API Base URL not specified - fetching problem details disabled")
	}

	// if config.Credentials.APIToken == "" {
	// 	return nil, errors.New("no api token specified")
	// }

	// if config.Credentials.APIBaseURL == "" {
	// 	return nil, errors.New("no api base url specified")
	// }

	return &config, nil
}

func readConfigFromEnv(target *Config) {
	var config Config
	var sListenPort string
	var listenPort int
	var apiBaseURL string
	var apiToken string
	var err error

	apiBaseURL = os.Getenv("DT_API_BASE_URL")
	if apiBaseURL != "" {
		config.Credentials.APIBaseURL = apiBaseURL
	}
	apiToken = os.Getenv("DT_API_TOKEN")
	if apiToken != "" {
		config.Credentials.APIToken = apiToken
	}
	sListenPort = os.Getenv("DT_LISTEN_PORT")
	if sListenPort != "" {
		if listenPort, err = strconv.Atoi(sListenPort); err != nil {
			fmt.Println("the value environment variable '" + "DT_LISTEN_PORT" + "' is not a valid listen port")
		} else {
			config.ListenPort = listenPort
		}
	}

	adoptConfig(target, &config)
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

	flagSet.BoolVar(&configFromFlags.Verbose, "v", false, "")
	flagSet.BoolVar(&configFromFlags.Insecure, "insecure", false, "")
	flagSet.BoolVar(&configFromFlags.NoProxy, "noproxy", false, "")
	flagSet.StringVar(&configFileName, "config", "", "")
	flagSet.IntVar(&configFromFlags.ListenPort, "listen", 0, "")
	flagSet.StringVar(&configFromFlags.Credentials.APIBaseURL, "api-base-url", "", "")
	flagSet.StringVar(&configFromFlags.Credentials.APIToken, "api-token", "", "")

	flagSet.Usage = func() {}
	if err = flagSet.Parse(args[1:]); err != nil {
		return err
	}

	if configFileName != "" {
		readConfigFromFile(&configFromFile, configFileName)
		adoptConfig(target, &configFromFile)
	}

	adoptConfig(target, &configFromFlags)

	return nil
}

func adoptConfig(target *Config, source *Config) {
	if source.Verbose {
		target.Verbose = source.Verbose
	}
	if source.ListenPort != 0 {
		target.ListenPort = source.ListenPort
	}
	if source.NoProxy {
		target.NoProxy = source.NoProxy
	}
	if source.Insecure {
		target.Insecure = source.Insecure
	}
	if source.Credentials.APIBaseURL != "" {
		target.Credentials.APIBaseURL = source.Credentials.APIBaseURL
	}
	if source.Credentials.APIToken != "" {
		target.Credentials.APIToken = source.Credentials.APIToken
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
			fromJSON(config, configFile)
			configFile.Close()
		} else {
			fmt.Println("some error happened: ", err.Error())
		}
	}
}
