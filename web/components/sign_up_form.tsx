import React, { useState } from "react";
import { SubmitButton } from "../elements/forms/button";
import { TextInput } from "../elements/forms/text";
import { signUp } from "../lib/services/api/sign_up";

export interface SignUp {
  email: string
  first: string
  last: string
  password: string
}

export function SignUpForm() {
  const [inputs, setInputs] = useState({
    email: "",
    first: "",
    last: "",
    password: "",
  } as SignUp);

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    await signUp(inputs);
  };

  const handleInputChange = (e: any) => {
    e.persist();
    setInputs({
      ...inputs,
      [e.target.name]: e.target.value,
    })
  };

  return <>
    <form className="container mx-auto max-w-sm" onSubmit={handleSubmit}>
      <TextInput type="text"
                 label="First"
                 name="first"
                 handleInputChange={handleInputChange}
                 value={inputs.first}
      />
      <TextInput type="text"
                 label="Last"
                 name="last"
                 handleInputChange={handleInputChange}
                 value={inputs.last}
      />
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
      <SubmitButton label="Sign Up"/>
    </form>
  </>;
}

