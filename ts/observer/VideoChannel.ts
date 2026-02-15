import { Suscriber } from "./types.d";

export class VideoChannel {
    private suscribers: Suscriber[] = [];

    subscrive(suscriber: Suscriber): void {
        this.suscribers.push(suscriber);
    }

    unsubscribe(suscriber: Suscriber): void {
        this.suscribers = this.suscribers.filter(s => s !== suscriber);
    }

    uploadVideo(videoTitle: string): void {
        console.log(`Video "${videoTitle}" uploaded.`);
        this.notifySubscribers(videoTitle);
    }

    private notifySubscribers(videoTitle: string): void {
        this.suscribers.forEach(suscriber => suscriber.update(videoTitle));
    }
}