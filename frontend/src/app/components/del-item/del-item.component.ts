import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {Item} from "../../interface/item";
import {MatButtonModule} from "@angular/material/button";
import {NgIf} from "@angular/common";
import {ConnectionService} from "../../services/connection.service";

@Component({
  selector: 'app-del-item',
  standalone: true,
  imports: [
    MatButtonModule,
    NgIf
  ],
  templateUrl: './del-item.component.html',
  styleUrl: './del-item.component.scss'
})
export class DelItemComponent {
  constructor(
    public dialogRef: MatDialogRef<DelItemComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { item: Item, check: boolean },
    public srv: ConnectionService,
  ) {
  }

  del_item(del: boolean) {
    this.data.check = del
    this.dialogRef.close()
    switch (this.data.item.type) {
      case "bucket":
        this.srv.del_bucket(this.data.item.name)
        break
      case "folder":
      case "file":
        this.srv.delete_file(this.data.item.name)
    }
  }
}
