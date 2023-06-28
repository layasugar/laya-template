package gcnf

import "github.com/layasugar/laya/core/constants"

func AlarmType() string {
	if IsSet(constants.KEY_APPALARMTYPE) {
		return GetString(constants.KEY_APPALARMTYPE)
	}
	return constants.DEFAULT_NULLSTRING
}

func AlarmKey() string {
	if IsSet(constants.KEY_APPALARMKEY) {
		return GetString(constants.KEY_APPALARMKEY)
	}
	return constants.DEFAULT_NULLSTRING
}

func AlarmHost() string {
	if IsSet(constants.KEY_APPALARMADDR) {
		return GetString(constants.KEY_APPALARMADDR)
	}
	return constants.DEFAULT_NULLSTRING
}
