import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { IssuesComponent } from './Pages/issues/issues.component';
import { CreateIssueComponent } from './Pages/create-issue/create-issue.component';
import { ShowIssueComponent } from './Pages/show-issue/show-issue.component';
import { LoginComponent } from './Pages/login/login.component';

const routes: Routes = [
  { path: '', component: IssuesComponent },
  { path: 'createIssues', component: CreateIssueComponent },
  { path: 'showIssue', component: ShowIssueComponent },
  { path: 'login', component: LoginComponent }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes)
  ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
