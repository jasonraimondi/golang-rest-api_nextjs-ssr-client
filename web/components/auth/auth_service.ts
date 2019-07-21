import decode from "jwt-decode";

export interface DecodedToken {
  user_id: string;
  email: string;
  exp: number;
}

export class AuthService {
  private token: DecodedToken;

  constructor(jwt: string) {
    try {
      this.token = decode(jwt);
    } catch (e) {
      this.token = this.blankToken;
    }
  }

  get auth() {
    return {
      userId: this.token.user_id,
      email: this.token.email,
    };
  }

  get isAuthenticated(): boolean {
    return !this.isExpired;
  }

  get isExpired(): boolean {
    return new Date() > this.expiresAt;
  }

  get expiresAt(): Date {
    return new Date(this.token.exp * 1000);
  }

  get blankToken(): DecodedToken {
    return {
      user_id: "",
      email: "",
      exp: 0,
    };
  }

  logout() {
    this.token = this.blankToken;
  }
}