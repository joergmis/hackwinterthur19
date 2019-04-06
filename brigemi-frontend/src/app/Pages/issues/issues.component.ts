import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {RestService} from "../../Services/rest-service";

@Component({
  selector: 'app-issues',
  templateUrl: './issues.component.html',
  styleUrls: ['./issues.component.css']
})
export class IssuesComponent implements OnInit {
  public issues;

  private restService;

  constructor(private http: HttpClient) {
    this.restService = new RestService(this.http);
  }

  ngOnInit() {
    this.issues = this.restService.getAll("issues");
  }

}
