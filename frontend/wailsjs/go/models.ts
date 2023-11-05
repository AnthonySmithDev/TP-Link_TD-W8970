export namespace wireless {
	
	export class Info {
	    ssid: string;
	    enable: boolean;
	    password: string;
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ssid = source["ssid"];
	        this.enable = source["enable"];
	        this.password = source["password"];
	    }
	}

}

