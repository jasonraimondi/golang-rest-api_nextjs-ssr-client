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


function catchAxiosError(error: any) {
    if (error.response) {
        // The request was made and the server responded with a status code
        // that falls out of the range of 2xx
        console.log(error.response.data);
        console.log(error.response.status);
        console.log(error.response.headers);
    } else if (error.request) {
        // The request was made but no response was received
        // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
        // http.ClientRequest in node.js
        console.log(error.request);
    } else {
        // Something happened in setting up the request that triggered an Error
        console.log('Error', error.message);
    }
    console.log(error.config);
}

function Home() {
    const rand = () => {
        const min=0;
        const max=1000;
        return Math.floor(Math.random() * (+max - +min)) + +min;
    };

    const [inputs, setInputs] = useState<SignUpForm>({
        email: `jason${rand()}@raimondi.us`,
    });


    const handleSubmit = async (event: any) => {
        event.preventDefault();
        const res = await appRestClient.post('/sign-up', inputs).catch(catchAxiosError);
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
                        <input type="email" name="email" onChange={handleInputChange} value={inputs.email} required/>
                    </label>
                </div>
                <div>
                    <label>
                        Password
                        <input type="password" name="password" onChange={handleInputChange} value={inputs.password}/>
                    </label>
                </div>
                <button type="submit">Sign Up</button>
            </form>
        </div>
    );
}

export default Home;
