package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"flag"
	"github.com/wfinn/gozapread"
)


func main() {
	pass := flag.String("p", "bbbbbb", "password")
	flag.Parse()
	if gozapread.Login("downvote2donate", *pass) != nil {
		fmt.Println("Login failed")
		return
	}
	for {
		if gozapread.UnreadMessages() {
			if messages, err := gozapread.GetMessageTable(); err == nil {
				for _, message := range messages.Data {
					if strings.Contains(message.Message, "@downvote2donate") && message.Status == "Unread" {
						if u, err := strconv.ParseUint(message.Link, 10, 32); err == nil {
							if postresp, err := gozapread.SubmitNewPost("downvote2donate", `<h2>downvote2donate</h2>If you downvote this post <a href="https://github.com/Horndev/zapread.com#vote-examples">80% go to this group, 10% to the community and 10% to zapread</a>.<br>Mention me and I'll create this post in the group you mentioned me in.`, gozapread.GetGroupId(uint(u))); err == nil {
							if commentId, err := strconv.ParseUint(message.Anchor, 10, 32); err == nil {
								gozapread.AddComment(fmt.Sprintf(`<a href="https://www.zapread.com/Post/Detail/%d">Here</a>'s the post you can downvote to donate to this group.`, postresp.PostID), uint(u), uint(commentId))
							}}
						}
					}
					gozapread.DismissMessage(message.ID)
				}

			}
		}
		time.Sleep(15 * time.Minute)
	}
}
