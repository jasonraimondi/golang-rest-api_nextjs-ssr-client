import Router from "next/router";
import React, { Component } from "react";
import { defaultLayout } from "../elements/layouts/default";
import { AuthService } from "../lib/auth/auth_service";
import { APP_ROUTES } from "../lib/routes";
import { signUpConfirmation } from "../lib/services/api/sign_up";

type State = { isValid: boolean, isLoading: boolean };
type Props = { userId: string, token: string };

class Page extends Component<Props, State> {
  state = {
    isValid: false,
    isLoading: true,
  };

  constructor(props: Props) {
    super(props);
    AuthService.redirectIfAuthenticated();
  }

  static async getInitialProps({ res, query }: any) {
    const { t, u } = query;
    if (!t || !u) AuthService.redirectToLogin(res);
    return { token: t, userId: u };
  };

  async componentDidMount(): Promise<void> {
    const res: any = await signUpConfirmation({
      t: this.props.token,
      u: this.props.userId,
    });

    if (res.status === 202) {
      this.setState({ isValid: true, isLoading: false });
      AuthService.redirectToLogin();
    } else {
      this.setState({ isValid: false, isLoading: false });
      await Router.push(APP_ROUTES.signUp);
    }
  }

  render() {
    const { isValid, isLoading } = this.state;
    if (isLoading) {
      return <div>Loading...</div>;
    }

    if (isValid && !isLoading) {
      return <div>Valid, redirecting to login.</div>;
    }

    return <div>Token is invalid.</div>;
  }
}

function redirectToLogin(res?: ServerResponse) {
    if (res) {
        res.writeHead(302, {
            Location: APP_ROUTES.auth.login,
        });
        res.end();
    } else {
        Router.push(APP_ROUTES.auth.login);
    }
}

export default defaultLayout(Page);
