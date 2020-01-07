package cmd

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	corecfg "git.ecd.axway.int/apigov/aws_apigw_discovery_agent/core/config"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// CommandHandler - Root command execution handler
type CommandHandler func() error

// InitConfigHandler - Handler to be invoked on config initialization
type InitConfigHandler func(centralConfig corecfg.CentralConfig) (interface{}, error)

// AgentRootCmd - Root Command for the Agents
type AgentRootCmd interface {
	RootCmd() *cobra.Command
	Execute() error

	// Methods for adding yaml properties and command flag
	AddStringProperty(name, flagName string, defaultVal string, description string)
	AddDurationProperty(name, flagName string, defaultVal time.Duration, description string)
	AddIntProperty(name, flagName string, defaultVal int, description string)
	AddBoolProperty(name, flagName string, defaultVal bool, description string)
	AddStringSliceProperty(name, flagName string, defaultVal []string, description string)

	// Methods to get the configured properties
	StringPropertyValue(name string) string
	DurationPropertyValue(name string) time.Duration
	IntPropertyValue(name string) int
	BoolPropertyValue(name string) bool
	StringSlicePropertyValue(name string) []string

	// Get the agentType
	GetAgentType() corecfg.AgentType
}

// agentRootCommand - Represents the agent root command
type agentRootCommand struct {
	configPath        string
	agentName         string
	rootCmd           *cobra.Command
	commandHandler    CommandHandler
	initConfigHandler InitConfigHandler
	agentType         corecfg.AgentType
}

// NewRootCmd - Creates a new Agent Root Command
func NewRootCmd(exeName, desc string, initConfigHandler InitConfigHandler, commandHandler CommandHandler, agentType corecfg.AgentType) AgentRootCmd {
	c := &agentRootCommand{
		configPath:        ".",
		agentName:         exeName,
		commandHandler:    commandHandler,
		initConfigHandler: initConfigHandler,
		agentType:         agentType,
	}

	c.rootCmd = &cobra.Command{
		Use:     c.agentName,
		Short:   desc,
		Version: BuildVersion,
		RunE:    c.run,
		PreRun:  c.initialize,
	}

	// APIC yaml properties and command flags
	c.AddStringProperty("central.tenantId", "centralTenantId", "", "Tenant ID for the owner of the environment")
	c.AddStringProperty("central.auth.privateKey", "authPrivateKey", "/etc/private_key.pem", "Path to the private key for API Central Authentication")
	c.AddStringProperty("central.auth.publicKey", "authPublicKey", "/etc/public_key", "Path to the public key for API Central Authentication")
	c.AddStringProperty("central.auth.keyPassword", "authKeyPassword", "", "Password for the private key, if needed")
	c.AddStringProperty("central.auth.url", "authUrl", "https://login-preprod.axway.com/auth", "API Central authentication URL")
	c.AddStringProperty("central.auth.realm", "authRealm", "Broker", "API Central authentication Realm")
	c.AddStringProperty("central.auth.clientId", "authClientId", "", "Client ID for the service account")
	c.AddDurationProperty("central.auth.timeout", "authTimeout", 10*time.Second, "Timeout waiting for AxwayID response")
	// ssl properties and command flags
	c.AddStringSliceProperty("central.ssl.nextProtos", "centralSSLNextProtos", []string{}, "List of supported application level protocols")
	c.AddBoolProperty("central.ssl.insecureSkipVerify", "centralSSLInsecureSkipVerify", false, "Controls whether a client verifies the server's certificate chain and host name")
	c.AddStringSliceProperty("central.ssl.cipherSuites", "centralSSLCipherSuites", corecfg.TLSDefaultCipherSuitesStringSlice(), "List of supported cipher suites")
	c.AddStringProperty("central.ssl.minVersion", "centralSSLMinVersion", corecfg.TLSDefaultMinVersionString(), "Minimum acceptable SSL/TLS protocol version")
	c.AddStringProperty("central.ssl.maxVersion", "centralSSLMaxVersion", "0", "Maximum acceptable SSL/TLS protocol version")

	if c.GetAgentType() == corecfg.TraceabilityAgent {

		c.AddStringProperty("central.deployment", "centralDeployment", "preprod", "API Central")
		c.AddStringProperty("central.environmentId", "centralEnvironmentId", "", "Environment ID for the current environment")
	} else {
		c.AddStringProperty("central.mode", "centralMode", "disconnected", "Agent Mode")
		c.AddStringProperty("central.apiServerEnvironment", "apiServerEnvironment", "", "The Environment that the APIs will be associated with in API Central")
		c.AddStringProperty("central.url", "centralUrl", "https://apicentral.preprod.k8s.axwayamplify.com", "URL of API Central")
		c.AddStringProperty("central.teamId", "centralTeamId", "", "Team ID for the current default team for creating catalog")
	}

	// Log yaml properties and command flags
	c.AddStringProperty("log.level", "logLevel", "info", "Log level (debug, info, warn, error)")
	c.AddStringProperty("log.format", "logFormat", "json", "Log format (json, line, package)")
	c.AddStringProperty("log.output", "logOutput", "stdout", "Log output type (stdout, file, both)")
	c.AddStringProperty("log.path", "logPath", "logs", "Log file path if output type is file or both")
	return c
}

