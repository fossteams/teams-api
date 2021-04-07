package teams_api_test

import (
	"fmt"
	teams_api "github.com/fossteams/teams-api"
	"github.com/fossteams/teams-api/pkg/csa"
	"github.com/logrusorgru/aurora"
	"sort"
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
