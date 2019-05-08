/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var (
	// Log 日志全局对象
	Log Logger
	// Common 通用包日志对象
	Common *zap.Logger
	// Discovery 发现服务包日志对象
	Discovery *zap.Logger
	// Examples 案例日志对象
	Examples *zap.Logger
	// Rivet 框架日志对象
	Rivet *zap.Logger
	// Server 关联接口服务日志对象
	Server *zap.Logger
	// Shunt 负载均衡日志对象
	Shunt *zap.Logger
	// Trans 请求处理日志对象
	Trans *zap.Logger
)

const (
	// DebugLevel 日志级别为 debug
	DebugLevel = "debug"
	// InfoLevel 日志级别为 info
	InfoLevel = "info"
)

// Logger 日志入口对象
type Logger struct {
	Config *Config
}

func init() {
	Log = Logger{}
	Log.Conf(&Config{
		FilePath:   strings.Join([]string{"./logs/rivet.log"}, ""),
		Level:      zapcore.DebugLevel,
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	})
	Common = Log.New("./logs/common.log", "common")
	Discovery = Log.New("./logs/discovery.log", "discovery")
	Examples = Log.New("./logs/examples.log", "examples")
	Rivet = Log.New("./logs/rivet.log", "rivet")
	Server = Log.New("./logs/server.log", "server")
	Shunt = Log.New("./logs/shunt.log", "shunt")
	Trans = Log.New("./logs/trans.log", "trans")
}

// Conf 配置日志基本信息
func (log *Logger) Conf(config *Config) {
	log.Config = config
}

// New 新建日志对象
func (log *Logger) New(filePath string, serviceName string) *zap.Logger {
	core := newCore(filePath, log.Config.Level, log.Config.MaxSize, log.Config.MaxBackups, log.Config.MaxAge, log.Config.Compress)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", serviceName))
	// 返回构造日志
	return zap.New(core, caller, development, filed)
}

// NewCustom 新建自定义日志对象
func (log *Logger) NewCustom(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", serviceName))
	// 返回构造日志
	return zap.New(core, caller, development, filed)
}

func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
}
