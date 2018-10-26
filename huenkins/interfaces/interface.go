package interfaces

var AllPluginFunctions []string = []string{"Init", "Version", "Methods"}

type GlobalContext struct {
}

type PluginParameter struct {
	StringList  []string
	IntList     []int
	Float64List []float64
	Float32List []float32

	StringVal  string
	IntVal     int
	Float64Val float64
	Float32Val float32

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
