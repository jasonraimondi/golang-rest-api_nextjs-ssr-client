import React from "react";
import { AuthService } from "../../lib/auth/auth_service";
import { privateRoute } from "../../lib/auth/private_route";
import { defaultLayout } from "../../elements/layouts/default";

type AuthProps = {
  auth: AuthService
}

type Props = {}

function DashboardPage({ auth }: Props & AuthProps) {
  console.log(auth);
  return <>
    <p>{JSON.stringify(auth.user)}</p>
  </>;
}

export default privateRoute(defaultLayout(DashboardPage));
