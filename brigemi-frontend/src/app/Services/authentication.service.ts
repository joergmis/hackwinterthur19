import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';

import {RestService} from "./rest-service";
import { User } from '../Objects/user';

@Injectable({ providedIn: 'root' })
export class AuthenticationService {
  private user : User;
  private endpoint = 'http://localhost:8090/';
  private restService;

  constructor(private http: HttpClient) { 
    this.restService = new RestService(this.http);
  }

  login(username: string, password: string) {
    var main = this;
    this.user = new User(username, password);
    return this.http.post<any>(this.endpoint + `users/authenticate`,  { "name": username, "password": password } ).pipe(
      (map(user => {
        // login successful if there's a user in the response
        console.log(user);
        if (user && user != "login not successful") {
            // store user details and basic auth credentials in local storage 
            // to keep user logged in between page refreshes
            localStorage.setItem('currentUser', JSON.stringify(user));
        }
        return user;
    })));
  }

  logout() {
      // remove user from local storage to log user out
      localStorage.removeItem('currentUser');
  }
}