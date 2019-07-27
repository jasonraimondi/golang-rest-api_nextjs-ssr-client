import axios, { AxiosRequestConfig } from "axios";
import { catchAxiosError } from "./services/error";

export interface StringProps {
  [id: string]: string
}

export function get<T = any>(
  url: string,
  params: URLSearchParams = new URLSearchParams(),
  headers: StringProps = {},
) {
  return axios.get<T>(url, mergeConfig({ headers, params })).catch(catchAxiosError);
}

export function post<T = any>(
  url: string,
  data: URLSearchParams = new URLSearchParams(),
  headers: StringProps = {},
) {
  return axios.post<T>(url, data, mergeConfig({ headers })).catch(catchAxiosError);
}

export function postMultipart(url: string, data: FormData = new FormData(), headers: StringProps = {}) {
  return axios.post(url, data, mergeConfig({ headers: mergeHeaders(headers) })).catch(catchAxiosError);
}


function mergeHeaders(headers: StringProps) {
  return {
    ...headers,
    "Content-Type": "multipart/form-data",
  };
}

function mergeConfig(config: AxiosRequestConfig): AxiosRequestConfig {
  return {
    ...config,
    baseURL: "http://localhost:1323",
  };
}
