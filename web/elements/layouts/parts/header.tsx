import Link from "next/link";
import React from "react";
import { AuthService } from "../../../lib/auth/auth_service";
import { APP_ROUTES } from "../../../lib/routes";

interface Props {
  auth?: AuthService,
}

const Header = (props: Props) => {
  let isAuthenticated = false;
  if (props.auth && props.auth.isAuthenticated) {
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
      <Link href={APP_ROUTES.dashboard}>
        <a>Dashboard</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.photos}>
        <a>Photo</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.photosUpload}>
        <a>Upload</a>
      </Link>
    </li>
    <li>
      <Link href={APP_ROUTES.logout}>
        <a>Logout</a>
      </Link>
    </li>
  </>;
}

function PublicRoutes() {
  return <>
    <li>
      <Link href={APP_ROUTES.login}>
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
