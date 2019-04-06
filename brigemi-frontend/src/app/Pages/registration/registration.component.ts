import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';

import { User } from '../../Objects/user';
import {RestService} from "../../Services/rest-service";


@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {

  private restService;
  private user : User;

  constructor(private http: HttpClient, private router : Router) { 
    this.restService = new RestService(this.http);
  }

  ngOnInit() {
  }

  createUser(form: NgForm) {
    var main = this;
    if(form.valid){
      console.log(form.value.txtUsername);
      main.user = new User(form.value.txtUsername, form.value.pwPassword);
      this.restService.post(main.user, "users").subscribe(
        data => {
          console.log(data);
          main.user.setUserID(data.id);
          this.router.navigate(['login']);
        },
        err => console.error('Observer got an error: ' + err)
      );
    }
  }

}
