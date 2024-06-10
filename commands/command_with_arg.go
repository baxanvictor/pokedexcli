package commands

import (
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandWithArg[T any](
	commandName string,
	cache *pokecache.AppCache,
	command func(arg string, cache *pokecache.AppCache, response *T) error,
	response *T,
	args ...string,
) error {
	switch argsLen := len(args); argsLen {
	case 0:
		return fmt.Errorf("one argument must be provided with the \"%s\" command", commandName)
	case 1:
		return command(args[0], cache, response)
	default:
		return fmt.Errorf("the \"%s\" command only takes one argument", commandName)
	}
}
