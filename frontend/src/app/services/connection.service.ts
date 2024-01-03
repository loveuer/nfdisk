import {Injectable} from '@angular/core';
import {
  AddConnection,
  CreateBucket,
  DelConnection,
  DeleteBucket,
  DoConnect,
  GetObject,
  HeadObject,
  ListBucket,
  ListConnection,
  ListObject,
  UploadObject,
  DeleteObject,
  ShareObject,
} from "../../../wailsjs/go/controller/App"
import {Connection} from "../interface/connection";
import {Item} from "../interface/item";
import {model} from "../../../wailsjs/go/models";
import {MsgService} from "./msg.service";
import {fromPromise} from "rxjs/internal/observable/innerFrom";
import {tap} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class ConnectionService {

  current_conn: Connection = {} as Connection
  current_bucket: string = ''
  current_start: string = ''
  current_paths: string[] = []
  current_filter: string = ''
  list: Connection[] = []
  items: Item[] = []

  constructor(
    private msg_srv: MsgService,
  ) {
    this.get_connections()
  }

  private init_current() {
    this.current_bucket = ""
    this.current_start = ""
    this.current_paths = []
  }

  private list_object(filter = "") {
    let path = this.current_paths.join('')

    if (filter) {
      if (path) {
        path = path.endsWith('/') ? path + filter : path + '/' + filter
      } else {
        path = filter
      }
    }


    console.log(`[D] list_object: id=${this.current_conn.id} bucket=${this.current_bucket} path=${path} start=${this.current_start}`)

    ListObject(this.current_conn.id, this.current_bucket, path, this.current_start).then(resp => {
      if (resp.status === 200) {
        this.items = resp.data.map(v => {
          let data = v as Item
          data.name = filter+data.name
          return data
        })
        console.log('[D] list_object: resp=', resp)
      } else {
        console.log('[E] list_object err=', resp)
      }
    }).catch(e => {
      console.error('[E] list_object err=', e)
    })
  }

  private list_bucket() {
    ListBucket(this.current_conn.id).then(resp => {
      if (resp.status === 200) {
        this.items = resp.data.map(v => v as Item)
      } else {
        console.log('[E] list_bucket err=', resp)
      }
    }).catch(err => console.error('[E] list_bucket err=', err))
  }

  back_list() {
    if (this.current_paths.length > 0) {
      this.current_paths.pop()
      this.list_object()
    } else {
      this.current_bucket = ''
      this.list_bucket()
    }
  }


  get_list(item: Item) {
    switch (item.type) {
      case 'bucket':
        this.current_filter = ''
        this.current_bucket = item.name
        this.current_paths = []
        break
      case 'folder':
        this.current_filter = ''
        this.current_paths.push(item.name)
        break
      case 'file':
        break
    }

    this.list_object()
  }

  filter_list(prefix: string) {
    this.list_object(prefix)
  }

  do_connect(item: Connection) {
    DoConnect(item.id).then(resp => {
      this.init_current()
      this.current_conn = item
      if (resp.status === 200) {
        this.items = resp.data.map(v => v as Item)
      } else {
        console.log('[E] do_connect resp=', resp)
      }
    }).catch(e => {
      console.error('[E] do_connect err=', e)
    })
  }

  get_connections() {
    console.log('[D] get_connections called...')
    ListConnection(0, -1).then(resp => {
      if (resp.status === 200) {
        this.list = resp.data
      } else {
        console.log('[E] get_connections: resp=', resp)
      }
    }).catch(e => {
        console.error('[E] get_connections err=', e)
      }
    )
  }

  add_connection(name: string, endpoint: string, access: string, secret: string) {
    AddConnection(name, endpoint, access, secret).then(resp => {
      if (resp.status === 200) {
        this.list = resp.data
      } else {
        console.log('[E] add_connection err=', resp)
      }
    }).catch(err => {
      console.error('[E] add_connection err=', err)
    })
  }

  del_connection(id: number) {
    DelConnection(id).then(resp => {
      if (resp.status === 200) {
        this.list = resp.data
      } else {
        console.log('[E] del_connection err=', resp)
      }
    }).catch(err => {
      console.error('[E] del_connection err=', err)
    })
  }

  del_bucket(name: string) {
    DeleteBucket(this.current_conn.id, name).then(resp => {
      if (resp.status === 200) {
        this.init_current()
        this.items = resp.data.map(v => v as Item)
        this.msg_srv.success(resp.msg)
      } else {
        this.msg_srv.error(resp.msg)
      }
    })
  }

  create_bucket(name: string) {
    CreateBucket(this.current_conn.id, name).then(resp => {
      if (resp.status === 200) {
        this.init_current()
        this.items = resp.data.map(v => v as Item)
        this.msg_srv.success(`新建桶 ${name} 成功`)
      } else {
        console.log('[E] create_bucket err=', resp)
      }
    }).catch(err => console.error('[E] create_bucket err=', err))
  }

  upload_file(key: string, content: string) {
    return UploadObject(
      this.current_conn.id,
      this.current_bucket,
      this.current_paths.join(''),
      key,
      content
    );
  }

  async get_file(name: string): Promise<model.RespMsg> {
    const path = this.current_paths.join('') + name
    return GetObject(this.current_conn.id, this.current_bucket, path)
  }

  head_file(name:string) {
    const path = this.current_paths.join('') + name
    const pr = HeadObject(this.current_conn.id, this.current_bucket, path)
    return fromPromise(pr)
  }

  delete_file(name: string) {
    let paths = [...this.current_paths]
    paths.push(name)
    DeleteObject(this.current_conn.id, this.current_bucket, paths.join('')).then(resp => {
      console.log('[D] delete_file resp=', resp)
      if (resp.status === 200) {
        this.msg_srv.success(resp.msg)
      }
    })
  }

  async share_file(item: Item): Promise<string> {
    const path = this.current_paths.join('') + item.name
    const resp = await ShareObject(this.current_conn.id, this.current_bucket, path)
    if (resp.status === 200) {
      return resp.data
    }

    return ''
  }
}
