package app_test

import (
	"net/rpc"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/app"
	"github.com/cloudfoundry/cli/cf/command_factory"
	"github.com/cloudfoundry/cli/cf/configuration/config_helpers"
	"github.com/cloudfoundry/cli/cf/manifest"
	"github.com/cloudfoundry/cli/cf/net"
	testconfig "github.com/cloudfoundry/cli/testhelpers/configuration"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"
	"github.com/codegangsta/cli"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Help", func() {
	BeforeEach(func() {
		rpc.DefaultServer = rpc.NewServer()
	})

	It("shows help for all commands", func() {
		commandFactory := createCommandFactory()

		dummyTemplate := `
{{range .Commands}}{{range .CommandSubGroups}}{{range .}}
{{.Name}}
{{end}}{{end}}{{end}}
`
		output := io_helpers.CaptureOutput(func() {
			app.ShowHelp(dummyTemplate, createApp(commandFactory))
		})

		for _, metadata := range commandFactory.CommandMetadatas() {
			Expect(commandInOutput(metadata.Name, output)).To(BeTrue(), metadata.Name+" not in help")
		}
	})

	It("shows help for all installed plugin's commands", func() {
		config_helpers.PluginRepoDir = func() string {
			return filepath.Join("..", "..", "fixtures", "config", "help-plugin-test-config")
		}

		commandFactory := createCommandFactory()

		dummyTemplate := `
{{range .Commands}}{{range .CommandSubGroups}}{{range .}}
{{.Name}}
{{end}}{{end}}{{end}}
`
		output := io_helpers.CaptureOutput(func() {
			app.ShowHelp(dummyTemplate, createApp(commandFactory))
		})

		Expect(commandInOutput("test_1_cmd1", output)).To(BeTrue(), "plugin command: test_1_cmd1 not in help")
		Expect(commandInOutput("test_1_cmd2", output)).To(BeTrue(), "plugin command: test_1_cmd2 not in help")
		Expect(commandInOutput("help", output)).To(BeTrue(), "plugin command: test_1_help not in help")
		Expect(commandInOutput("test_2_cmd1", output)).To(BeTrue(), "plugin command: test_2_cmd1 not in help")
		Expect(commandInOutput("test_2_cmd2", output)).To(BeTrue(), "plugin command: test_2_cmd2 not in help")

	})

})

func createCommandFactory() command_factory.Factory {
	fakeUI := &testterm.FakeUI{}
	configRepo := testconfig.NewRepository()
	manifestRepo := manifest.NewManifestDiskRepository()
	apiRepoLocator := api.NewRepositoryLocator(configRepo, map[string]net.Gateway{
		"auth":             net.NewUAAGateway(configRepo),
		"cloud-controller": net.NewCloudControllerGateway(configRepo, time.Now),
		"uaa":              net.NewUAAGateway(configRepo),
	})

	return command_factory.NewFactory(fakeUI, configRepo, manifestRepo, apiRepoLocator)
}

func createApp(commandFactory command_factory.Factory) *cli.App {
	new_app := cli.NewApp()
	new_app.Commands = []cli.Command{}
	for _, metadata := range commandFactory.CommandMetadatas() {
		new_app.Commands = append(new_app.Commands, cli.Command{Name: metadata.Name})
	}

	return new_app
}

func commandInOutput(cmdName string, output []string) bool {
	for _, line := range output {
		if strings.TrimSpace(line) == strings.TrimSpace(cmdName) {
			return true
		}
	}
	return false
}
