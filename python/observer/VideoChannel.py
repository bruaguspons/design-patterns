from typing import List
from Subscriber import Subscriber

class VideoChannel:
    def __init__(self):
        self._subscribers: List[Subscriber] = []

    def subscribe(self, subscriber: Subscriber) -> None:
        self._subscribers.append(subscriber)

    def unsubscribe(self, subscriber: Subscriber) -> None:
        self._subscribers = [
            s for s in self._subscribers if s != subscriber
        ]

    def upload_video(self, video_title: str) -> None:
        print(f'Video "{video_title}" uploaded.')
        self._notify_subscribers(video_title)

    def _notify_subscribers(self, video_title: str) -> None:
        for subscriber in self._subscribers:
            subscriber.update(video_title)
