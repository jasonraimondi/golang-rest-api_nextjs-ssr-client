import React from "react";
import { AuthService } from "../lib/auth/auth_service";
import { defaultLayout } from "../elements/layouts/default";
import { LoginForm } from "../components/login_form";

function LoginPage() {
  AuthService.redirectIfAuthenticated();
  return <LoginForm/>;
}

export default defaultLayout(LoginPage);
