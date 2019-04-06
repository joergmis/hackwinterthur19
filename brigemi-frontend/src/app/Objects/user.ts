export class User {
    private id;
    private name;
    private password;
    
    constructor(name: string, password: string) {
        this.name = name;
        this.password = password;
    }
    
    public static fromJSON = (json: string): User => {
        const jsonObject = JSON.parse(json);
        return new User(
            jsonObject.name,
            jsonObject.password
        );
    };
    
    public setUserID(id: number) {
        this.id = id;
    }
    
    public getId(): string {
        return this.id;
    };     
}
