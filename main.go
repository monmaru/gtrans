package main

import (
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

const apiKey = "YOUR_TRANSLATE_API_KEY"

func main() {
	app := cli.NewApp()
	app.Name = "gtrans"
	app.Version = "1.0"
	app.Usage = "Google Translate API from the terminal"
	app.Action = func(c *cli.Context) error {
		if c.NArg() < 2 {
			fmt.Println("Ex: gtrans en|jp target text")
			return nil
		}

		lang := c.Args().Get(0)
		text := strings.Join(c.Args()[1:], " ")
		ret, err := trans(lang, text)
		if err != nil {
			return err
		}
		fmt.Println(ret)
		return nil
	}
	app.Run(os.Args)
}

func trans(lang, text string) (string, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	tag, err := language.Parse(lang)
	if err != nil {
		return "", err
	}

	resp, err := client.Translate(ctx, []string{text}, tag, nil)
	if err != nil {
		return "", err
	}

	return resp[0].Text, nil
}
