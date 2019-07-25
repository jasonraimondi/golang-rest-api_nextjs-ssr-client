import Link from "next/link";
import React from "react";
import { AuthService } from "../../../lib/auth/auth_service";

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
          <Link href="/">
            <a>Home</a>
          </Link>
        </li>
        {isAuthenticated ? <PrivateRoutes/> : <PublicRoutes/>}
      </ul>
    </nav>
  </header>;
};

Header.getInitialProps = async ({ req }) => {
  console.log(req, "FETCHING");
  const res = await fetch("https://api.github.com/repos/zeit/next.js");
  const json = await res.json();
  return { stars: json.stargazers_count };
};

function PrivateRoutes() {
  console.log("hi private route");
  return <>
    <li>
      <Link href="/app/dashboard">
        <a>Dashboard</a>
      </Link>
    </li>
    <li>
      <Link href="/app/upload_photos">
        <a>Upload Photo</a>
      </Link>
    </li>
    <li>
      <Link href="/logout">
        <a>Logout</a>
      </Link>
    </li>
  </>;
}

function PublicRoutes() {
  return <>
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
