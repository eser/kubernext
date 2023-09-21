package log

import (
	"time"

	"github.com/eser/go-service/pkg/infra/config"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Field  = zapcore.Field
	Level  = zapcore.Level
	Logger struct {
		// *zapcore.Core
		*zap.Logger
	}
)

const (
	DebugLevel  = zapcore.DebugLevel
	InfoLevel   = zapcore.InfoLevel
	WarnLevel   = zapcore.WarnLevel
	ErrorLevel  = zapcore.ErrorLevel
	DPanicLevel = zapcore.DPanicLevel
	PanicLevel  = zapcore.PanicLevel
	FatalLevel  = zapcore.FatalLevel
)

var (
	Error      = zap.Error
	Int        = zap.Int
	String     = zap.String
	Duration   = zap.Duration
	Time       = zap.Time
	Any        = zap.Any
	WithCaller = zap.WithCaller
	L          = zap.L
)

func NewLogger(conf *config.Config) (*Logger, error) {
	var zapConf zap.Config

	if conf.Env == "production" {
		zapConf = zap.NewProductionConfig()
		zapConf.Level.SetLevel(InfoLevel)
		zapConf.DisableStacktrace = true
		zapConf.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		zapConf.EncoderConfig.TimeKey = "timestamp"
		zapConf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
	} else {
		zapConf = zap.NewDevelopmentConfig()
		zapConf.Level.SetLevel(DebugLevel)
		zapConf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapConf.EncoderConfig.TimeKey = "timestamp"
		zapConf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.TimeOnly)
		zapConf.EncoderConfig.ConsoleSeparator = " "
	}

	zapLogger, err := zapConf.Build()

	if err != nil {
		return nil, err
	}

	logger := &Logger{
		zapLogger.With(String("app", conf.AppName)).WithOptions(WithCaller(false)),
	}

	defer logger.Sync()

	return logger, nil
}

func GetFxLogger(logger *Logger) fxevent.Logger {
	return logger
}

func (l *Logger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Debug(
			"OnStart hook executing: ",
			String("callee", e.FunctionName),
			String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Debug(
				"OnStart hook failed: ",
				String("callee", e.FunctionName),
				String("caller", e.CallerName),
				Error(e.Err),
			)
		} else {
			l.Logger.Debug(
				"OnStart hook executed: ",
				String("callee", e.FunctionName),
				String("caller", e.CallerName),
				String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Debug(
			"OnStop hook executing: ",
			String("callee", e.FunctionName),
			String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Debug(
				"OnStop hook failed: ",
				String("callee", e.FunctionName),
				String("caller", e.CallerName),
				Error(e.Err),
			)
		} else {
			l.Logger.Debug(
				"OnStop hook executed: ",
				String("callee", e.FunctionName),
				String("caller", e.CallerName),
				String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.Logger.Debug(
			"supplied: ",
			String("type", e.TypeName),
			Error(e.Err),
		)
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug(
				"provided: ",
				String("constructor", e.ConstructorName),
				String("type", rtype),
			)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("decorated: ",
				String("decorator", e.DecoratorName),
				String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.Logger.Debug(
			"invoking: ",
			String("function", e.FunctionName),
		)
	case *fxevent.Started:
		if e.Err == nil {
			l.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.Logger.Debug(
				"initialized: custom fxevent.Logger",
				String("constructor", e.ConstructorName),
			)
		}
	}
}
