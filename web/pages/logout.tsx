import { Component } from "react";
import { AuthToken } from "../lib/services/auth_token";

class Page extends Component {
  componentDidMount(): void {
    AuthToken.logout();
  }

  render() {
    return "Logging Out...";
  }
}

export default Page;
