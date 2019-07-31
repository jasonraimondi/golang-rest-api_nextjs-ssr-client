import ServerCookie from "next-cookies";
import Router from "next/router";
import React, { Component } from "react";
import { COOKIES } from "../cookie";
import { APP_ROUTES } from "../routes";
import { AuthService } from "./auth_service";

export type AuthProps = {
  auth: AuthService
}

export function privateRoute(C: any) {
  return class extends Component<AuthProps> {
    static async getInitialProps(ctx: any) {
      const jwt = ServerCookie(ctx)[COOKIES.authToken];
      const props = { auth: new AuthService(jwt) };
      if (C.getInitialProps) return C.getInitialProps(props);
      return props;
    }

    componentDidMount(): void {
      if (this.props.auth.isExpired) Router.push(APP_ROUTES.home);
    }

    componentDidUpdate(): void {
      if (this.props.auth.isExpired) Router.push(APP_ROUTES.home);
    }

    render() {
      return <C auth={this.props.auth} {...this.props} />;
    }
  };
}
