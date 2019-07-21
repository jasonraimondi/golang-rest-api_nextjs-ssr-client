import { Component } from "react";
import { AuthService } from "../components/auth/auth_service";
import { privateRoute } from "../components/auth/private_route";

class LogoutPage extends Component<{ auth: AuthService }> {
  componentDidMount(): void {
    this.props.auth.logout();
  }
  render() {
    return "Logging Out..."
  }
}

export default privateRoute(LogoutPage);
