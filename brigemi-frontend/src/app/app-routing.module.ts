import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { IssuesComponent } from './Pages/issues/issues.component';
import { CreateIssueComponent } from './Pages/create-issue/create-issue.component';
import { ShowIssueComponent } from './Pages/show-issue/show-issue.component';
import { LoginComponent } from './Pages/login/login.component';
import { RegistrationComponent } from './Pages/registration/registration.component';
import { AuthGuard } from './auth/auth.guard';

const routes: Routes = [
  { path: '', component: LoginComponent },
  { path: 'createIssues', component: CreateIssueComponent, canActivate: [AuthGuard] },
  { path: 'showIssue', component: ShowIssueComponent, canActivate: [AuthGuard] },
  { path: 'issues', component: IssuesComponent, canActivate: [AuthGuard] },
  { path: 'register', component: RegistrationComponent},
  { path: '**', redirectTo: '' }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes)
  ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }
