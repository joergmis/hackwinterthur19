import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';

import { AppComponent } from './app.component';
import { IssuesComponent } from './Pages/issues/issues.component';
import { CreateIssueComponent } from './Pages/create-issue/create-issue.component';

@NgModule({
  declarations: [
    AppComponent,
    IssuesComponent,
    CreateIssueComponent
  ],
  imports: [
    BrowserModule,
    NgbModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
