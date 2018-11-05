package interfaces

var AllPluginFunctions []string = []string{"Init", "Version", "Methods"}

const (
	STAGE_LOAD_REPO = "stage_load_repo"
)

type PluginInOut struct {
	StringList  []string
	IntList     []int
	Float64List []float64
	Float32List []float32
	ByteList    []byte

	StringVal  string
	IntVal     int
	Float64Val float64
	Float32Val float32
	ByteVal    byte

	MSI   map[string]interface{}
	MSS   map[string]string
	Stage string
	Error error
}

// type PluginResult struct {
// 	StringList  []string
// 	IntList     []int
// 	Float64List []float64
// 	Float32List []float32
// 	ByteList    []byte

// 	StringVal  string
// 	IntVal     int
// 	Float64Val float64
// 	Float32Val float32
// 	ByteVal    byte

// 	Error error
// }

type PluginFunction func(PluginInOut) PluginInOut

type BasePlugin interface {
	Init(PluginInOut) error
	Name() string
	Version() string
	Methods() []string
	Run(string, ...interface{}) error
}

type Plugin interface {
	Name() string
	Version() string
	Methods() []string

	Run(string, ...interface{}) error
	Call(methodName string, vals ...interface{}) (res interface{}, err error)
	Int64(methodName string, vals ...interface{}) (res int64, err error)
	String(methodName string, vals ...interface{}) (res string, err error)
}

type Job interface {
	Run() error
}

// type JobLoader interface {
// 	LoadJob(Job) error
// }

type GlobalJob interface {
	Stage(string, error) error
	NextStep(...string) error
	Log(...interface{})
	DeleteDir()
	Dir(string)
	LoadJob(Job) error
}

type GlobalContext interface {
}

type PluginCreateFunction func() (Plugin, error)
