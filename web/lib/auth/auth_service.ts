import { ServerResponse } from "http";
import Cookie from "js-cookie";
import decode from "jwt-decode";
import Router from "next/router";
import { LoginInputs } from "../../pages/login";
import { COOKIES } from "../cookie";
import { post } from "../rest_client";

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

  static async login(inputs: LoginInputs | any): Promise<string | void> {
    try {
      const res: any = await post<{ token: string }>("/login", inputs);
      if (res.data && res.data.token) {
        Cookie.set(COOKIES.authToken, res.data.token);
        Router.push("/app/dashboard");
      }
      return;
    } catch (err) {
      if (!err.response) {
        return err.message;
      } else if (err.response.status === 404) {
        return "user not found";
      } else if (err.response.data && err.response.data.message) {
        return err.response.data;
      }
    }
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
    this.decodedJWT = this.blankToken;
    Router.push("/login");
  }
}