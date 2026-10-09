package apperrors

import "errors"

var ErrWrongJsonFormat = errors.New("wrong json format")
var ErrEmptyModelName = errors.New("empty model name")
var ErrEmptyNodesNames = errors.New("empty nodes names")
var ErrNodesDuplicates = errors.New("nodes duplicates")
var ErrNoNodes = errors.New("no one nodes")
var ErrMismatchingNodes = errors.New("mismatching nodes")
var ErrModelNotFound = errors.New("model is not found")
var ErrProcessNotFound = errors.New("process is not found")

var ErrProcessTokenIsEmpty = errors.New("process token cannot be empty")

var ErrProcessCannotBeLocked = errors.New("process cannot be locked")

var ErrProcessCannotBeFinished = errors.New("process cannot be finished")

var ErrActualStateNotFound = errors.New("actual state for step hasn't been found")
