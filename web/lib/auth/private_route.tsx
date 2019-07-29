import ServerCookie from "next-cookies";
import Router from "next/router";
import React, { Component } from "react";
import { COOKIES } from "../cookie";
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

    get auth() {
      return this.props.auth;
    }

    componentDidMount(): void {
      if (this.auth.isExpired) Router.push("/");
    }

    componentDidUpdate(): void {
      if (this.auth.isExpired) Router.push("/");
    }

    render() {
      return <C auth={this.auth} {...this.props} />;
    }
  };
}
