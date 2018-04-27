package golog

import (
	"os"
	"fmt"
	"io"
)

// Logger
type Logger struct {
	// LoggerName
	// Public Required
	//
	// It can be used as an identifier when outputting logs.
	// If you specify an empty string, the default logger name is substituted
	Name string

	// levelAppender
	// Private Required
	levelAppender map[LogLevel][]Appender

	// enabledMetadata
	// Private Required
	enabledMetadata bool

	// metadataFormatter
	// Private Option
	//
	// If not specified, the default formatter will be used
	metadataFormatter *MetadataFormatter

	// metadataConfig
	// Private Option
	//
	// If not specified, the default config wil be used
	metadataConfig *MetadataConfig
}

// doAppendIfLevelEnabled
func (logger *Logger) doAppendIfLevelEnabled(event []byte, level LogLevel) {

	// recover
	defer func(writer io.Writer) {
		if err := recover(); err != nil {
			fmt.Fprintln(writer, "Error: golog exit appending error:", err)
		}
	}(os.Stderr)

	if appenders, ok := logger.levelAppender[level]; ok {
		for _, appender := range appenders {
			appender.Write(event)
		}
	}
}

// newMetadata
func (logger *Logger) newMetadata(level LogLevel) LogEventMetadata {
	var metadata LogEventMetadata
	metadata = NewLogEventMetadata(logger.metadataConfig, logger.metadataFormatter)
	metadata.setLogLevel(level)
	metadata.setLoggerName(logger.Name)
	metadata.setSource(5)
	return metadata
}

// Trace calls specified appender to print string.
func (logger *Logger) Trace(string string) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_TRACE)
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(&metadata), LogLevel_TRACE)
	} else {
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(nil), LogLevel_TRACE)
	}
}

// Debug calls specified appender to print string.
func (logger *Logger) Debug(string string) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_DEBUG)
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(&metadata), LogLevel_DEBUG)
	} else {
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(nil), LogLevel_DEBUG)
	}
}

// Info calls specified appender to print string.
func (logger *Logger) Info(string string) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_INFO)
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(&metadata), LogLevel_INFO)
	} else {
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(nil), LogLevel_INFO)
	}
}

// Warn calls specified appender to print string.
func (logger *Logger) Warn(string string) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_WARN)
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(&metadata), LogLevel_WARN)
	} else {
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(nil), LogLevel_WARN)
	}
}

// Error calls specified appender to print string.
func (logger *Logger) Error(string string) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_ERROR)
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(&metadata), LogLevel_ERROR)
	} else {
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(nil), LogLevel_ERROR)
	}
}

// Fatal calls specified appender to print string.
func (logger *Logger) Fatal(string string) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_FATAL)
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(&metadata), LogLevel_FATAL)
	} else {
		logger.doAppendIfLevelEnabled(TextLogEvent{Event: string}.Encode(nil), LogLevel_FATAL)
	}

	os.Exit(1)
}

// TraceF encodes according to format specifier and calls specified appender to print.
func (logger *Logger) TraceF(format string, args ...interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_TRACE)
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(&metadata), LogLevel_TRACE)
	} else {
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_TRACE)
	}
}

// DebugF encodes according to format specifier and calls specified appender to print.
func (logger *Logger) DebugF(format string, args ...interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_DEBUG)
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(&metadata), LogLevel_DEBUG)
	} else {
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_DEBUG)
	}
}

// InfoF encodes according to format specifier and calls specified appender to print.
func (logger *Logger) InfoF(format string, args ...interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_INFO)
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(&metadata), LogLevel_INFO)
	} else {
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_INFO)
	}
}

// WarnF encodes according to format specifier and calls specified appender to print.
func (logger *Logger) WarnF(format string, args ...interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_WARN)
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(&metadata), LogLevel_WARN)
	} else {
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_WARN)
	}
}

// ErrorF encodes according to format specifier and calls specified appender to print.
func (logger *Logger) ErrorF(format string, args ...interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_ERROR)
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(&metadata), LogLevel_ERROR)
	} else {
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_ERROR)
	}
}

// FatalF encodes according to format specifier and calls specified appender to print.
func (logger *Logger) FatalF(format string, args ...interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_FATAL)
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(&metadata), LogLevel_FATAL)
	} else {
		logger.doAppendIfLevelEnabled(FormatLogEvent{format: format, args: args,}.Encode(nil), LogLevel_FATAL)
	}
}

// TraceJ encodes as Json binary and calls specified appender to print.
func (logger *Logger) TraceJ(obj interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_TRACE)
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(&metadata), LogLevel_TRACE)
	} else {
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_TRACE)
	}
}

// DebugJ encodes as Json binary and calls specified appender to print.
func (logger *Logger) DebugJ(obj interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_DEBUG)
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(&metadata), LogLevel_DEBUG)
	} else {
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_DEBUG)
	}
}

// InfoJ encodes as Json binary and calls specified appender to print.
func (logger *Logger) InfoJ(obj interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_INFO)
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(&metadata), LogLevel_INFO)
	} else {
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_INFO)
	}
}

