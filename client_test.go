package teams_api_test

import (
	"bytes"
	"fmt"
	teams_api "github.com/fossteams/teams-api"
	"github.com/fossteams/teams-api/pkg/csa"
	"github.com/logrusorgru/aurora"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"sort"
	"strings"
	"testing"
)

func TestTeamsClient_GetConversations(t *testing.T) {
	c, err := teams_api.New()
	if err != nil {
		t.Fatalf("unable to create teams client: %v", err)
	}

	c.Debug(true)
	convs, err := c.GetConversations()
	if err != nil {
		t.Fatalf("unable to get conversations: %v", err)
	}

	if convs == nil {
		t.Fatal("convs should never be nil!")
	}

	// Pretty print conversations
	fmt.Printf("%s\n", aurora.Bold("Teams"))
	sort.Sort(csa.TeamsByName(convs.Teams))
	for _, t := range convs.Teams {
		fmt.Printf("%s (%d users)\n",
			aurora.Magenta(t.DisplayName),
			aurora.Green(t.MembershipSummary.UserRoleCount),
		)
		sort.Sort(csa.ChannelsByName(t.Channels))
		for _, channel := range t.Channels {
			fmt.Printf("\t%s\n",
				aurora.Red(channel.DisplayName))
		}
		fmt.Printf("\n")
	}
}

func TestTeamsClient_GetMessages(t *testing.T) {
	c, err := teams_api.New()
	if err != nil {
		t.Fatalf("unable to create teams client: %v", err)
	}

	c.Debug(true)
	convs, err := c.GetConversations()
	if err != nil {
		t.Fatalf("unable to get conversations: %v", err)
	}

	if convs == nil {
		t.Fatal("convs should never be nil!")
	}

	// Get first team, first channel
	team := convs.Teams[0]
	channel := team.Channels[0]
	assert.NotNil(t, channel)

	fmt.Printf("%s\n", aurora.Red(team.DisplayName))
	fmt.Printf("%s\n", aurora.Yellow(channel.DisplayName))

	messages, err := c.GetMessages(&channel)
	if err != nil {
		t.Fatalf("unable to get channel messages: %v", err)
	}
	assert.Greater(t, len(messages), 0)

	for _, m := range messages {

		fmt.Printf("%s\n", aurora.Bold(aurora.Green(m.ImDisplayName)))
		z := html.NewTokenizer(bytes.NewBuffer([]byte(m.Content)))
		for {
			tt := z.Next()
			if tt == html.ErrorToken {
				break
			}

			switch tt {
			case html.TextToken:
				text := string(z.Text())
				if strings.TrimSpace(text) == "" {
					continue
				}
				fmt.Printf("\t%v\n", aurora.Blue(text))
			}
			if tt == html.ErrorToken {
				break
			}
		}
		fmt.Printf("\n")
	}
}
