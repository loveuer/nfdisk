// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';

export function AddConnection(arg1:string,arg2:string,arg3:string,arg4:string):Promise<model.RespConnectionList>;

export function CreateBucket(arg1:number,arg2:string):Promise<model.RespObjectList>;

export function DelConnection(arg1:number):Promise<model.RespConnectionList>;

export function DeleteBucket(arg1:number,arg2:string):Promise<model.RespObjectList>;

export function DeleteObject(arg1:number,arg2:string,arg3:string):Promise<model.RespMsg>;

export function DoConnect(arg1:number):Promise<model.RespObjectList>;

export function GetObject(arg1:number,arg2:string,arg3:string):Promise<model.RespMsg>;

export function HeadObject(arg1:number,arg2:string,arg3:string):Promise<model.RespObject>;

export function ListBucket(arg1:number):Promise<model.RespObjectList>;

export function ListConnection(arg1:number,arg2:number):Promise<model.RespConnectionList>;

export function ListObject(arg1:number,arg2:string,arg3:string,arg4:string):Promise<model.RespObjectList>;

export function ShareObject(arg1:number,arg2:string,arg3:string):Promise<model.RespShare>;

export function UploadObject(arg1:number,arg2:string,arg3:string,arg4:string,arg5:string):Promise<model.RespObjectList>;
