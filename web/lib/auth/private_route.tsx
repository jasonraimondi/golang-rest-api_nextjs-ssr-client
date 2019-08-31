import { NextPageContext } from "next";
import React, { Component } from "react";
import { AuthToken } from "../services/auth_token";
import { redirectToLogin } from "../services/redirect_service";


export type AuthProps = Props & {
  auth: AuthToken
}

type Props = {
  token: string;
}

export function privateRoute(Page: any) {
  return class extends Component<AuthProps> {
    state = {
      auth: new AuthToken(this.props.token),
    };

    static async getInitialProps(ctx: NextPageContext) {
      const auth = AuthToken.fromNext(ctx);
      if (auth.isExpired) await redirectToLogin(ctx.res);
      return {
        ...(Page.getInitialProps ? await Page.getInitialProps(ctx) : {}),
        token: auth.token,
      };
    }

    componentDidMount(): void {
      // since getInitialProps returns our props after they've JSON.stringify
      // we need to reinitialize it as an AuthToken to have the full class
      // with all instance methods available
      this.setState({ auth: new AuthToken(this.props.token) });
    }

    render() {
      const { auth, ...propsWithoutAuth } = this.props;
      return <Page auth={this.state.auth} {...propsWithoutAuth} />;
    }
  };
}