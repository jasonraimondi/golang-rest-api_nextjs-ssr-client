import Cookie from "js-cookie";
import Router from "next/router";

import { LoginInputs } from "@/components/auth/login_form";
import { COOKIES } from "@/lib/cookie";
import { post } from "@/lib/rest_client";
import { API_ROUTES, APP_ROUTES } from "@/lib/routes";

export async function login(inputs: LoginInputs): Promise<string | void> {
  const res: any = await post<{ token: string }>(API_ROUTES.login.create(), new URLSearchParams(inputs));

  if (res.error) {
    return res.error;
  }

  if (res.data && res.data.token) {
    Cookie.set(COOKIES.authToken, res.data.token);
    await Router.push(APP_ROUTES.admin.dashboard.create());
    return;
  }

  return "Something went wrong!";
}