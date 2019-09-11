import { NextPage } from "next";
import React, { useEffect, useState } from "react";

import { AuthToken } from "@/lib/services/auth_token";
import { redirectToLogin } from "@/lib/services/redirect_service";

export type AuthProps = {
  token: string
  auth: AuthToken
}

export function privateRoute(Page: any) {

  const PrivateRoute: NextPage<any> = (props: any) => {
    const [auth, setAuth] = useState(AuthToken.fromToken(props.token));

    // componentDidMount
    useEffect(() => {
      setAuth(AuthToken.fromToken(props.token));
    }, []);

    return <Page auth={auth} {...props} />;
  };

  PrivateRoute.getInitialProps = async (ctx) => {
    const auth = AuthToken.fromNext(ctx);
    if (auth.isExpired) await redirectToLogin(ctx.res);
    return {
      ...(Page.getInitialProps ? await Page.getInitialProps(ctx) : {}),
      token: auth.token,
    };
  };

  return PrivateRoute;
}