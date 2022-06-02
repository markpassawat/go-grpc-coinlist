package errx

import "reflect"

type matchOptions struct {
	ignoreFields bool
}

type matchFunc func(opts *matchOptions)

// IgnoreFields when use with Match will make Match ignore equality of fields.
func IgnoreFields(opts *matchOptions) {
	opts.ignoreFields = true
}

// Match matches err1 and err2. Return true if err1 and err2 can be considered
// the same errors, false otherwise. Optional match options can be provided.
func Match(err1, err2 error, opts ...matchFunc) bool {
	matchOpts := matchOptions{}
	for _, opt := range opts {
		opt(&matchOpts)
	}

	e1, ok := err1.(*Errorx)
	if !ok {
		return false
	}
	e2, ok := err2.(*Errorx)
	if !ok {
		return false
	}
	if e1.terr.StatusCode != 0 && e1.terr.StatusCode != e2.terr.StatusCode {
		return false
	}
	if e1.terr.Code != "" && e1.terr.Code != e2.terr.Code {
		return false
	}
	if !matchOpts.ignoreFields {
		if len(e1.terr.Fields) > 0 && !reflect.DeepEqual(e1.terr.Fields, e2.terr.Fields) {
			return false
		}
	}

	return true
}
