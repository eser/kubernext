package log

type RuntimeContext struct {
	HasLoggerSet bool
}

func NewRuntimeContext() (*RuntimeContext, error) {
	runtimeContext := &RuntimeContext{}

	return runtimeContext, nil
}
