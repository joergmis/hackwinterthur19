import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-show-issue',
  templateUrl: './show-issue.component.html',
  styleUrls: ['./show-issue.component.css']
})
export class ShowIssueComponent implements OnInit {

  constructor(private router : Router) { }

  ngOnInit() {
  }

  redirectFunction(url : string, id : number) : void {
    this.router.navigate([url],{ queryParams: { id: id } });
  }

}
