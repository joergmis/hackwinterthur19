export class Issue {
    private id;
    private name;
    private password;
    
    constructor(name: string, password: string) {
        this.name = name;
        this.password = password;
    }
    
    public static fromJSON = (json: string): Issue => {
        const jsonObject = JSON.parse(json);
        return new Issue(
            jsonObject.name,
            jsonObject.password
        );
    };
    
    setIssueID(id: number) {
        this.id = id;
    }
    
    public getId(): string {
        return this.id;
    };     
}
