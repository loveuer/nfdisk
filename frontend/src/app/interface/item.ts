export interface Item {
  name: string
  last_modified: number
  type: 'bucket' | 'folder' | 'file'
  content_type: string | null
}
