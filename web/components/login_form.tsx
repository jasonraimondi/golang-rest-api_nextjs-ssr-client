import React, { Component } from "react";
import { AuthService } from "./auth/auth_service";

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

  render() {
    const inputs = this.state.inputs;
    return <>
      <form onSubmit={this.handleSubmit}>
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
        <button type="submit">Login</button>
      </form>
    </>;
  }
}