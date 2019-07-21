import React from "react";
import { PrivateRoute } from "../../components/auth/private_route";
import Header from "../../components/layout/header";

function Dashboard() {
  return <>
    <Header/>
    <p>Hello Dashboard</p>
  </>;
}

export default PrivateRoute(Dashboard);
