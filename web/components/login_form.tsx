import React, { Component } from "react";
import { AuthService } from "../lib/auth/auth_service";
import { SubmitButton } from "../elements/forms/button";
import { TextInput } from "../elements/forms/text";

interface Props {
}

export interface LoginInputs {
  email: string
  password: string
}

interface State {
  inputs: LoginInputs;
}

export class LoginForm extends Component<Props, State> {
  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleInputChange = this.handleInputChange.bind(this);
    this.state = {
      inputs: {
        email: "",
        password: "",
      },
    };
  }

  render() {
    const inputs = this.state.inputs;
    return <>
      <form className="container mx-auto max-w-sm" onSubmit={this.handleSubmit}>
        <TextInput type="email"
                   label="Email"
                   name="email"
                   handleInputChange={this.handleInputChange}
                   value={inputs.email} required
        />
        <TextInput type="password"
                   label="Password"
                   name="password"
                   handleInputChange={this.handleInputChange}
                   value={inputs.password}
        />
        <SubmitButton label="Login"/>
      </form>
    </>;
  }

  private async handleSubmit(e: any) {
    e.preventDefault();
    await AuthService.login(this.state.inputs);
  };

  private handleInputChange(e: any) {
    e.persist();
    this.setState({
      inputs: {
        ...this.state.inputs,
        [e.target.name]: e.target.value,
      },
    });
  };
}