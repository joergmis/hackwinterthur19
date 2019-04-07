// Source: https://x-team.com/blog/webcam-image-capture-angular/

import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { RestService } from 'src/app/Services/rest-service';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';

import { Issue } from "src/app/Objects/issue";
import { Tag } from "src/app/Objects/tag";
import { modelGroupProvider } from '@angular/forms/src/directives/ng_model_group';
import { forEach } from '@angular/router/src/utils/collection';
import { jsonpCallbackContext } from '@angular/common/http/src/module';

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
  tags = '';
  tag = [];

  private httpClient;
  private restService;
  private imageToUpload;


  constructor(httpClient: HttpClient, private router : Router) {
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

    let canvasElement = this.canvas.nativeElement;
    canvasElement.width = width;
    canvasElement.height = height;
    canvasElement.getContext("2d").drawImage(this.video.nativeElement, 0, 0, width, height);

    this.imageToUpload = canvasElement.toDataURL("image/png");

    this.image.nativeElement.src = this.imageToUpload;
    this.image.nativeElement.width = width;
    this.image.nativeElement.height = height;
  }

  createIssue() {
    var main = this;
    console.log(main.tags);
    var allTags = main.tags.split(/[ ,]+/).filter(Boolean);
    // create all the tags
    for (var i = 0; i < allTags.length; i++) {
      var payload = {name: allTags[i]}
      this.restService.post(payload, "tags").subscribe(
        data => {
          var t = new Tag();
          t.Id = data.ID;
          t.Name = data.Name;
          console.log("created tag "+ JSON.stringify(t));
          this.tag.push(t);
        },
        err => {
          console.log(err);
        }
      )
    }
    // create the issue itself
    this.restService.post(this.model, "issues").subscribe(
      data => {
        // got the issue back with th id
        main.model = data;
        main.uploadImage()
      },
      err => {
        console.log("error occurred");
        let modelAsString = JSON.stringify(this.model);
        localStorage.setItem("issue_" + hashCode(modelAsString), modelAsString);
      }
    );
    
    for (let i =  0; i < localStorage.length; i++) {
      let itemName = localStorage.key(i);
      if (itemName.startsWith("issue")) {
        this.restService.post(JSON.parse(localStorage.getItem(itemName)), "issues").subscribe(
          data => { 
            console.log("POST done");
            localStorage.removeItem(itemName);
            console.log("\"" + itemName + "\" from localStorage removed");
          }
        );
      }
    }

  }

  uploadImage() {
    var main = this;
    var blob = '';
    this.canvas.nativeElement.toBlob(function (data) {
      blob = data;
      let formData = new FormData();
      let time = new Date();
      formData.append('image', blob, time.getMilliseconds() + '.png')
      let headers = new HttpHeaders({ "contentType": "multipart/form-data" });
      main.httpClient.post("http://localhost:8090/fileupload", formData, headers).subscribe(
        data => {
          console.log(data);
          main.model.Fileid = data.ID;
          main.restService.put(main.model, "issues").subscribe(
            data => {
              console.log(data);
              main.router.navigate(["/issues"]);
            },
            err => console.error("Erroro: " + err)
          );
        },
        err => console.error(err)
      );
    })
  }
}

function hashCode(data) {
  var hash = 0, i, chr;
  if (data.length === 0) return hash;
  for (i = 0; i < data.length; i++) {
    chr   = data.charCodeAt(i);
    hash  = ((hash << 31) - hash) + chr;
    hash |= 0;
  }
  return hash;
}
