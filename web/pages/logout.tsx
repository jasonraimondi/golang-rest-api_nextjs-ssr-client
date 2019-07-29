import { Component } from "react";
import { AuthProps, privateRoute } from "../lib/auth/private_route";

class Page extends Component<AuthProps> {
  componentDidMount(): void {
    this.props.auth.logout();
  }

  render() {
    return "Logging Out...";
  }
}

export default privateRoute(Page);
