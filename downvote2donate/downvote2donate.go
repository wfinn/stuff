package main

import (
	"flag"
	"fmt"
	"github.com/wfinn/gozapread"
	"strconv"
	"strings"
	"time"
)

func main() {
	pass := flag.String("p", "bbbbbb", "password")
	flag.Parse()
	api, err := gozapread.Login("downvote2donate", *pass)
	if err != nil {
		fmt.Println("Login failed")
		return
	}
	for {
		if api.UnreadMessages() {
			if messages, err := api.GetMessageTable(); err == nil {
				for _, message := range messages.Data {
					if strings.Contains(message.Message, "@downvote2donate") && message.Status == "Unread" {
						if postid, err := strconv.ParseUint(message.Link, 10, 32); err == nil {
							if commentId, err := strconv.ParseUint(message.Anchor, 10, 32); err == nil {
								api.AddComment(fmt.Sprintf(`Hi, %s!<br>You can <b>downvote this comment to donate to this group!</b><br><a href="https://github.com/Horndev/zapread.com#vote-examples">80%% of the donations go to this group, 10%% to the community and 10%% to zapread.</a>`, message.From), uint(postid), uint(commentId))
							}
						}
					}
					api.DismissMessage(message.ID)
				}
			}
		}
		time.Sleep(20 * time.Minute)
	}
}
