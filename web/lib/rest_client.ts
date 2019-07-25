import axios, { AxiosRequestConfig } from "axios";

const http = axios.create({
  baseURL: "http://localhost:1323",
});

export interface StringProps {
  [id: string]: string
}

export function get<T = any>(url: string, params: StringProps = {}, headers: StringProps = {}) {
  const config = {
    params: new URLSearchParams(params),
    headers: new Headers(headers),
  };
  return http.get<T>(url, config);
}

export function post<T = any>(url: string, data: StringProps = {}, headers: StringProps = {}) {
  return http.post<T>(url, new URLSearchParams(data), {
    headers: new Headers(headers),
  });
}

export function postMultipart(url: string, data: FormData = new FormData(), headers: StringProps = {}) {
  headers = {
    ...headers,
    "Content-Type": "multipart/form-data",
  };
  const config = {
    headers: new Headers(headers),
  } as AxiosRequestConfig;
  return http.post(url, data, config);
}
