import {
  AfterRenderRef,
  AfterViewChecked,
  AfterViewInit,
  Component,
  ElementRef,
  Input,
  OnInit,
  ViewChild
} from '@angular/core';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatDialog} from "@angular/material/dialog";
import {AddConnectionComponent} from "../../components/add-connection/add-connection.component";
import {ConnectionService} from "../../services/connection.service";
import {CommonModule} from "@angular/common";
import {Item} from "../../interface/item";
import {MatTooltipModule} from "@angular/material/tooltip";
import {Connection} from "../../interface/connection";
import {DelConnectionComponent} from "../../components/del-connection/del-connection.component";
import {NewBucketComponent} from "../../components/new-bucket/new-bucket.component";
import {UploadComponent} from "../../components/upload/upload.component";
import {CdkContextMenuTrigger, CdkMenu, CdkMenuItem} from "@angular/cdk/menu";
import {DelItemComponent} from "../../components/del-item/del-item.component";
import {NewFolderComponent} from "../../components/new-folder/new-folder.component";
import {MsgService} from "../../services/msg.service";
import {MatInputModule} from "@angular/material/input";
import {FormsModule} from "@angular/forms";
import {debounce, debounceTime, fromEvent, map, timer} from "rxjs";
import {PreviewImageComponent} from "../../components/preview-image/preview-image.component";

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    MatIconModule,
    MatButtonModule,
    CommonModule,
    MatTooltipModule,
    CdkContextMenuTrigger,
    CdkMenu,
    CdkMenuItem,
    MatInputModule,
    FormsModule,
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent implements AfterViewInit{

  constructor(
    public dialog: MatDialog,
    public connection_srv: ConnectionService,
    private msg_srv: MsgService,
  ) {
  }

  @ViewChild('downloadRef') downloadLink!: ElementRef<HTMLLinkElement>;
  @ViewChild('pathFilterRef') pathFilterInput!: ElementRef<HTMLInputElement>;

  ngAfterViewInit() {
    const input = fromEvent(this.pathFilterInput.nativeElement, "input")
    const result = input.pipe(
      debounceTime(1000),
    )
    result.subscribe(_ => {
      this.connection_srv.filter_list(this.connection_srv.current_filter)
    })
  }


  openAddConnection() {
    const addConnectionDialog = this.dialog.open(AddConnectionComponent)
  }

  connect(item: Connection) {
    this.connection_srv.do_connect(item)
  }

  open(item: Item) {
    switch (item.type) {
      case 'bucket':
        this.connection_srv.get_list(item)
        break
      case 'folder':
        this.connection_srv.get_list(item)
        break
      case 'file':
        this.preview_item(item)
        break
    }
  }

  back() {
    this.connection_srv.back_list()
  }

  del_connection(conn: Connection) {
    let data = {conn: conn, del: false}
    const delConnectionDialog = this.dialog.open(DelConnectionComponent, {data: data}).afterClosed().subscribe(val => {
      if (data.del) {
        this.connection_srv.del_connection(conn.id)
      }
    })
  }

  create_bucket() {
    let data = {name: '', create: false}
    const newBucketDialog = this.dialog.open(NewBucketComponent, {data: data}).afterClosed().subscribe(val => {
      if (data.create) {
        this.connection_srv.create_bucket(data.name)
      }
    })
  }

  create_folder() {
    let data = {name: '', create: false}
    const newFolderDialog = this.dialog.open(NewFolderComponent, {data: data}).afterClosed().subscribe(val => {
      if (data.create) {
        this.connection_srv.create_bucket(data.name)
      }
    })
  }

  upload_item() {
    const uploadFileDialog = this.dialog.open(UploadComponent).afterClosed().subscribe(_ => {
    })
  }

  del_item(item: Item) {
    let data = {item: item, check: false}
    const delItemDialog = this.dialog.open(DelItemComponent, {data: data}).afterClosed().subscribe(val => {
      if (data.check) {
      }
    })
  }

  download_item(item: Item) {
    this.connection_srv.get_file(item.name).then(resp => {
      this.msg_srv.success('下载成功')
    })
  }

  // todo: get item header(content-type)
  head_item(item: Item):Item {
    return item
  }

  preview_item(item: Item) {
    this.connection_srv.head_file(item.name).subscribe(data => {
      if (data.status === 200) {
        if (data.data?.content_type.startsWith('image/')) {
          this.dialog.open(PreviewImageComponent, {data: item})
        }
      } else {
        this.msg_srv.warning(`获取文件类型失败`)
      }
    })
  }
}
