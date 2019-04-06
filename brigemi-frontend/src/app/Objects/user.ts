export class User {
    private Id;
    private Name;
    private Password;
    
    constructor(Name: string, Password: string) {
        this.Name = Name;
        this.Password = Password;
    }
    
    public static fromJSON = (json: string): User => {
        const jsonObject = JSON.parse(json);
        return new User(
            jsonObject.Name,
            jsonObject.Password
        );
    };
    
    public setUserID(Id: number) {
        this.Id = Id;
    }
    
    public getId(): string {
        return this.Id;
    };     
}
