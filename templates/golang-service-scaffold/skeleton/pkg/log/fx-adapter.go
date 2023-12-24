package log

import (
	"go.uber.org/fx/fxevent"
)

type (
	FxLogger struct {
		*Logger
	}
)

func GetFxLogger(logger *Logger) fxevent.Logger { //nolint:ireturn
	return FxLogger{logger}
}

func (l FxLogger) LogEvent(event fxevent.Event) { //nolint:cyclop
	switch e := event.(type) { //nolint:varnamelen
	case *fxevent.OnStartExecuting:
		l.logOnStartExecuting(e)
	case *fxevent.OnStartExecuted:
		l.logOnStartExecuted(e)
	case *fxevent.OnStopExecuting:
		l.logOnStopExecuting(e)
	case *fxevent.OnStopExecuted:
		l.logOnStopExecuted(e)
	case *fxevent.Supplied:
		l.logSupplied(e)
	case *fxevent.Provided:
		l.logProvided(e)
	case *fxevent.Decorated:
		l.logDecorated(e)
	case *fxevent.Invoking:
		l.logInvoking(e)
	case *fxevent.Started:
		l.logStarted(e)
	case *fxevent.LoggerInitialized:
		l.logLoggerInitialized(e)
	}
}

func (l *FxLogger) logOnStartExecuting(e *fxevent.OnStartExecuting) {
	l.Logger.Debug(
		"OnStart hook executing: ",
		String("callee", e.FunctionName),
		String("caller", e.CallerName),
	)
}

func (l *FxLogger) logOnStartExecuted(e *fxevent.OnStartExecuted) { //nolint:varnamelen
	if e.Err != nil {
		l.Logger.Debug(
			"OnStart hook failed: ",
			String("callee", e.FunctionName),
			String("caller", e.CallerName),
			ErrorObject(e.Err),
		)

		return
	}

	l.Logger.Debug(
		"OnStart hook executed: ",
		String("callee", e.FunctionName),
		String("caller", e.CallerName),
		String("runtime", e.Runtime.String()),
	)
}

func (l *FxLogger) logOnStopExecuting(e *fxevent.OnStopExecuting) {
	l.Logger.Debug(
		"OnStop hook executing: ",
		String("callee", e.FunctionName),
		String("caller", e.CallerName),
	)
}

func (l *FxLogger) logOnStopExecuted(e *fxevent.OnStopExecuted) { //nolint:varnamelen
	if e.Err != nil {
		l.Logger.Debug(
			"OnStop hook failed: ",
			String("callee", e.FunctionName),
			String("caller", e.CallerName),
			ErrorObject(e.Err),
		)

		return
	}

	l.Logger.Debug(
		"OnStop hook executed: ",
		String("callee", e.FunctionName),
		String("caller", e.CallerName),
		String("runtime", e.Runtime.String()),
	)
}

func (l *FxLogger) logSupplied(e *fxevent.Supplied) {
	l.Logger.Debug(
		"supplied: ",
		String("type", e.TypeName),
		ErrorObject(e.Err),
	)
}

func (l *FxLogger) logProvided(e *fxevent.Provided) {
	for _, rtype := range e.OutputTypeNames {
		l.Logger.Debug(
			"provided: ",
			String("constructor", e.ConstructorName),
			String("type", rtype),
		)
	}
}

func (l *FxLogger) logDecorated(e *fxevent.Decorated) {
	for _, rtype := range e.OutputTypeNames {
		l.Logger.Debug("decorated: ",
			String("decorator", e.DecoratorName),
			String("type", rtype),
		)
	}
}

func (l *FxLogger) logInvoking(e *fxevent.Invoking) {
	l.Logger.Debug(
		"invoking: ",
		String("function", e.FunctionName),
	)
}

func (l *FxLogger) logStarted(e *fxevent.Started) {
	if e.Err == nil {
		l.Logger.Debug("started")
	}
}

func (l *FxLogger) logLoggerInitialized(e *fxevent.LoggerInitialized) {
	if e.Err == nil {
		l.Logger.Debug(
			"initialized: custom fxevent.Logger",
			String("constructor", e.ConstructorName),
		)
	}
}
