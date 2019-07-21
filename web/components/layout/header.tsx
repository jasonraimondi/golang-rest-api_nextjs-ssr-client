import Link from "next/link";
import React from "react";

// The Header creates links that can be used to navigate
// between routes.
const Header = () => (
  <header>
    <nav>
      <ul style={{ listStyleType: "none" }}>
        <li>
          <Link href="/">
            <a>Home</a>
          </Link>
        </li>
        <li>
          <Link href="/login">
            <a>Login</a>
          </Link>
        </li>
        <li>
          <Link href="/app/upload_photos">
            <a>Upload Photos</a>
          </Link>
        </li>
        <li>
          <Link href="/sign_up">
            <a>SignUp</a>
          </Link>
        </li>
      </ul>
    </nav>
  </header>
);

export default Header;
