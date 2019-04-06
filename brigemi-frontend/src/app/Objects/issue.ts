export class Issue {
    public Id: number;
    public Name: string;
    public Description: string;
    Userid: number;
    Fileid?: number;
    Documentid?: number;
    
    constructor(
        Id: number, 
        Name: string,
        Description: string,
        Userid: number,
        Fileid?: number,
        Documentid?: number
        ) { 
            this.Name = Name;
            this.Description = Description;
            this.Userid = Userid;
            this.Fileid = Fileid;
            this.Documentid = Documentid;
        }
    
    setIssueID(Userid: number) {
        this.Userid = Userid;
    }
    
    public getId(): number {
        return this.Userid;
    };     
}
