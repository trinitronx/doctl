package commands

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bryanl/doit"
	"github.com/digitalocean/godo"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// DoitCmd is the base command.
	DoitCmd = &cobra.Command{
		Use: "doit",
	}

	// Token holds the global authorization token.
	Token string

	// Output holds the global output format.
	Output string

	requiredColor = color.New(color.Bold, color.FgWhite).SprintfFunc()

	writer = os.Stdout
)

func init() {
	viper.SetConfigType("yaml")

	DoitCmd.PersistentFlags().StringVarP(&Token, "access-token", "t", "", "DigtialOcean API V2 Access Token")
	DoitCmd.PersistentFlags().StringVarP(&Output, "output", "o", "text", "output formt [text|json]")
}

// LoadConfig loads out configuration.
func LoadConfig() error {
	cf := doit.NewConfigFile()
	r, err := cf.Open()
	if err != nil {
		return fmt.Errorf("can't open configuration file: %v", err)
	}

	return viper.ReadConfig(r)
}

// Execute executes the base command.
func Execute() {
	initializeConfig()
	addCommands()
	DoitCmd.Execute()
}

// AddCommands adds sub commands to the base command.
func addCommands() {
	DoitCmd.AddCommand(Account())
	DoitCmd.AddCommand(Actions())
	DoitCmd.AddCommand(Auth())
	DoitCmd.AddCommand(Domain())
	DoitCmd.AddCommand(DropletAction())
	DoitCmd.AddCommand(Droplet())
	DoitCmd.AddCommand(FloatingIP())
	DoitCmd.AddCommand(FloatingIPAction())
	DoitCmd.AddCommand(Images())
	DoitCmd.AddCommand(Region())
	DoitCmd.AddCommand(Size())
	DoitCmd.AddCommand(SSHKeys())
	DoitCmd.AddCommand(SSH())
	DoitCmd.AddCommand(Version())
}

func initFlags() {
	viper.SetEnvPrefix("DIGITALOCEAN")
	viper.BindEnv("access-token", "DIGITALOCEAN_ACCESS_TOKEN")
	viper.BindPFlag("access-token", DoitCmd.PersistentFlags().Lookup("access-token"))
	viper.BindPFlag("output", DoitCmd.PersistentFlags().Lookup("output"))
}

func loadDefaultSettings() {
	viper.SetDefault("output", "text")
}

// InitializeConfig initializes the doit configuration.
func initializeConfig() {
	loadDefaultSettings()
	LoadConfig()
	initFlags()

	if DoitCmd.PersistentFlags().Lookup("access-token").Changed {
		viper.Set("access-token", Token)
	}

	if DoitCmd.PersistentFlags().Lookup("output").Changed {
		viper.Set("output", Output)
	}
}

type flagOpt func(c *command, name, key string)

func requiredOpt() flagOpt {
	return func(c *command, name, key string) {
		c.MarkFlagRequired(key)
		key = requiredKey(key)
		viper.Set(key, true)

		u := c.Flag(name).Usage
		c.Flag(name).Usage = fmt.Sprintf("%s %s", u, requiredColor("(required)"))
	}
}

func requiredKey(key string) string {
	return fmt.Sprintf("%s.required", key)
}

func addStringFlag(cmd *command, name, dflt, desc string, opts ...flagOpt) {
	fn := flagName(cmd, name)
	cmd.Flags().String(name, dflt, desc)
	viper.BindPFlag(fn, cmd.Flags().Lookup(name))

	for _, o := range opts {
		o(cmd, name, fn)
	}
}

func addIntFlag(cmd *command, name string, def int, desc string, opts ...flagOpt) {
	fn := flagName(cmd, name)
	cmd.Flags().Int(name, def, desc)
	viper.BindPFlag(fn, cmd.Flags().Lookup(name))

	for _, o := range opts {
		o(cmd, name, fn)
	}
}

func addBoolFlag(cmd *command, name string, def bool, desc string, opts ...flagOpt) {
	fn := flagName(cmd, name)
	cmd.Flags().Bool(name, def, desc)
	viper.BindPFlag(fn, cmd.Flags().Lookup(name))

	for _, o := range opts {
		o(cmd, name, fn)
	}
}

func addStringSliceFlag(cmd *command, name string, def []string, desc string, opts ...flagOpt) {
	fn := flagName(cmd, name)
	cmd.Flags().StringSlice(name, def, desc)
	viper.BindPFlag(fn, cmd.Flags().Lookup(name))

	for _, o := range opts {
		o(cmd, name, fn)
	}
}

func flagName(cmd *command, name string) string {
	parentName := doit.NSRoot
	if cmd.Parent() != nil {
		parentName = cmd.Parent().Name()
	}

	return fmt.Sprintf("%s.%s.%s", parentName, cmd.Name(), name)
}

func cmdNS(cmd *cobra.Command) string {
	parentName := doit.NSRoot
	if cmd.Parent() != nil {
		parentName = cmd.Parent().Name()
	}

	return fmt.Sprintf("%s.%s", parentName, cmd.Name())
}

type cmdRunner func(ns string, config doit.Config, out io.Writer, args []string) error

type cmdOption func(*command)

type command struct {
	*cobra.Command

	fmtCols []string
}

func aliasOpt(aliases ...string) cmdOption {
	return func(c *command) {
		if c.Aliases == nil {
			c.Aliases = []string{}
		}

		for _, a := range aliases {
			c.Aliases = append(c.Aliases, a)
		}
	}
}

func displayerType(d displayer) cmdOption {
	return func(c *command) {
		c.fmtCols = d.Cols()
	}
}

func cmdBuilder(parent *cobra.Command, cr cmdRunner, cliText, desc string, out io.Writer, options ...cmdOption) *command {
	cc := &cobra.Command{
		Use:   cliText,
		Short: desc,
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			err := cr(cmdNS(cmd), doit.DoitConfig, out, args)
			checkErr(err, cmd)
		},
	}

	if parent != nil {
		parent.AddCommand(cc)
	}

	c := &command{Command: cc}

	for _, co := range options {
		co(c)
	}

	if cols := c.fmtCols; cols != nil {
		formatHelp := fmt.Sprintf("Columns for output in a comma seperated list. Possible values: %s",
			strings.Join(cols, ","))
		addStringFlag(c, doit.ArgFormat, "", formatHelp)
		addBoolFlag(c, doit.ArgNoHeader, false, "hide headers")
	}

	return c
}

func listDroplets(client *godo.Client) ([]godo.Droplet, error) {
	f := func(opt *godo.ListOptions) ([]interface{}, *godo.Response, error) {
		list, resp, err := client.Droplets.List(opt)
		if err != nil {
			return nil, nil, err
		}

		si := make([]interface{}, len(list))
		for i := range list {
			si[i] = list[i]
		}

		return si, resp, err
	}

	si, err := doit.PaginateResp(f)
	if err != nil {
		return nil, err
	}

	list := make([]godo.Droplet, len(si))
	for i := range si {
		list[i] = si[i].(godo.Droplet)
	}

	return list, nil
}

type dropletIPTable map[ifaceType]string

type ifaceType string

const (
	ifacePublic  ifaceType = "public"
	ifacePrivate ifaceType = "private"
)

func extractDropletIPs(droplet *godo.Droplet) dropletIPTable {
	t := dropletIPTable{}
	for _, in := range droplet.Networks.V4 {
		switch in.Type {
		case "public":
			t[ifacePublic] = in.IPAddress
		case "private":
			t[ifacePrivate] = in.IPAddress
		}
	}

	return t
}