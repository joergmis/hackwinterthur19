import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { IssuesComponent } from './issues/issues.component';
import { CreateIssueComponent } from './create-issue/create-issue.component';

@NgModule({
  declarations: [
    AppComponent,
    IssuesComponent,
    CreateIssueComponent
  ],
  imports: [
    BrowserModule,
    NgbModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
