import {AxiosInstance, AxiosRequestConfig} from 'axios';

export class AxiosRestClient {
    constructor(private readonly http: AxiosInstance) {
    }

    get<T = any>(url: string, config?: AxiosRequestConfig) {
        return this.http.get<T>(url, config);
    }

    post<T = any>(url: string, data: any, config?: AxiosRequestConfig) {
        return this.http.post<T>(url, data, config);
    }
}

export class AppRestClient {
    constructor(private readonly http: AxiosRestClient) {
    }

    get<T = any>(url: string, config?: AxiosRequestConfig) {
        return this.http.get<T>(url, this.mergeConfig(config))
    }

    post<T = any>(url: string, data: any, config?: AxiosRequestConfig) {
        return this.http.post<T>(url, data, this.mergeConfig(config));
    }

    private mergeConfig(config?: AxiosRequestConfig): AxiosRequestConfig {
        if (config) return {...config};
        return {}
    }
}

