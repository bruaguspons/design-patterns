package main

import "fmt"

type VideoChannel struct {
	subscribers []Subscriber
}

func (vc *VideoChannel) Subscribe(sub Subscriber) {
	vc.subscribers = append(vc.subscribers, sub)
}

func (vc *VideoChannel) Unsubscribe(sub Subscriber) {
	for i, s := range vc.subscribers {
		if s == sub {
			vc.subscribers = append(vc.subscribers[:i], vc.subscribers[i+1:]...)
			break
		}
	}
}

func (vc *VideoChannel) UploadVideo(videoTitle string) {
	fmt.Printf("Video \"%s\" uploaded.\n", videoTitle)
	vc.notifySubscribers(videoTitle)
}

func (vc *VideoChannel) notifySubscribers(videoTitle string) {
	for _, sub := range vc.subscribers {
		sub.Update(videoTitle)
	}
}
