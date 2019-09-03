import { NextPageContext } from "next";
import React from "react";
import { AuthToken } from "../../lib/services/auth_token";
import { redirectToLogin } from "../../lib/services/redirect_service";

export type AuthProps = {
  token: string
}

export function privateRoute(Page: any) {

  const PrivateRoute = (props: any) => <Page {...props} />;

  PrivateRoute.getInitialProps = async (ctx: NextPageContext) => {
    const auth = AuthToken.fromNext(ctx);
    if (auth.isExpired) await redirectToLogin(ctx.res);
    return {
      ...(Page.getInitialProps ? await Page.getInitialProps(ctx) : {}),
      token: auth.token,
    };
  };

  return PrivateRoute;
}