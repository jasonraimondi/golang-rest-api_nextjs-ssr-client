import axios, { AxiosInstance, AxiosRequestConfig } from "axios";

class AppRestClient {
  constructor(private readonly http: AxiosInstance) {
  }

  get<T = any>(url: string, config?: AxiosRequestConfig) {
    return this.http.get<T>(url, this.mergeConfig(config));
  }

  post<T = any>(url: string, data?: any, config?: AxiosRequestConfig) {
    const params = new URLSearchParams();
    for (const key in data) {
      params.append(key, data[key]);
    }
    return this.http.post<T>(url, params, config);
  }

  private mergeConfig(config?: AxiosRequestConfig): AxiosRequestConfig {
    if (!config) {
      config = {};
    }
    config = {
      ...config,
      headers: {
        ...config.headers,
      },
    };
    return config;
  }
}

const http = axios.create({
  baseURL: "http://localhost:1323",
});
export const appRestClient = new AppRestClient(http);
