import React, { Component } from "react";
import { signUp } from "../lib/services/sign_up";

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
  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleInputChange = this.handleInputChange.bind(this);
    this.state = {
      inputs: {
        email: "",
      },
    };
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

  render() {
    const inputs = this.state.inputs;
    return <>
      <form onSubmit={this.handleSubmit}>
        <div>
          <label>
            First Name
            <input type="text" name="first" onChange={this.handleInputChange} value={inputs.first}/>
          </label>
        </div>
        <div>
          <label>
            Last Name
            <input type="text" name="last" onChange={this.handleInputChange} value={inputs.last}/>
          </label>
        </div>
        <div>
          <label>
            Email Address
            <input type="email" name="email" onChange={this.handleInputChange} value={inputs.email}
                   required/>
          </label>
        </div>
        <div>
          <label>
            Password
            <input type="password" name="password" onChange={this.handleInputChange}
                   value={inputs.password}/>
          </label>
        </div>
        <button type="submit">Sign Up</button>
      </form>
    </>;
  }
}