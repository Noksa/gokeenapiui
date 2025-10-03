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
	export class RouteConfig {
	    interfaceId: string;
	    batFiles: string[];
	    batUrls: string[];
	
	    static createFrom(source: any = {}) {
	        return new RouteConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.interfaceId = source["interfaceId"];
	        this.batFiles = source["batFiles"];
	        this.batUrls = source["batUrls"];
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

