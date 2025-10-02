export namespace main {
	
	export class AWGConfig {
	    name: string;
	    filePath: string;
	
	    static createFrom(source: any = {}) {
	        return new AWGConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.filePath = source["filePath"];
	    }
	}
	export class RouterConfig {
	    url: string;
	    login: string;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new RouterConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.login = source["login"];
	        this.password = source["password"];
	    }
	}

}

