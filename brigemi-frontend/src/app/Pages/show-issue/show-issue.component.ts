import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { RestService } from 'src/app/Services/rest-service';

import { Issue } from '../../Objects/issue';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-show-issue',
  templateUrl: './show-issue.component.html',
  styleUrls: ['./show-issue.component.css']
})
export class ShowIssueComponent implements OnInit {
  public issue : Issue;

  private restService;

  constructor(private http: HttpClient, private router : Router, private ar : ActivatedRoute) {
    var main = this;

    this.issue = new Issue(0,'','',0,0,0);
    this.restService = new RestService(this.http);

    ar.queryParams.subscribe(params => {
      this.issue.setIssueID(params['id']);
    })

    console.log(this.issue.getId());

    this.restService.get(this.issue.getId(), "issues/").subscribe(
      data => { main.issue = data; },
      err => { console.log(err); }
    );

  }

  ngOnInit() {
    
  }

  redirectFunction(url : string, id : number) : void {
    this.router.navigate([url],{ queryParams: { id: id } });
  }

}
