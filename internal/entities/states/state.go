package states

import "slices"

const (
	STATE_NONE    string = "NONE"
	STATE_RUN     string = "RUN"
	STATE_SKIP    string = "SKIP"
	STATE_TIMEOUT string = "TIMEOUT"
	STATE_DONE    string = "DONE"
	STATE_ERROR   string = "ERROR"
)

func CheckStateIsValid(stateName string) bool {
	states := []string{STATE_NONE, STATE_RUN, STATE_SKIP, STATE_TIMEOUT, STATE_DONE, STATE_ERROR}
	res := slices.Contains(states, stateName)
	return res
}

func CheckStateSequence(stateName string, prevStateName string) bool {
	if stateName == STATE_NONE {
		return true
	}

	if prevStateName == STATE_NONE && (stateName == STATE_RUN || stateName == STATE_SKIP || stateName == STATE_TIMEOUT) {
		return true
	}

	if prevStateName == STATE_RUN && (stateName == STATE_DONE || stateName == STATE_ERROR || stateName == STATE_TIMEOUT) {
		return true
	}

	return false
}

func CheckStateIsFinishState(stateName string) bool {
	finishStates := []string{STATE_SKIP, STATE_TIMEOUT, STATE_DONE, STATE_ERROR}
	res := slices.Contains(finishStates, stateName)
	return res
}
