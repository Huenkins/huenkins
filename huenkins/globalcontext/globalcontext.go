package globalcontext

import (
	"errors"
	"sync"
)

/*

	How to use:

	print.Printf(
		`String is: {{.StringData}}, Int is: {{.IntData}}, Bool is: {{.BoolData}}`,
		print.Add("StringData", "stringData", "IntData", 9999, "BoolData", true)
	)

*/

type Envirement struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`

	Key   string
	Value string
}

// JobStack is object for manager all jobs
type GlobalContext struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`

	Plugins     map[string][]*Plugin
	Credentials []*Credential
	Envirement  []*Envirement
}

func (gc *GlobalContext) GetPlugin(pluginID string, versionIDs ...string) (*Plugin, error) {
	/*
		This function
	*/

	listPlugins, find := gc.Plugins[pluginID]
	if !find {
		return nil, errors.New("Unknown plugin type.")
	}

	if len(listPlugins) == 0 {
		return nil, errors.New("Unknown plugin type.")
	}

	if len(versionIDs) == 0 {
		return listPlugins[len(listPlugins)-1], nil
	}

	return nil, nil
}
