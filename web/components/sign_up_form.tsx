import React, { Component } from "react";
import { signUp } from "../lib/services/sign_up";
import { SubmitButton } from "./forms/button";
import { TextInput } from "./forms/text";

interface Props {
  setMessage: (message: string) => void
  setSubmitted: (isSubmitted: boolean) => void
}

export interface SignUp {
  email: string
  first?: string
  last?: string
  password?: string
}

interface State {
  inputs: SignUp;
}

export class SignUpForm extends Component<Props, State> {
  state = {
    inputs: {
      email: "",
    } as SignUp,
  };

  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleInputChange = this.handleInputChange.bind(this);
  }

  render() {
    const { inputs } = this.state;
    return <>
      <form className="container mx-auto max-w-sm" onSubmit={this.handleSubmit}>
        <TextInput type="text"
                   label="First"
                   name="first"
                   handleInputChange={this.handleInputChange}
                   value={inputs.first}
        />
        <TextInput type="text"
                   label="Last"
                   name="last"
                   handleInputChange={this.handleInputChange}
                   value={inputs.last}
        />
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
        <SubmitButton label="Sign Up"/>
      </form>
    </>;
  }

  private async handleSubmit(e: any) {
    e.preventDefault();
    this.props.setMessage(await signUp(this.state.inputs));
    this.props.setSubmitted(true);
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

