package url

import (
	"context"
	"fmt"
	"strings"

	urlpkg "net/url"

	"github.com/pkg/browser"
	"github.com/urfave/cli/v3"
)

type Providers map[string]string

func (p Providers) In(s string) (string, bool) {
	v, ok := p[s]
	if !ok {
		return "", false
	}

	return v, true
}

// TODO: will move to configuration file
var providers = Providers{
	"gg":          "https://www.google.com/search?q=%s",
	"ggtranslate": "https://translate.google.com/?sl=en&tl=th&text=%s&op=translate",
	"gmail":       "https://mail.google.com/mail/u/%s",
	"gh":          "https://github.com/search?q=%s",
	"mygh":        "https://github.com/imkk000?tab=repositories&q=%s",
	"arch":        "https://archlinux.org/packages/?q=%s",
	"aur":         "https://aur.archlinux.org/packages?K=%s",
	"archwiki":    "https://wiki.archlinux.org/index.php?search=%s",
	"yt":          "https://www.youtube.com/results?search_query=%s",
	"gopkg":       "https://pkg.go.dev/search?q=%s",
	"leetcode":    "https://leetcode.com/search/?q=%s",
	"rfc":         "https://www.rfc-editor.org/search/rfc_search_detail.php?title=%s",
	"reddit":      "https://www.reddit.com/search/?q=%s",
	"docker":      "https://hub.docker.com/search?q=%s",
	"cve":         "https://www.cve.org/CVERecord/SearchResults?query=%s",
	"wiki":        "https://en.wikipedia.org/w/index.php?search=%s",
	"mozilla":     "https://developer.mozilla.org/en-US/search?q=%s",
	"imdb":        "https://www.imdb.com/find/?q=%s",
	"opensub":     "https://www.opensubtitles.com/en/en/search-all/q-%s/hearing_impaired-exclude/machine_translated-/trusted_sources-",
	"fish":        "https://fishshell.com/docs/current/search.html?q=%s",
	"cheat":       "https://cheat.sh/%s",
}

func buildProvidersCommands() []*cli.Command {
	cmd := make([]*cli.Command, 0, len(providers))

	for alias := range providers {
		cmd = append(cmd, &cli.Command{
			Name: alias,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "provider",
					Value: alias,
				},
				&cli.BoolFlag{
					Name:    "web",
					Aliases: []string{"w"},
					Value:   true,
				},
			},
			Action: func(_ context.Context, c *cli.Command) error {
				provider := c.String("provider")
				format, valid := providers.In(provider)
				if !valid {
					return fmt.Errorf("provider %s: not found", provider)
				}

				keywords := strings.Join(c.Args().Slice(), " ")
				q := urlpkg.QueryEscape(keywords)
				fullURL := fmt.Sprintf(format, q)

				if !c.Bool("web") {
					fmt.Println(fullURL)

					return nil
				}
				if err := browser.OpenURL(fullURL); err != nil {
					return fmt.Errorf("open browser: %w", err)
				}

				return nil
			},
		})
	}

	return cmd
}