func (c *agentRootCommand) initialize(cmd *cobra.Command, args []string) {
	viper.SetConfigName(c.agentName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.configPath)
	viper.SetTypeByDefaultValue(true)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}

func (c *agentRootCommand) bindOrPanic(key string, flg *flag.Flag) {
	if err := viper.BindPFlag(key, flg); err != nil {
		panic(err)
	}
}

// initConfig - Initializes the central config and invokes initConfig handler
// to initialize the agent config. Performs validation on returned agent config
func (c *agentRootCommand) initConfig() error {
	logLevel := c.StringPropertyValue("log.level")
	logFormat := c.StringPropertyValue("log.format")
	logOutput := c.StringPropertyValue("log.output")
	logPath := c.StringPropertyValue("log.path")
	SetupLogging(c.agentName, logLevel, logFormat, logOutput, logPath)

	// Init Central Config
	centralCfg, err := c.parseCentralConfig()
	if err != nil {
		return err
	}

	// Initialize Agent Config
	agentCfg, err := c.initConfigHandler(centralCfg)
	if err != nil {
		return err
	}

	// Validate Agent Config
	return c.validateAgentConfig(agentCfg)
}

func (c *agentRootCommand) parseCentralConfig() (corecfg.CentralConfig, error) {
	cfg := &corecfg.CentralConfiguration{
		AgentType: c.agentType,
		TenantID:  c.StringPropertyValue("central.tenantId"),
		Auth: &corecfg.AuthConfiguration{
			URL:        c.StringPropertyValue("central.auth.url"),
			Realm:      c.StringPropertyValue("central.auth.realm"),
			ClientID:   c.StringPropertyValue("central.auth.clientID"),
			PrivateKey: c.StringPropertyValue("central.auth.privateKey"),
			PublicKey:  c.StringPropertyValue("central.auth.publicKey"),
			KeyPwd:     c.StringPropertyValue("central.auth.keyPassword"),
			Timeout:    c.DurationPropertyValue("central.auth.timeout"),
		},
		TLS: &corecfg.TLSConfiguration{
			NextProtos:         c.StringSlicePropertyValue("central.ssl.nextProtos"),
			InsecureSkipVerify: c.BoolPropertyValue("central.ssl.insecureSkipVerify"),
			CipherSuites:       corecfg.NewCipherArray(c.StringSlicePropertyValue("central.ssl.cipherSuites")),
			MinVersion:         corecfg.TLSVersionAsValue(c.StringPropertyValue("central.ssl.minVersion")),
			MaxVersion:         corecfg.TLSVersionAsValue(c.StringPropertyValue("central.ssl.maxVersion")),
		},
	}

	if c.GetAgentType() == corecfg.TraceabilityAgent {
		cfg.APICDeployment = c.StringPropertyValue("central.deployment")
		cfg.EnvironmentID = c.StringPropertyValue("central.environmentId")
	} else {
		cfg.URL = c.StringPropertyValue("central.url")
		cfg.Mode = corecfg.StringAgentModeMap[strings.ToLower(c.StringPropertyValue("central.mode"))]
		cfg.APIServerEnvironment = c.StringPropertyValue("central.apiServerEnvironment")
		cfg.APIServerVersion = c.StringPropertyValue("central.apiServerVersion")
		cfg.TeamID = c.StringPropertyValue("central.teamId")
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

// validateAgentConfig - Validates the agent config
// Uses reflection to get the Validate method on the config struct or
// struct variable.
// Makes call to Validate method except if the struct variable is of CentralConfig type
// as the validation for CentralConfig is already done during parseCentralConfig
func (c *agentRootCommand) validateAgentConfig(agentCfg interface{}) error {
	// Check if top level struct has Validate. If it does then call Validate
	// only at top level
	if objInterface, ok := agentCfg.(interface{ Validate() error }); ok {
		return objInterface.Validate()
	}

	// If the parameter is of struct pointer, use indirection to get the
	// real value object
	v := reflect.ValueOf(agentCfg)
	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}

	// Look for Validate method on struct properties and invoke it
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			fieldInterface := v.Field(i).Interface()
			// Skip the property it is CentralConfig type as its already Validated
			// during parseCentralConfig
			if _, ok := fieldInterface.(corecfg.CentralConfig); !ok {
				if objInterface, ok := fieldInterface.(interface{ Validate() error }); ok {
					err := objInterface.Validate()
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

// run - Executes the agent command
func (c *agentRootCommand) run(cmd *cobra.Command, args []string) (err error) {
	err = c.initConfig()

	if err == nil {
		log.Infof("Starting %s (%s)", c.rootCmd.Short, c.rootCmd.Version)
		err = c.commandHandler()
	}

	return
}

func (c *agentRootCommand) RootCmd() *cobra.Command {
	return c.rootCmd
}

func (c *agentRootCommand) Execute() error {
	return c.rootCmd.Execute()
}

func (c *agentRootCommand) AddStringProperty(name, flagName string, defaultVal string, description string) {
	if c.rootCmd != nil {
		c.rootCmd.Flags().String(flagName, defaultVal, description)
		c.bindOrPanic(name, c.rootCmd.Flags().Lookup(flagName))
	}
}

func (c *agentRootCommand) AddStringSliceProperty(name, flagName string, defaultVal []string, description string) {
	if c.rootCmd != nil {
		c.rootCmd.Flags().StringSlice(flagName, defaultVal, description)
		c.bindOrPanic(name, c.rootCmd.Flags().Lookup(flagName))
	}
}

func (c *agentRootCommand) AddDurationProperty(name, flagName string, defaultVal time.Duration, description string) {
	if c.rootCmd != nil {
		c.rootCmd.Flags().Duration(flagName, defaultVal, description)
		c.bindOrPanic(name, c.rootCmd.Flags().Lookup(flagName))
	}
}

func (c *agentRootCommand) AddIntProperty(name, flagName string, defaultVal int, description string) {
	if c.rootCmd != nil {
		c.rootCmd.Flags().Int(flagName, defaultVal, description)
		c.bindOrPanic(name, c.rootCmd.Flags().Lookup(flagName))
	}
}

func (c *agentRootCommand) AddBoolProperty(name, flagName string, defaultVal bool, description string) {
	if c.rootCmd != nil {
		c.rootCmd.Flags().Bool(flagName, defaultVal, description)
		c.bindOrPanic(name, c.rootCmd.Flags().Lookup(flagName))
	}
}

func (c *agentRootCommand) StringSlicePropertyValue(name string) []string {
	val := viper.Get(name)

	// special check to differentiate between yaml and commandline parsing. For commandline, must
	// turn it into an array ourselves
	switch val.(type) {
	case string:
		return c.convertStringToSlice(fmt.Sprintf("%v", viper.Get(name)))
	default:
		return viper.GetStringSlice(name)
	}
}

func (c *agentRootCommand) convertStringToSlice(value string) []string {
	slc := strings.Split(value, ",")
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}
	return slc
}

func (c *agentRootCommand) StringPropertyValue(name string) string {
	return viper.GetString(name)
}

func (c *agentRootCommand) DurationPropertyValue(name string) time.Duration {
	return viper.GetDuration(name)
}

func (c *agentRootCommand) IntPropertyValue(name string) int {
	return viper.GetInt(name)
}

func (c *agentRootCommand) BoolPropertyValue(name string) bool {
	return viper.GetBool(name)
}

func (c *agentRootCommand) GetAgentType() corecfg.AgentType {
	return c.agentType
}