// WarnJ encodes as Json binary and calls specified appender to print.
func (logger *Logger) WarnJ(obj interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_WARN)
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(&metadata), LogLevel_WARN)
	} else {
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_WARN)
	}
}

// ErrorJ encodes as Json binary and calls specified appender to print.
func (logger *Logger) ErrorJ(obj interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_ERROR)
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(&metadata), LogLevel_ERROR)
	} else {
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_ERROR)
	}
}

// FatalJ encodes as Json binary and calls specified appender to print.
func (logger *Logger) FatalJ(obj interface{}) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_FATAL)
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(&metadata), LogLevel_FATAL)
	} else {
		logger.doAppendIfLevelEnabled(JsonLogEvent{event: obj,}.Encode(nil), LogLevel_FATAL)
	}
}

// TraceS encodes as user defined logEvent and calls specified appender to print it.
func (logger *Logger) TraceS(logEvent LogEvent) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_TRACE)
		logger.doAppendIfLevelEnabled(logEvent.Encode(&metadata), LogLevel_TRACE)
	} else {
		logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_TRACE)
	}
}

// DebugS encodes as user defined logEvent and calls specified appender to print it.
func (logger *Logger) DebugS(logEvent LogEvent) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_DEBUG)
		logger.doAppendIfLevelEnabled(logEvent.Encode(&metadata), LogLevel_DEBUG)
	} else {
		logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_DEBUG)
	}
}

// InfoS encodes as user defined logEvent and calls specified appender to print it.
func (logger *Logger) InfoS(logEvent LogEvent) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_INFO)
		logger.doAppendIfLevelEnabled(logEvent.Encode(&metadata), LogLevel_INFO)
	} else {
		logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_INFO)
	}
}

// WarnS encodes as user defined logEvent and calls specified appender to print it.
func (logger *Logger) WarnS(logEvent LogEvent) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_WARN)
		logger.doAppendIfLevelEnabled(logEvent.Encode(&metadata), LogLevel_WARN)
	} else {
		logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_WARN)
	}
}

// ErrorS encodes as user defined logEvent and calls specified appender to print it.
func (logger *Logger) ErrorS(logEvent LogEvent) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_ERROR)
		logger.doAppendIfLevelEnabled(logEvent.Encode(&metadata), LogLevel_ERROR)
	} else {
		logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_ERROR)
	}
}

// FatalS encodes as user defined logEvent and calls specified appender to print it.
func (logger *Logger) FatalS(logEvent LogEvent) {
	if logger.enabledMetadata {
		metadata := logger.newMetadata(LogLevel_FATAL)
		logger.doAppendIfLevelEnabled(logEvent.Encode(&metadata), LogLevel_FATAL)
	} else {
		logger.doAppendIfLevelEnabled(logEvent.Encode(nil), LogLevel_FATAL)
	}
}

// SetAppender
func (logger *Logger) SetAppender(appender ...Appender) {
	for k := range logger.levelAppender {
		logger.levelAppender[k] = appender
	}
}

// DisableLogEventMetadata
// If metadata is unnecessary, please disable it.
// It is possible to prevent unnecessary allocation.
// It is enabled by default.
func (logger *Logger) DisableLogEventMetadata() {
	logger.enabledMetadata = false
}

// SetMetadataFormatter
func (logger *Logger) SetMetadataFormatter(formatter *MetadataFormatter) {
	logger.metadataFormatter = formatter
}

// SetMetadataConfig
func (logger *Logger) SetMetadataConfig(config *MetadataConfig) {
	logger.metadataConfig = config
}

// SetLogLevel enables the specified log level
func (logger *Logger) SetAppenderWithLevel(logLevel LogLevel, appender ...Appender) {
	logger.levelAppender[logLevel] = appender
}

// SetLogLevel enables the specified log level
func (logger *Logger) SetAppenderWithLevels(logLevels []LogLevel, appender ...Appender) {
	for _, v := range logLevels {
		logger.levelAppender[v] = appender
	}
}

// Close implements io.Closer
func (logger *Logger) Close() error {

	for _,v := range logger.levelAppender {
		for _, appender := range v {
			err := appender.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "close appender is failed , error : %s\n",  err.Error())
			}
		}
	}
	return nil
}

// NewLogger
func NewLogger(loggerName string, logLevel LogLevel, appender ...Appender) Logger {
	levelAppender := map[LogLevel][]Appender{}
	logLevels := NewDefaultLevelFilter().DoFilter(logLevel)

	if len(logLevels) == 0{
		fmt.Fprintln(os.Stderr, "Warn: no levels is specified")
	}

	if len(appender) == 0 {
		fmt.Fprintln(os.Stderr, "Warn: no appender is specified")
	}

	for _, logLevel := range logLevels {
		levelAppender[logLevel] = appender
	}

	return Logger{
		Name:            loggerName,
		levelAppender:   levelAppender,
		enabledMetadata: true,
	}
}

// NewDefaultLogger
func NewDefaultLogger() Logger {
	return NewLogger("defaultLogger", LogLevel_TRACE, NewDefaultConsoleAppender())
}
