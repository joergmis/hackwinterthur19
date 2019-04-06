import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import { Router } from '@angular/router';

import {RestService} from "../../Services/rest-service";

import { Issue } from '../../Objects/issue';

@Component({
  selector: 'app-issues',
  templateUrl: './issues.component.html',
  styleUrls: ['./issues.component.css']
})
export class IssuesComponent implements OnInit {
  public issues : Issue[] = [];

  private restService;

  constructor(private http: HttpClient, private router : Router) {
    this.restService = new RestService(this.http);
  }

  ngOnInit() {
    this.issues = this.restService.getAll("issues");
  }

  redirectFunction(url : string, id : number) : void {
    this.router.navigate([url],{ queryParams: { id: id } });
  }

}
