// Source: https://x-team.com/blog/webcam-image-capture-angular/

import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { RestService } from 'src/app/Services/rest-service';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Issue } from "src/app/Objects/issue";

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

  @ViewChild("image")
  public image: ElementRef

  @ViewChild("file")
  public file: ElementRef

  isCameraShown = false;

  isPhotoTaken = false;

  // An issue id of 0 indicates a new issue to be created
  model = new Issue(0, "", "", 1);

  private httpClient;
  private restService;
  private imageToUpload;

  constructor(httpClient: HttpClient) {
    this.httpClient = httpClient;
    this.restService = new RestService(httpClient);
  }

  ngOnInit() { }

  ngAfterViewInit() {
    if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
      navigator.mediaDevices.getUserMedia({ video: true }).then(stream => {
        this.video.nativeElement.srcObject = stream;
        this.video.nativeElement.play();
      });
    }
  }

  capture() {
    let width = this.video.nativeElement.offsetWidth;
    let height = this.video.nativeElement.offsetHeight;

    this.canvas.nativeElement.getContext("2d").drawImage(this.video.nativeElement, 0, 0, width, height);
    this.imageToUpload = this.canvas.nativeElement.toDataURL("image/png");

    this.image.nativeElement.src = this.imageToUpload;
    this.image.nativeElement.width = width;
    this.image.nativeElement.height = height;
  }

  createIssue() {
    this.uploadImage();
    this.restService.post(this.model, "issues").subscribe(
      data => {
        console.log("POST done");
      },
      err => console.error("Erroro: " + err)
    );
  }

  uploadImage() {
    var main = this;
    var blob = '';
    this.canvas.nativeElement.toBlob(function (data) {
      blob = data;
      let formData = new FormData();
      let time = new Date(100);
      formData.append('image', blob, time.getMilliseconds() + '.png')
      let headers = new HttpHeaders({ "contentType": "multipart/form-data" });
      main.httpClient.post("http://localhost:8090/fileupload", formData, headers).subscribe(
        data => { console.log(data) },
        err => console.error(err)
      );
    })
  }
}
