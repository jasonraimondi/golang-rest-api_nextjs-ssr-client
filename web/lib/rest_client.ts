import axios from "axios";

const http = axios.create({
  baseURL: "http://localhost:1323",
});

export interface StringProps {
  [id: string]: string
}

export function get<T = any>(url: string, params?: StringProps, headers?: StringProps) {
  if (!params) params = {};
  if (!headers) headers = {};
  return http.get<T>(url, {
    params: new URLSearchParams(params),
    headers: new Headers(headers),
  });
}

export function post<T = any>(url: string, data?: StringProps, headers?: StringProps) {
  if (!data) data = {};
  if (!headers) headers = {};
  return http.post<T>(url, new URLSearchParams(data), {
    headers: new Headers(headers),
  });
}

export function postMultipart(url: string, data?: FormData, headers?: StringProps) {
  if (!headers) headers = {};
  return http.post(url, data, {
    headers: new Headers(headers),
  });
}
