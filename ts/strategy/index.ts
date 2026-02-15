import { BikeStrategy, CarStrategy, Commuter } from "./Commuter";

const commuter = new Commuter();
commuter.setStrategy(new CarStrategy());
commuter.commute();

commuter.setStrategy(new BikeStrategy());
commuter.commute();