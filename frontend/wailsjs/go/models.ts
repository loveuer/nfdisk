export namespace model {
	
	export class NFObject {
	    name: string;
	    key: string;
	    last_modified: number;
	    size: number;
	    type: string;
	    content_type: string;
	
	    static createFrom(source: any = {}) {
	        return new NFObject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.key = source["key"];
	        this.last_modified = source["last_modified"];
	        this.size = source["size"];
	        this.type = source["type"];
	        this.content_type = source["content_type"];
	    }
	}
	export class Connection {
	    id: number;
	    name: string;
	    endpoint: string;
	    access_key: string;
	    secret_key: string;
	    active: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.endpoint = source["endpoint"];
	        this.access_key = source["access_key"];
	        this.secret_key = source["secret_key"];
	        this.active = source["active"];
	    }
	}
	export class RespConnectionList {
	    status: number;
	    msg: string;
	    err: string;
	    data: Connection[];
	
	    static createFrom(source: any = {}) {
	        return new RespConnectionList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.msg = source["msg"];
	        this.err = source["err"];
	        this.data = this.convertValues(source["data"], Connection);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RespMsg {
	    status: number;
	    msg: string;
	    err: string;
	
	    static createFrom(source: any = {}) {
	        return new RespMsg(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.msg = source["msg"];
	        this.err = source["err"];
	    }
	}
	export class RespObject {
	    status: number;
	    msg: string;
	    err: string;
	    data?: NFObject;
	
	    static createFrom(source: any = {}) {
	        return new RespObject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.msg = source["msg"];
	        this.err = source["err"];
	        this.data = this.convertValues(source["data"], NFObject);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RespObjectList {
	    status: number;
	    msg: string;
	    err: string;
	    data: NFObject[];
	
	    static createFrom(source: any = {}) {
	        return new RespObjectList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.msg = source["msg"];
	        this.err = source["err"];
	        this.data = this.convertValues(source["data"], NFObject);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RespShare {
	    status: number;
	    msg: string;
	    err: string;
	    data: string;
	
	    static createFrom(source: any = {}) {
	        return new RespShare(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.msg = source["msg"];
	        this.err = source["err"];
	        this.data = source["data"];
	    }
	}

}

