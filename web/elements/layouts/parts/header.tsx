import Link from "next/link";
import React from "react";
import { APP_ROUTES } from "../../../lib/routes";
import { AuthToken } from "../../../lib/services/auth_token";

interface Props {
  auth?: AuthToken,
}

const Header = (props: Props) => {
  let isAuthenticated = false;
  if (props.auth && props.auth.isValid) {
    isAuthenticated = true;
  }

  return <header>
    <nav>
      <ul className="flex justify-around">
        <li>
          <Link href={APP_ROUTES.home}>
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
    <li>
      <Link href={APP_ROUTES.admin.dashboard}>
        <a>Dashboard</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.admin.photos.index}>
        <a>Photo</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.admin.photos.upload}>
        <a>Upload</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.auth.logout}>
        <a>Logout</a>
      </Link>
    </li>
  </>;
}

function PublicRoutes() {
  return <>
    <li>
      <Link href={APP_ROUTES.auth.login}>
        <a>Login</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.signUp}>
        <a>Sign Up</a>
      </Link>
    </li>
  </>;
}

export default Header;
