package gcnf

import "github.com/layasugar/laya/core/constants"

func TraceType() string {
	if IsSet(constants.KEY_APPTRACETYPE) {
		return GetString(constants.KEY_APPTRACETYPE)
	}
	return constants.DEFAULT_TRACETYPE
}

func TraceAddr() string {
	if IsSet(constants.KEY_APPTRACEADDR) {
		return GetString(constants.KEY_APPTRACEADDR)
	}
	return constants.DEFAULT_TRACEADDR
}

func TraceMod() float64 {
	if IsSet(constants.KEY_APPTRACEMOD) {
		return GetFloat64(constants.KEY_APPTRACEMOD)
	}
	return constants.DEFAULT_TRACEMOD
}
