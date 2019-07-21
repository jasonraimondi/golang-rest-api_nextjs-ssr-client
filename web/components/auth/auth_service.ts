import Cookie from "js-cookie";
import decode from "jwt-decode";
import Router from "next/router";
import { COOKIES } from "../../lib/cookie";
import { appRestClient } from "../../lib/rest_client";
import { LoginInputs } from "../login_form";

export interface DecodedToken {
  user_id: string;
  email: string;
  exp: number;
}

export class AuthService {
  public readonly authorizationString: string;

  private decodedJWT: DecodedToken;

  constructor(private token?: string) {
    if (!this.token) this.token = "ERR";
    try {
      this.decodedJWT = decode(this.token);
    } catch (e) {
      this.decodedJWT = this.blankToken;
    }

    this.authorizationString = `Bearer ${this.token}`;
    this.logout = this.logout.bind(this);
  }

  get user() {
    return {
      id: this.decodedJWT.user_id,
      email: this.decodedJWT.email,
    };
  }

  get isAuthenticated(): boolean {
    return !this.isExpired;
  }

  get isExpired(): boolean {
    return new Date() > this.expiresAt;
  }

  get expiresAt(): Date {
    return new Date(this.decodedJWT.exp * 1000);
  }

  get blankToken(): DecodedToken {
    return {
      user_id: "",
      email: "",
      exp: 0,
    };
  }

  static async login(inputs: LoginInputs) {
    const res = await appRestClient.post<{ token: string }>("/login", inputs);
    if (res.data.token) {
      Cookie.set(COOKIES.authToken, res.data.token);
      Router.push("/app/dashboard");
    }
  }

  static redirectIfAuthenticated() {
    const authService = new AuthService(Cookie.get(COOKIES.authToken));
    if (authService.isAuthenticated) {
      Router.push("/app/dashboard");
    }
  }

  logout() {
    Cookie.remove(COOKIES.authToken);
    this.decodedJWT = this.blankToken;
    Router.push("/login");
  }
}