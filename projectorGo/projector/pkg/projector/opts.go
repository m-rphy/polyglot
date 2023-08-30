package projector

import (
	"errors"
	"go/parser"

	"github.com/hellflame/argparse"
)


type ProjectorOps struct {
    Pwd string
    Config string
    Args []string
}

func GetOptions() (*ProjectorOps, error) {
    parser := argparse.NewParser("projector", "get all the values", &argparse.ParserConfig{DisableDefaultShowHelp: true})
    args := parser.Strings("f", "foo", &argparse.Option{
        Positional: true,
        Default: "",
        Required: false,
    })

    config := parser.String("c", "config", &argparse.Option{
        Required: false,
        Default: "",
    })

    pwd := parser.String("p", "pwd", &argparse.Option{
        Required: false,
        Default: "",
    })

    err := parser.Parse(nil)
    if err != nil {
        return nil, err
    }

    return &ProjectorOps{
        Pwd: *pwd,
        Config: *config,
        Args: *args,
    }, nil
}
