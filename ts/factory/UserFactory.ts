export interface User {
    name: string;
    age: number;
}

export class AdminUser implements User {
    constructor(public name: string, public age: number) {}
}

export class GuestUser implements User {
    constructor(public name: string, public age: number) {}
}

export class RegularUser implements User {
    constructor(public name: string, public age: number) {}
}

export class UserFactory{
    public static createUser(type: "admin" | "guest" | "regular", payload: User): User {
        switch (type) {
            case "admin":
                return new AdminUser(payload.name, payload.age);
            case "guest":
                return new GuestUser(payload.name, payload.age);
            case "regular":
                return new RegularUser(payload.name, payload.age);
            default:
                throw new Error("Invalid user type");
        }
    }
}