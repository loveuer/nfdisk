import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {Item} from "../../interface/item";
import {ConnectionService} from "../../services/connection.service";
import {MsgService} from "../../services/msg.service";

@Component({
  selector: 'app-preview-image',
  standalone: true,
  imports: [],
  templateUrl: './preview-image.component.html',
  styleUrl: './preview-image.component.scss'
})
export class PreviewImageComponent {
  constructor(
    public dialogRef: MatDialogRef<PreviewImageComponent>,
    @Inject(MAT_DIALOG_DATA) public data: Item,
    public srv: ConnectionService,
    private msg: MsgService,
  ) {
    this.srv.share_file(this.data).then(url => {
      if (url) {
        this.src = url
      }
    })
  }

  src = ''

  previewFail(evt: Event) {
    console.log('PreviewFail: called with event=', evt)
    this.msg.error("预览失败")
  }

  close() {
    this.dialogRef.close()
  }
}
