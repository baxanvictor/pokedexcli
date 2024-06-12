package commands

import (
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandWithArg[T any](
	commandName string,
	cache *pokecache.AppCache,
	command func(arg string, cache *pokecache.AppCache, dataHolder *T) error,
	dataHolder *T,
	args ...string,
) error {
	switch argsLen := len(args); argsLen {
	case 0:
		return fmt.Errorf("one argument must be provided with the \"%s\" command", commandName)
	case 1:
		return command(args[0], cache, dataHolder)
	default:
		return fmt.Errorf("the \"%s\" command only takes one argument", commandName)
	}
}
