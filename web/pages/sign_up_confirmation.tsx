import { NextPageContext } from "next";
import Router from "next/router";
import React, { Component } from "react";
import { defaultLayout } from "../elements/layouts/default";
import { APP_ROUTES } from "../lib/routes";
import { signUpConfirmation } from "../lib/services/api/sign_up";
import { redirectIfAuthenticated, redirectToLogin } from "../lib/services/redirect_service";

type State = { isValid: boolean, isLoading: boolean };
type Props = { userId: string, token: string };

class Page extends Component<Props, State> {
  state = {
    isValid: false,
    isLoading: true,
  };

  constructor(props: Props) {
    super(props);
  }

  static async getInitialProps(ctx: NextPageContext) {
    await redirectIfAuthenticated(ctx);

    const { t, u } = ctx.query;
    if (!t || !u) redirectToLogin(ctx.res);
    return { token: t, userId: u };
  };

  async componentDidMount(): Promise<void> {
    const res: any = await signUpConfirmation({
      t: this.props.token,
      u: this.props.userId,
    });

    if (res.status === 202) {
      this.setState({ isValid: true, isLoading: false });
      redirectToLogin();
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

export default defaultLayout(Page);
