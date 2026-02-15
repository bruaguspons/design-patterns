import { UserFactory } from "./UserFactory";

const admin = UserFactory.createUser("admin", { name: "John", age: 30 });
const guest = UserFactory.createUser("guest", { name: "Jane", age: 25 });
const regular = UserFactory.createUser("regular", { name: "Bob", age: 40 });