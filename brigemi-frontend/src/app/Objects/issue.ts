export class Issue {
    constructor(
        public id: number, 
        public name: string,
        public description: string,
        public userid: number,
        public fileid?: number,
        public documentid?: number
        ) { }  
    
    setIssueID(id: number) {
        this.id = id;
    }
    
    public getId(): number {
        return this.id;
    };     
}
