import {Component, Inject} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import {MatIconModule} from "@angular/material/icon";
import {FormsModule} from "@angular/forms";
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {MatButtonModule} from "@angular/material/button";

@Component({
  selector: 'app-new-bucket',
  standalone: true,
  imports: [CommonModule, MatFormFieldModule, MatInputModule, MatIconModule, FormsModule, MatButtonModule],
  templateUrl: './new-bucket.component.html',
  styleUrl: './new-bucket.component.scss'
})
export class NewBucketComponent {
  constructor(
    public dialogRef: MatDialogRef<NewBucketComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { name: string, create: boolean }
  ) {
  }

  create(check: boolean) {
    this.data.create = check
    this.dialogRef.close()
  }
}
