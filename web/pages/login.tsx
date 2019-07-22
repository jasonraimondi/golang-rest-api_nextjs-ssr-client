import React from "react";
import { LoginForm } from "../components/login_form";
import { defaultLayout } from "../elements/layouts/default";
import { AuthService } from "../lib/auth/auth_service";

function LoginPage() {
  AuthService.redirectIfAuthenticated();
  return <LoginForm/>;
}

export default defaultLayout(LoginPage);
