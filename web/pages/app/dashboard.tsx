import React from "react";
import { AuthService } from "../../components/auth/auth_service";
import { privateRoute } from "../../components/auth/private_route";
import { defaultLayout } from "../../components/layouts/default";

type AuthProps = {
  auth: AuthService
}

type Props = {}

function DashboardPage({ auth }: Props & AuthProps) {
  console.log(auth);
  return <>
    <p>Hello Dashboard</p>
    <p>{auth.authorizationString}</p>
  </>;
}

export default privateRoute(defaultLayout(DashboardPage));
