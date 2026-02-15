import { TransportStrategy } from "./types";

export class CarStrategy implements TransportStrategy {
    transport(): void {
        console.log("Transporting by car");
    }
}

export class BikeStrategy implements TransportStrategy {
    transport(): void {
        console.log("Transporting by bike");
    }
}

export class BusStrategy implements TransportStrategy {
    transport(): void {
        console.log("Transporting by bus");
    }
}

export class Commuter {
    private strategy: TransportStrategy | null = null;

    setStrategy(strategy: TransportStrategy): void {
        this.strategy = strategy;
    }

    commute(): void {
        if (!this.strategy) {
            console.log("No transport strategy set");
            return;
        }
        this.strategy.transport();
    }
}