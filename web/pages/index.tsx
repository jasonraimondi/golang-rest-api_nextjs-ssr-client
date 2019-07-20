import React, {useState} from 'react';
import axios from 'axios';

import {AppRestClient, AxiosRestClient} from "../lib/rest_client";

interface SignUpForm {
    email: string
    first?: string
    last?: string
    password?: string
}

const http = axios.create({
    baseURL: 'http://localhost:1323'
});
const restClient = new AxiosRestClient(http);
const appRestClient = new AppRestClient(restClient);

function Home() {
    const [inputs, setInputs] = useState<SignUpForm>({
        email: "Z",
    });
    const handleSubmit = async (event: any) => {
        event.preventDefault();
        const res = await appRestClient.post('/sign-up', inputs);
        console.log(res);
    };
    const handleInputChange = (event: any) => {
        event.persist();
        setInputs(inputs => ({...inputs, [event.target.name]: event.target.value}));
    };
    return (
        <div>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>First Name</label>
                    <input type="text" name="first" onChange={handleInputChange} value={inputs.first}/>
                    <label>Last Name</label>
                    <input type="text" name="last" onChange={handleInputChange} value={inputs.last}/>
                </div>
                <div>
                    <label>Email Address</label>
                    <input type="email" name="email" onChange={handleInputChange} value={inputs.email} required/>
                </div>
                <div>
                    <label>Password</label>
                    <input type="password" name="password" onChange={handleInputChange} value={inputs.password}/>
                </div>
                <button type="submit">Sign Up</button>
            </form>
        </div>
    );
}

export default Home;
