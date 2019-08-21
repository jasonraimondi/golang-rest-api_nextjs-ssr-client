import Cookie from "js-cookie";
import Router from "next/router";
import { LoginInputs } from "../../../pages/login";
import { COOKIES } from "../../cookie";
import { post } from "../../rest_client";
import { APP_ROUTES } from "../../routes";

export async function login(inputs: LoginInputs): Promise<string | void> {
  const res: any = await post<{ token: string }>("/login", new URLSearchParams(inputs));

  if (res.error) {
    return res.error;
  }

  if (!res.data || !res.data.token) {
    return "Something went wrong!";
  }

  Cookie.set(COOKIES.authToken, res.data.token);
  await Router.push(APP_ROUTES.admin.dashboard);
}