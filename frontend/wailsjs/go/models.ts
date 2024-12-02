export namespace main {
	
	export class VideoRequest {
	    page: number;
	    per_page: number;
	
	    static createFrom(source: any = {}) {
	        return new VideoRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.per_page = source["per_page"];
	    }
	}

}

