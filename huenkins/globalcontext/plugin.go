package globalcontext

import (
	"fmt"
	"plugin"
	"sync"
)

type Plugin struct {
	// safeness :)
	mu sync.RWMutex `json:"-"`

	FileSO     string `json:"file_so"`
	PluginType string `json:"plugin_type"`
}

func NewPlugin(globalContext *GlobalContext, pluginType, fileSO string) (*Plugin, error) {
	pl := &Plugin{
		FileSO:     fileSO,
		PluginType: pluginType,
	}

	return pl, nil
}

// Main function
func (pl *Plugin) CheckCredential(globalContext *GlobalContext, data map[string][]byte) (map[string]string, map[string][]byte, error) {
	return map[string]string{}, map[string][]byte{}, nil
}

// Main function
func (pl *Plugin) Load(globalContext *GlobalContext) error {

	p, err := plugin.Open(pl.FileSO)
	if err != nil {
		panic(err)
	}

	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	initFunc, err := p.Lookup("Init")
	if err != nil {
		panic(err)
	}

	// t, err := p.Lookup("T")
	// if err != nil {
	// 	panic(err)
	// }

	*v.(*GlobalContext) = *globalContext
	f.(func())() // prints "Hello, number 7"

	// out := t.(func() int)() // prints "Hello, number 7"
	// fmt.Printf("Hello, out %d\n", out)
	fmt.Printf("Hello, initFunc: %+v\n", initFunc)
	fmt.Printf("Hello from Load\n")

	return nil
}
