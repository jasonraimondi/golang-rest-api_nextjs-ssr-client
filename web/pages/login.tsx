import React from "react";
import { AuthService } from "../components/auth/auth_service";
import { defaultLayout } from "../components/layouts/default";
import { LoginForm } from "../components/login_form";

function LoginPage() {
  AuthService.redirectIfAuthenticated();

  return (
    <>
      <LoginForm/>
    </>
  );
}


export default defaultLayout(LoginPage);
