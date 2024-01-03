export interface Connection {
  id: number;
  name: string;
  endpoint: string;
  access_key: string;
  secret_key: string;
  active: boolean;
}
