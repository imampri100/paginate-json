package configrequest

/*
ILogger

# Interface replicate from log4j java logging framework

Rankings: ALL < TRACE < DEBUG < INFO < WARN < ERROR < FATAL < OFF.

Reference:
- https://www.section.io/engineering-education/how-to-choose-levels-of-logging/
- https://www.tutorialspoint.com/log4j/log4j_logging_levels.htm
*/
type ILogger interface {
	LogAll(method, field string, value interface{})
	LogTrace(method, field string, value interface{})
	LogDebug(method, field string, value interface{})
	LogInfo(method, field string, value interface{})
	LogWarn(method, field string, value interface{})
	LogError(method, field string, value interface{})
	LogFatal(method, field string, value interface{})
	LogOff(method, field string, value interface{})
}
