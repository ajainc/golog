package golog

// TextLogEvent
type TextLogEvent struct {
	Event string
}

// Encode implements LogEvent.Encode
func (logEvent *TextLogEvent) Encode(metadata *LogEventMetadata) []byte {
	if metadata != nil {

		data := metadata.GetLogLevel() + " " +
			metadata.GetTime() + " " +
			metadata.GetLoggerName() + " " +
			metadata.GetSourceFile() + "(" +
				metadata.GetSourceLine() + ") " +
					logEvent.Event

					return []byte(data)
	}


	return []byte(logEvent.Event)
}