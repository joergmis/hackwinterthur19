export class Issue {
    private ID: number;
    private Name: string;
    private Description: string;
    private Userid: number;
    private Fileid?: number;
    private Documentid?: number;
    
    constructor(
        ID: number, 
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
    
    setIssueID(ID: number) {
        this.ID = ID;
    }
    
    public getId(): number {
        return this.ID;
    };     
}
