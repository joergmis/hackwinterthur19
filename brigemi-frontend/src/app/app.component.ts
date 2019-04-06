import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'brigemi-frontend';

  public currentPath : Location;

  constructor() {
    this.currentPath = window.location;
    console.log(this.currentPath.pathname);
  }
}
