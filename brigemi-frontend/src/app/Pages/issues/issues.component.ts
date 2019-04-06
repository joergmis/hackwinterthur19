import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Router } from '@angular/router';
import { Observable, of } from 'rxjs';

import {RestService} from "../../Services/rest-service";

import { Issue } from '../../Objects/issue';

@Component({
  selector: 'app-issues',
  templateUrl: './issues.component.html',
  styleUrls: ['./issues.component.css']
})
export class IssuesComponent implements OnInit {
  public issues;
  public documents = null;

  private restService;

  model = {
    name: ''
  };

  constructor(private http: HttpClient, private router : Router) {
    this.restService = new RestService(this.http);
  }

  ngOnInit() {
    this.issues = this.restService.getAll("issues");
  }

  searchcall() {
    this.restService.post(null, "search?tag=" + this.model.name).subscribe(
      data => {
        this.documents = of(data);
      },
      err => console.log(err)
    );
  }

  redirectFunction(url : string, id : number) : void {
    this.router.navigate([url],{ queryParams: { id: id } });
  }
}