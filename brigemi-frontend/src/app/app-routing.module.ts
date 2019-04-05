import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { IssuesComponent } from './issues/issues.component';
import { CreateIssueComponent } from './create-issue/create-issue.component';

const routes: Routes = [
  { path: '', component: IssuesComponent },
  { path: 'createIssues', component: CreateIssueComponent }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes)
  ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
