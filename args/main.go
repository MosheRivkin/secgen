package args

import (
	"errors"
)

func Parse(args []string) (*Args, error) {
	args = setDefaultArgs(args)

	if len(args) > 2 {
		return nil, errors.New(errTooManyArgs)
	}

	length, err := parseLength(args[0])
	if err != nil {
		return nil, err
	}

	charSet, err := parseCharacterFlags(args[1])
	if err != nil {
		return nil, err
	}

	return &Args{Length: length, Characters: charSet}, nil
}
