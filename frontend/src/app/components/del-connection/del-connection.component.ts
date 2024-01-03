import {Component, Inject} from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatButtonModule} from "@angular/material/button";
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {model} from "../../../../wailsjs/go/models";
import Connection = model.Connection;

@Component({
  selector: 'app-del-connection',
  standalone: true,
  imports: [CommonModule, MatButtonModule],
  templateUrl: './del-connection.component.html',
  styleUrl: './del-connection.component.scss'
})
export class DelConnectionComponent {
  constructor(
    public dialogRef: MatDialogRef<DelConnectionComponent>,
    @Inject(MAT_DIALOG_DATA) public data: {conn:Connection, del: boolean}
  ) {
  }

  del_confirm(check: boolean) {
    this.data.del = check
    this.dialogRef.close()
  }
}
