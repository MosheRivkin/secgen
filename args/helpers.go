package args

import (
	"errors"
	"strconv"
)

func setDefaultArgs(args []string) []string {
	if len(args) == 0 {
		return []string{defaultLength, defaultFlags}
	}
	if len(args) == 1 {
		if isInteger(args[0]) {
			return append(args, defaultFlags)
		}
		return append([]string{defaultLength}, args[0])
	}
	return args
}

func parseLength(length string) (int, error) {
	value, err := strconv.Atoi(length)
	if err != nil {
		return 0, errors.New(errInvalidLen)
	}
	if value <= 0 {
		return 0, errors.New(errInvalidLenSign)
	}
	return value, nil
}

func parseCharacterFlags(flags string) (*Characters, error) {

	invertFlags := flags[0] == '!'
	if invertFlags {
		flags = flags[1:]
	}

	charSet := createCharacters(invertFlags)

	for _, flag := range flags {
		switch flag {
		case 'n':
			charSet.Number = !invertFlags
		case 'l':
			charSet.Lowercase = !invertFlags
		case 'u':
			charSet.Uppercase = !invertFlags
		case 's':
			charSet.Special = !invertFlags
		case '!':
			continue
		default:
			return nil, errors.New(errInvalidFlag + string(flag))

		}
	}

	return charSet, nil
}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func createCharacters(initialValue bool) *Characters {
	return &Characters{
		Number:    initialValue,
		Lowercase: initialValue,
		Uppercase: initialValue,
		Special:   initialValue,
	}
}
