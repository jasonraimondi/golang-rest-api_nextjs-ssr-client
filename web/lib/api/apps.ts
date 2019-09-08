import { App } from "../entity/app";
import { get } from "../rest_client";
import { API_ROUTES } from "../routes";
import { ApiResponse } from "./api_response";

function ToApp(data: any) {
  return data as App;
}


export async function listApps(page: number, itemsPerPage: number): Promise<ApiResponse<App[]>> {
  const inputs = { page, itemsPerPage };
  const res: any = await get(API_ROUTES.apps.index.create(), inputs);

  if (res.data && res.data.records) {
    return [
      res.data.records.map((photo: any) => ToApp(photo)),
      undefined,
    ];
  }

  if (res.error) {
    console.log("res error here");
    return [
      [],
      res.error,
    ];
  }

  return [
    [],
    "something went wrong!",
  ];
}
