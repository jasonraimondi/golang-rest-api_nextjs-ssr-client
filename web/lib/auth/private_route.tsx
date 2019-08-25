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

export function privateRoute(WrappedComponent: any) {
  return class extends Component<AuthProps> {
    state = {
      auth: new AuthToken(this.props.auth.token),
    };

    static async getInitialProps(ctx: NextPageContext) {
      const auth = AuthToken.fromNext(ctx);
      const initialProps = { ...ctx, auth };
      if (auth.isExpired) redirectToLogin(ctx.res);
      if (WrappedComponent.getInitialProps) {
        const wrappedProps = await WrappedComponent.getInitialProps(initialProps);
        // make sure our `auth: AuthToken` is always returned
        return { ...wrappedProps, auth };
      }
      return initialProps;
    }

    componentDidMount(): void {
      // since getInitialProps returns our props after they've JSON.stringify
      // we need to reinitialize it as an AuthToken to have the full class
      // with all instance methods available
      this.setState({ auth: new AuthToken(this.props.auth.token) });
    }

    render() {
      const { auth, ...propsWithoutAuth } = this.props;
      return <WrappedComponent auth={this.state.auth} {...propsWithoutAuth} />;
    }
  };
}