package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	proctor_config "github.com/gojektech/proctor/config"
	"github.com/gojektech/proctor/io"
	"github.com/spf13/cobra"
)

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func NewCmd(printer io.Printer) *cobra.Command {
	return &cobra.Command{
		Use:     "config",
		Short:   "Configure proctor client",
		Long:    "This command helps configure client with proctord host, email id and access token",
		Example: fmt.Sprintf("proctor config set %s=example.proctor.com %s=example@proctor.com %s=XXXXX", proctor_config.ProctorHost, proctor_config.EmailId, proctor_config.AccessToken),
		Args:    cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			configFile := filepath.Join(proctor_config.ConfigFileDir(), "proctor.yaml")
			if _, err := os.Stat(configFile); err == nil {
				printer.Println("[Warning] This will overwrite current config:", color.FgYellow)
				existingProctorConfig, err := ioutil.ReadFile(configFile)
				if err != nil {
					printer.Println(fmt.Sprintf("Error reading config file: %s", configFile), color.FgRed)
					return
				}

				printer.Println(string(existingProctorConfig))
				printer.Println("\nDo you want to continue (Y/n)?\t", color.FgYellow)

				in := bufio.NewReader(os.Stdin)
				userPermission, err := in.ReadString('\n')

				if err != nil {
					printer.Println("Error getting user permission for overwriting config", color.FgRed)
					return
				}

				if userPermission != "y\n" && userPermission != "Y\n" {
					printer.Println("Skipping configuring proctor client", color.FgYellow)
					return
				}
			}

			CreateDirIfNotExist(proctor_config.ConfigFileDir())
			var proctorHost, emailID, accessToken string
			for _, v := range args {
				arg := strings.Split(v, "=")

				if len(arg) != 2 {
					printer.Println(fmt.Sprintf("%-40s %-100s", "\nIncorrect config format: %s. Correct format: CONFIG_KEY=VALUE\n", v), color.FgRed)
					return
				}

				switch arg[0] {
				case proctor_config.ProctorHost:
					proctorHost = arg[1]
				case "EMAIL_ID":
					emailID = arg[1]
				case "ACCESS_TOKEN":
					accessToken = arg[1]
				default:
					printer.Println(fmt.Sprintf("Proctor doesn't support config key: %s", arg[0]), color.FgYellow)
				}
			}

			configFileContent := []byte(fmt.Sprintf("%s: "+proctorHost+"\n"+"%s: "+emailID+"\n"+"%s: "+accessToken, proctor_config.ProctorHost, proctor_config.EmailId, proctor_config.AccessToken))
			f, err := os.Create(configFile)
			if err != nil {
				printer.Println(fmt.Sprintf("Error creating config file %s: %s", configFile, err.Error()), color.FgRed)
			}
			_, err = f.Write(configFileContent)
			if err != nil {
				printer.Println(fmt.Sprintf("Error writing content %v \n to config file %s: %s", configFileContent, configFile, err.Error()), color.FgRed)
			}
			defer f.Close()
			printer.Println("Proctor client configured successfully", color.FgGreen)
		},
	}
}
