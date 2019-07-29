import ClientCookie from "js-cookie";
import ServerCookie from "next-cookies";
import Router from "next/router";
import React, { Component } from "react";
import { COOKIES } from "../cookie";
import { AuthService } from "./auth_service";

export type AuthProps = {
  auth: AuthService
}

export function privateRoute(C: any) {
  const authService = new AuthService(ClientCookie.get(COOKIES.authToken));

  return class extends Component {
    static async getInitialProps(ctx: any) {
      const cookies = ServerCookie(ctx);
      const props = { auth: new AuthService(cookies[COOKIES.authToken]) };
      if (C.getInitialProps) return await C.getInitialProps(props);
      return props;
    }

    componentDidMount(): void {
      if (authService.isExpired) Router.push("/");
    }

    componentDidUpdate(): void {
      if (authService.isExpired) Router.push("/");
    }

    render() {
      return <C auth={authService} {...this.props} />;
    }
  };
}
