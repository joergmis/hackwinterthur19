import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

export class RestService{
  private endpoint = 'http://localhost:8090/';
  private httpOptions = {
    headers: new HttpHeaders({
      'Content-Type':  'application/json',
      'dataType': 'json'
    })
  };

  constructor(private http: HttpClient) { }

  post(object : any, url : string): Observable<any> {
    return this.http.post<any>(this.endpoint + url, JSON.stringify(object), this.httpOptions).pipe();
  }

  get(id : number, url : string): Observable<any> {
     return this.http.get<any>(this.endpoint + url + id, this.httpOptions).pipe();
  }

  getAll(url: string): Observable<any> {
    return this.http.get<any>(this.endpoint + url, this.httpOptions);
  }
}
