import { ServerResponse } from "http";
import Cookie from "js-cookie";
import decode from "jwt-decode";
import Router from "next/router";
import { COOKIES } from "../cookie";

export interface DecodedToken {
  user_id: string;
  email: string;
  exp: number;
}

export class AuthService {
  public readonly user: any;
  public readonly authorizationString: string;
  public readonly expiresAt: Date;
  public readonly isExpired: boolean;
  public readonly isAuthenticated: boolean;

  private decodedToken: DecodedToken;

  constructor(private token?: string) {
    if (!this.token) this.token = "ERR";
    try {
      this.decodedToken = decode(this.token);
    } catch (e) {
      this.decodedToken = this.blankToken;
    }
    this.user = {
      id: this.decodedToken.user_id,
      email: this.decodedToken.email,
    };
    this.authorizationString = `Bearer ${this.token}`;
    this.expiresAt = new Date(this.decodedToken.exp * 1000);
    this.isExpired = new Date() > this.expiresAt;
    this.isAuthenticated = !this.isExpired;
    this.logout = this.logout.bind(this);
  }

  get blankToken(): DecodedToken {
    return {
      user_id: "",
      email: "",
      exp: 0,
    };
  }

  static redirectToLogin(res?: ServerResponse) {
    if (res) {
      res.writeHead(302, {
        Location: "/login",
      });
      res.end();
    } else {
      Router.push("/login");
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
    this.decodedToken = this.blankToken;
    Router.push("/login");
  }
}