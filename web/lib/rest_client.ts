import axios, { AxiosRequestConfig } from "axios";
import { catchAxiosError } from "./services/error";

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
  return http.get<T>(url, config).catch(catchAxiosError);
}

export function post<T = any>(url: string, data: StringProps = {}, headers: StringProps = {}) {
  const body = new URLSearchParams(data);
  const config = {
    headers: new Headers(headers),
  };
  return http.post<T>(url, body, config).catch(catchAxiosError);
}

export function postMultipart(url: string, data: FormData = new FormData(), headers: StringProps = {}) {
  const config = {
    headers: new Headers({
      ...headers,
      "Content-Type": "multipart/form-data",
    }),
  } as AxiosRequestConfig;
  return http.post(url, data, config);
}
