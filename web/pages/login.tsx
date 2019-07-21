import React from "react";
import { defaultLayout } from "../components/layouts/default";

import { LoginForm } from "../components/login_form";

function LoginPage() {
  return (
    <>
      <LoginForm/>
    </>
  );
}


export default defaultLayout(LoginPage);
