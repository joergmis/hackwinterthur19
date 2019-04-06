// Source: https://x-team.com/blog/webcam-image-capture-angular/

import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';

@Component({
  selector: 'app-create-issue',
  templateUrl: './create-issue.component.html',
  styleUrls: ['./create-issue.component.css']
})
export class CreateIssueComponent implements OnInit {

  @ViewChild("video")
  public video: ElementRef

  @ViewChild("canvas")
  public canvas: ElementRef

  public constructor() {  }

  ngOnInit() {
  }

  public ngAfterViewInit() {
    if(navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
        navigator.mediaDevices.getUserMedia({ video: true }).then(stream => {
            this.video.nativeElement.srcObject = stream; // .src = window.URL.createObjectURL(stream);
            this.video.nativeElement.play();
        });
    }
  }

  public capture() {
    var context = this.canvas.nativeElement.getContext("2d").drawImage(this.video.nativeElement, 0, 0, 640, 480);
    // this.captures.push(this.canvas.nativeElement.toDataURL("image/png"));
  }

  public createIssue() {
    
  }
}
