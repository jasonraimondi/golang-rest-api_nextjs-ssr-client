import { Component } from "react";
import { AuthProps, privateRoute } from "../lib/auth/private_route";

class Page extends Component<AuthProps> {
  static async getInitialProps({ auth }: AuthProps) {
    auth.logout();
  }

  render() {
    return "Logging Out...";
  }
}

export default privateRoute(Page);
