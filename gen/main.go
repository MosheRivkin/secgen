package gen

import (
	"crypto/rand"
	"math/big"

	"github.com/mosherivkin/secgen/args"
)

func Gen(args args.Args) string {
	var chars string

	if args.Characters.Lowercase {
		chars += lowercase
	}
	if args.Characters.Uppercase {
		chars += uppercase
	}
	if args.Characters.Number {
		chars += number
	}
	if args.Characters.Special {
		chars += special
	}

	result := make([]byte, args.Length)
	for i := 0; i < args.Length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[n.Int64()]
	}

	return string(result)
}
