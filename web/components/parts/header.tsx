import Link from "next/link";
import React from "react";
import { AuthService } from "../auth/auth_service";

interface Props {
  auth?: AuthService,
}

// The Header creates links that can be used to navigate
// between routes.
const Header = (props: Props) => {
  let isAuthenticated = false;
  if (props.auth && props.auth.isAuthenticated) {
    isAuthenticated = true;
  }

  return <header>
    <nav>
      <ul style={{ listStyleType: "none" }}>
        <li>
          <Link href="/">
            <a>Home</a>
          </Link>
        </li>
        {isAuthenticated ? <PrivateRoutes/> : <PublicRoutes/>}
      </ul>
    </nav>
  </header>;
};

function PrivateRoutes() {
  return <>
    <li>PRIVATE ROUTES</li>
    <li>
      <Link href="/app/dashboard">
        <a>Dashboard</a>
      </Link>
      <Link href="/logout">
        <a>Logout</a>
      </Link>
    </li>
  </>;
}

function PublicRoutes() {
  return <>
    <li>Public ROUTES</li>
    <li>
      <Link href="/login">
        <a>Login</a>
      </Link>
    </li>
    <li>
      <Link href="/sign_up">
        <a>SignUp</a>
      </Link>
    </li>
  </>;
}

export default Header;
