package main

import (
	"fmt"
	"log"
	"os"

	mailinator "github.com/manybrain/mailinator-go-client"
)

var (
	domainName string
	inboxName  string
	apiToken   string
)

func main() {
	fmt.Printf("Go Mailinator\n")
	if len(os.Args) < 4 {
		log.Fatal("This test program requires 3 arguments: [apiToken] [domainName] [inboxName]")
	}
	apiToken = os.Args[1]
	domainName = os.Args[2]
	inboxName = os.Args[3]

	// Replace API_TOKEN with your real key
	client := mailinator.NewMailinatorClient(apiToken)

	//Get Team
	stats, err := client.GetTeamStats()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Team stats: %+v\n", stats)

	//Get TeamStats
	team, err := client.GetTeam()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("TeamName:%s\n", team.TeamName)
	for _, member := range team.Members {
		fmt.Printf("Member email:%s role:%s id:%s\n", member.Email, member.Role, member.Id)

	}

	//Fetch Inbox
	inbox, err := client.FetchInbox(&mailinator.FetchInboxOptions{Domain: domainName, Inbox: inboxName})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inbox:%+v\n", inbox)
	for _, message := range inbox.Messages {
		fmt.Printf("message id:%s subject:%s", message.Id, message.Subject)
		//Fetch Message Links
		ml, err := client.FetchMessageLinks(&mailinator.FetchMessageLinksOptions{domainName, inboxName, message.Id})
		if err != nil {
			log.Fatal(err)
		}
		for _, link := range ml.Links {
			fmt.Printf(" link:%s", link)
		}
		fmt.Printf("\n")
	}
}
