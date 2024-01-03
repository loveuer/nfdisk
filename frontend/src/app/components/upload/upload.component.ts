import {Component} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatButtonModule} from "@angular/material/button";
import {MatInputModule} from "@angular/material/input";
import {MatIconModule} from "@angular/material/icon";
import {FormsModule} from "@angular/forms";
import {MatTooltipModule} from "@angular/material/tooltip";
import {MatDialogRef} from "@angular/material/dialog";
import {ConnectionService} from "../../services/connection.service";
import {MsgService} from "../../services/msg.service";

@Component({
  selector: 'app-upload',
  standalone: true,
  imports: [CommonModule, MatButtonModule, MatInputModule, MatIconModule, FormsModule, MatTooltipModule],
  templateUrl: './upload.component.html',
  styleUrl: './upload.component.scss'
})
export class UploadComponent {
  constructor(
    public dialogRef: MatDialogRef<UploadComponent>,
    public srv: ConnectionService,
    private msg_srv: MsgService,
  ) {
  }

  name: string = ''
  filename: string = ''
  file: File | null = null
  ele: HTMLInputElement | null = null

  onFileSelected(event: Event) {
    console.log('[D] on file selected event=', event)
    this.ele = event.target as HTMLInputElement
    this.file = this.ele.files ? this.ele.files[0] : null
    this.filename = this.file?.name || ''
  }

  cleanSelected() {
    this.file = null
    this.filename = ''
    if (this.ele) {
      this.ele.files = null
    }
  }

  doUpload(check: boolean) {
    if (this.file && check) {
      let reader = new FileReader()
      reader.readAsDataURL(this.file)
      reader.onload = () => {
        if (typeof reader.result === "string") {
          this.srv.upload_file(this.name, reader.result).then(resp => {
            if (resp.status===200) {
              this.msg_srv.success(resp.msg)
            }
          })
        }
      }
    }
    this.dialogRef.close()
  }
}
