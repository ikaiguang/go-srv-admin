package setup

import (
	"io"
	stdlog "log"
	"sync"

	writerutil "github.com/ikaiguang/go-srv-kit/kit/writer"
)

// LoggerFileWriter 文件日志写手柄
func (s *engines) LoggerFileWriter() (io.Writer, error) {
	var err error
	s.loggerFileWriterMutex.Do(func() {
		s.loggerFileWriter, err = s.loadingLoggerFileWriter()
		if err != nil {
			s.loggerFileWriterMutex = sync.Once{}
		}
	})
	return s.loggerFileWriter, err
}

// loadingLoggerFileWriter 启动日志文件写手柄
func (s *engines) loadingLoggerFileWriter() (io.Writer, error) {
	fileLoggerConfig := s.Config.LoggerConfigForFile()
	if !s.Config.EnableLoggingFile() || fileLoggerConfig == nil {
		stdlog.Println("|*** 加载：日志工具：虚拟的文件写手柄")
		return writerutil.NewDummyWriter()
	}
	rotateConfig := &writerutil.ConfigRotate{
		Dir:            fileLoggerConfig.Dir,
		Filename:       fileLoggerConfig.Filename,
		RotateTime:     fileLoggerConfig.RotateTime.AsDuration(),
		RotateSize:     fileLoggerConfig.RotateSize,
		StorageCounter: uint(fileLoggerConfig.StorageCounter),
		StorageAge:     fileLoggerConfig.StorageAge.AsDuration(),
	}
	if appConfig := s.Config.AppConfig(); appConfig != nil {
		if appConfig.Env != "" {
			rotateConfig.Filename += "_" + appConfig.Env
		}
		if appConfig.Version != "" {
			rotateConfig.Filename += "_" + appConfig.Version
		}
	}
	return writerutil.NewRotateFile(rotateConfig)
}
