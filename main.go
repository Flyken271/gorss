package main

import (
	"flag"

	"github.com/marcusolsson/tui-go"
	"github.com/marcusolsson/tui-go/wordwrap"
	"github.com/mmcdole/gofeed"
)

type feedItem struct {
	idx         int
	Title       string
	Content     string
	Description string
	Published   string
}

func main() {

	var feedFlag = flag.String("url", "example.com/example.rss", "URL to rss feed.")
	flag.Parse()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(*feedFlag)

	var feedlist []feedItem

	for idx, val := range feed.Items {
		feedlist = append(feedlist, feedItem{
			idx:         idx,
			Content:     val.Content,
			Description: val.Description,
			Published:   val.Published,
			Title:       val.Title,
		})
	}

	//fmt.Println(feed.Title)
	feedList := tui.NewList()
	feedTitle := tui.NewLabel(feed.Title)
	feedTitleBox := tui.NewVBox(feedTitle)
	feedTitleBox.SetSizePolicy(tui.Expanding, tui.Maximum)
	feedTitleBox.SetBorder(true)

	feedBox := tui.NewVBox(feedList)
	feedBox.SetBorder(true)

	for _, item := range feedlist {
		feedList.AddItems(item.Title)
	}

	feedList.SetFocused(true)

	itemDesc := tui.NewLabel("")
	itemBox := tui.NewVBox(itemDesc)
	itemBox.SetBorder(true)

	feedList.OnItemActivated(func(l *tui.List) {
		for idx, val := range feedlist {
			if l.Selected() == idx {
				itemDesc.SetText(wordwrap.WrapString(val.Description, 60))
			}
		}
	})

	descContainer := tui.NewHBox(feedBox, itemBox)

	root := tui.NewVBox(feedTitleBox, descContainer)

	ui, err := tui.New(root)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}

}
