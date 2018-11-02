package plugins

/*
	"plugins" provides interface for:
	- loading plugins
	- executing plugins
*/

import (
	"fmt"
	"reflect"

	"github.com/iostrovok/huenkins/huenkins/interfaces"
)

func _recover(err *error, methodName string) {
	//handle panic
	if v := recover(); v != nil {
		*err = fmt.Errorf("Error %s for %s\n", v, methodName)
	}
}

func (pl *Plugin) _call(methodName string, p interfaces.PluginInOut) (res interface{}, err error) {

	args := []reflect.Value{reflect.ValueOf(interfaces.PluginInOut{})}

	// []reflect.Value
	out := reflect.ValueOf(pl.Object).MethodByName(methodName).Call(args)

	if len(out) != 1 {
		err = fmt.Errorf("Wrogn [ %d != 1 ] length of returned parameters for %s\n", len(out), methodName)
	} else {
		res = out[0].Interface()
	}

	fmt.Printf("_call: resresresresresresres: %+v\n", res)
	fmt.Printf("_call: resresresresresresres: %T\n", res)

	return
}

func (pl *Plugin) _YYYCallToInt(methodName string, val reflect.Value, itype reflect.Type) (res interface{}, err error) {

	pl.mu.RLock()
	defer pl.mu.RUnlock()
	defer _recover(&err, methodName)

	if err == nil {
		switch t := val.(type) {
		case itype:
			res = val.(itype)
		default:
			err = fmt.Errorf("Wrogn [%s] type of returned parameters for %s", t, methodName)
		}
	}

	return
}

func (pl *Plugin) Call(methodName string, p interfaces.PluginInOut) (res interfaces.PluginInOut, err error) {

	pl.mu.RLock()
	defer pl.mu.RUnlock()
	defer _recover(&err, methodName)

	var val interface{}
	val, err = pl._call(methodName, p)
	if err == nil {
		switch t := val.(type) {
		case interfaces.PluginInOut:
			res = val.(interfaces.PluginInOut)
		default:
			err = fmt.Errorf("Wrogn [%s] type of returned parameters for %s", t, methodName)
		}
	}

	return
}

func (pl *Plugin) CallToInt(methodName string, p interfaces.PluginInOut) (res int64, err error) {

	pl.mu.RLock()
	defer pl.mu.RUnlock()
	defer _recover(&err, methodName)

	var val interface{}
	val, err = pl._call(methodName, p)
	if err == nil {
		switch t := val.(type) {
		case int64:
			res = val.(int64)
		default:
			err = fmt.Errorf("Wrogn [%s] type of returned parameters for %s", t, methodName)
		}
	}

	return
}

func (pl *Plugin) CallToString(methodName string, p interfaces.PluginInOut) (res string, err error) {

	pl.mu.RLock()
	defer pl.mu.RUnlock()
	defer _recover(&err, methodName)

	var val interface{}
	val, err = pl._call(methodName, p)
	if err == nil {
		switch t := val.(type) {
		case string:
			res = val.(string)
		default:
			err = fmt.Errorf("Wrogn [%s] type of returned parameters for %s", t, methodName)
		}
	}

	return
}
