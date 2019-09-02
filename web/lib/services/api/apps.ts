import { get } from "../../rest_client";
import { API_ROUTES } from "../../routes";
import { App } from "./photos";

function ToApp(data: any) {
  return data as App;
}

export async function listApps(page: number, itemsPerPage: number) {
  const inputs = { page, itemsPerPage };
  const res: any = await get(API_ROUTES.apps.index.create(), inputs);
  if (res.error) {
    return res.error;
  }
  return res.data.records.map((photo: any) => ToApp(photo));
}
