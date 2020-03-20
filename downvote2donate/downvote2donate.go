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
						if u, err := strconv.ParseUint(message.Link, 10, 32); err == nil {
							if postresp, err := api.SubmitNewPost("downvote2donate", `<h2>downvote2donate</h2>If you downvote this post <a href="https://github.com/Horndev/zapread.com#vote-examples">80% go to this group, 10% to the community and 10% to zapread</a>.<br>Mention me and I'll create this post in the group you mentioned me in.`, api.GetGroupId(uint(u))); err == nil {
								if commentId, err := strconv.ParseUint(message.Anchor, 10, 32); err == nil {
									api.AddComment(fmt.Sprintf(`<a href="https://www.zapread.com/Post/Detail/%d">Here</a>'s the post you can downvote to donate to this group.`, postresp.PostID), uint(u), uint(commentId))
								}
							}
						}
					}
					api.DismissMessage(message.ID)
				}

			}
		}
		time.Sleep(15 * time.Minute)
	}
}
