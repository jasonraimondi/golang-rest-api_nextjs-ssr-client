import { Component } from "react";
import { AuthService } from "../lib/auth/auth_service";
import { privateRoute } from "../lib/auth/private_route";

class Page extends Component<{ auth: AuthService }> {
  componentDidMount(): void {
    this.props.auth.logout();
  }

  render() {
    return "Logging Out...";
  }
}

export default privateRoute(Page);
