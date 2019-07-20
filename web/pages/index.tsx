import React, {useState} from "react";

import {GlobalHeader} from "../components/head";
import {signUp} from "../lib/services/sign_up";

export interface SignUpForm {
    email: string
    first?: string
    last?: string
    password?: string
}


function Home() {
    const rand = () => {
        const min = 0;
        const max = 1000;
        return Math.floor(Math.random() * (+max - +min)) + +min;
    };
    const [message, setMessage] = useState("");
    const [submitted, setSubmitted] = useState(false);
    const [inputs, setInputs] = useState<SignUpForm>({
        email: `jason${rand()}@raimondi.us`,
    });

    const handleSubmit = async (e: any) => {
        e.preventDefault();
        setMessage(await signUp(inputs));
        setSubmitted(true);
    };

    const handleInputChange = (e: any) => {
        e.persist();
        setInputs(inputs => ({...inputs, [e.target.name]: e.target.value}));
    };

    return (
        <div>
            <GlobalHeader/>
            <h1>This page has a title {inputs.first}🤔</h1>
            {submitted ? message : (
                <form onSubmit={handleSubmit}>
                    <div>
                        <label>
                            First Name
                            <input type="text" name="first" onChange={handleInputChange} value={inputs.first}/>
                        </label>
                    </div>
                    <div>
                        <label>
                            Last Name
                            <input type="text" name="last" onChange={handleInputChange} value={inputs.last}/>
                        </label>
                    </div>
                    <div>
                        <label>
                            Email Address
                            <input type="email" name="email" onChange={handleInputChange} value={inputs.email}
                                   required/>
                        </label>
                    </div>
                    <div>
                        <label>
                            Password
                            <input type="password" name="password" onChange={handleInputChange}
                                   value={inputs.password}/>
                        </label>
                    </div>
                    <button type="submit">Sign Up</button>
                </form>
            )}
        </div>
    );
}

export default Home;
