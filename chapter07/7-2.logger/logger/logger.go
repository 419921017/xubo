package logger

type LogWritter interface {
	Write(data interface{}) error
}

type Logger struct {
	writerList []LogWritter
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) RegisterWriter(writer LogWritter) {
	l.writerList = append(l.writerList, writer)
}

func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}
