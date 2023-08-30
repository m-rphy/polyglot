package projector

import (
	"github.com/hellflame/argparse"
)


type Opts struct {
    Pwd string
    Config string
    Args []string
}

func GetOpts() (*Opts, error) {
    parser := argparse.NewParser("projector", "get all the values", &argparse.ParserConfig{DisableDefaultShowHelp: true})
    
    args := parser.Strings("a", "args", &argparse.Option{
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

    return &Opts{
        Pwd: *pwd,
        Config: *config,
        Args: *args,
    }, nil
}
