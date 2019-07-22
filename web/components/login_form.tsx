import React, { useState } from "react";
import { SubmitButton } from "../elements/forms/button";
import { TextInput } from "../elements/forms/text";
import { AuthService } from "../lib/auth/auth_service";

export interface LoginInputs {
  email: string
  password: string
}

export function LoginForm() {
  const [inputs, setInputs] = useState({
    email: "",
    password: "",
  } as LoginInputs);

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    await AuthService.login(inputs);
  };

  const handleInputChange = (e: any) => {
    e.persist();
    setInputs({
      ...inputs,
      [e.target.name]: e.target.value,
    });
  };

  return <>
    <form className="container mx-auto max-w-sm" onSubmit={handleSubmit}>
      <TextInput type="email"
                 label="Email"
                 name="email"
                 handleInputChange={handleInputChange}
                 value={inputs.email} required
      />
      <TextInput type="password"
                 label="Password"
                 name="password"
                 handleInputChange={handleInputChange}
                 value={inputs.password}
      />
      <SubmitButton label="Login"/>
    </form>
  </>;
}