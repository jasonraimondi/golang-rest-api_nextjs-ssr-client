import { ServerResponse } from "http";
import Cookie from "js-cookie";
import jwtDecode from "jwt-decode";
import { NextPageContext } from "next";
import ServerCookie from "next-cookies";
import Router from "next/router";
import { COOKIES } from "../cookie";
import { APP_ROUTES } from "../routes";
import { redirectToLogin } from "./redirect_service";


export interface DecodedToken {
  readonly userId: string;
  readonly email: string;
  readonly isValid: boolean;
  readonly isConfirmed: boolean;
  readonly exp: number;
}

export class AuthToken {
  readonly decodedToken: DecodedToken;
  public logout = AuthToken.logout;

  constructor(readonly token?: string) {
    this.decodedToken = { userId: "anonymous", email: "anonymous", exp: 0, isValid: false, isConfirmed: false };
    try {
      if (token) this.decodedToken = jwtDecode(token);
    } catch (e) {
    }
  }

  get user() {
    const { userId, email, isConfirmed, isValid } = this.decodedToken;
    return { id: userId, email, isConfirmed, isValid };
  }

  get authorizationString() {
    return `Bearer ${this.token}`;
  }

  get expiresAt(): Date {
    return new Date(this.decodedToken.exp * 1000);
  }

  get isExpired(): boolean {
    return new Date() > this.expiresAt;
  }

  get isValid(): boolean {
    return !this.isExpired;
  }

  static fromNext(ctx: NextPageContext) {
    const token = ServerCookie(ctx)[COOKIES.authToken];
    return new AuthToken(token);
  }

  static async logout(server?: ServerResponse) {
    Cookie.remove(COOKIES.authToken);
    await redirectToLogin(server);
  };

  static async storeToken(token: string) {
    Cookie.set(COOKIES.authToken, token);
    await Router.push(APP_ROUTES.admin.dashboard.create());
  }
}