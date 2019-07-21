import { privateRoute } from "../components/auth/private_route";

function LogoutPage({ auth }) {
  auth.logout();
}

export default privateRoute(LogoutPage);
