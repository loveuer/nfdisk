import {Component, Inject} from '@angular/core';
import {MatIconModule} from '@angular/material/icon';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatButtonModule} from "@angular/material/button";
import { FormsModule, ReactiveFormsModule} from "@angular/forms";
import {AddConnection} from '../../../../wailsjs/go/controller/App'
import {ConnectionService} from "../../services/connection.service";
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";

@Component({
  selector: 'app-add-connection',
  standalone: true,
  imports: [ReactiveFormsModule,MatIconModule, MatInputModule, MatFormFieldModule, MatButtonModule, FormsModule, ],
  templateUrl: './add-connection.component.html',
  styleUrl: './add-connection.component.scss'
})
export class AddConnectionComponent {
  constructor(
    public connection_srv: ConnectionService,
    public dialogRef: MatDialogRef<AddConnectionComponent>,
  ) {
  }

  newConnection = {
    name:  "",
    endpoint: "",
    access: "",
    secret: "",
  }

  createConnection(check: boolean) {
    if (check) {
      this.connection_srv.add_connection(
        this.newConnection.name,
        this.newConnection.endpoint,
        this.newConnection.access,
        this.newConnection.secret,
      )
    }
    this.dialogRef.close()
  }
}
