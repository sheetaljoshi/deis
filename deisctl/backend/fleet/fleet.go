package fleet

import (
	"io"
	"os"
	"path"
	"text/tabwriter"

	"github.com/deis/deis/deisctl/config"

	"github.com/coreos/fleet/client"
	"github.com/coreos/fleet/machine"
)

// FleetClient used to wrap Fleet API calls
type FleetClient struct {
	Fleet         client.API
	configBackend config.Backend

	// used to cache MachineStates
	machineStates map[string]*machine.MachineState

	templatePaths []string
	runner        commandRunner
	out           *tabwriter.Writer
	errWriter     io.Writer
}

// NewClient returns a client used to communicate with Fleet
// using the Registry API
func NewClient(cb config.Backend) (*FleetClient, error) {
	client, err := getRegistryClient()
	if err != nil {
		return nil, err
	}

	// path hierarchy for finding systemd service templates
	templatePaths := []string{
		os.Getenv("DEISCTL_UNITS"),
		path.Join(os.Getenv("HOME"), ".deis", "units"),
		"/var/lib/deis/units",
	}

	out := new(tabwriter.Writer)
	out.Init(os.Stdout, 0, 8, 1, '\t', 0)

	return &FleetClient{Fleet: client, configBackend: cb, templatePaths: templatePaths, runner: sshCommandRunner{},
		out: out, errWriter: os.Stderr}, nil
}
