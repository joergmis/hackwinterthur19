import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowIssueComponent } from './show-issue.component';

describe('ShowIssueComponent', () => {
  let component: ShowIssueComponent;
  let fixture: ComponentFixture<ShowIssueComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ShowIssueComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ShowIssueComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
