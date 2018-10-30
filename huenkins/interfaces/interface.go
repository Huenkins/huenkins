package interfaces

var AllPluginFunctions []string = []string{"Init", "Version", "Methods"}

const (
	STAGE_LOAD_REPO = "stage_load_repo"
)

type GlobalContext struct {
}

type PluginParameter struct {
	StringList  []string
	IntList     []int
	Float64List []float64
	Float32List []float32
	ByteList    []byte

	StringVal  string
	IntVal     int
	Float64Val float64
	Float32Val float32
	ByteVal    []byte

	MSI map[string]interface{}
}

type PluginResult struct {
	StringList  []string
	IntList     []int
	Float64List []float64
	Float32List []float32

	StringVal  string
	IntVal     int
	Float64Val float64
	Float32Val float32

	Error error
}

type PluginFunction func(PluginParameter) PluginResult

type Plugin interface {
	Init(PluginParameter) PluginResult
	Version(PluginParameter) PluginResult
	Methods(PluginParameter) PluginResult
}

type PluginCreateFunction func() (Plugin, error)
