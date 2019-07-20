import {SignUpForm} from "../../pages";
import {appRestClient} from "../rest_client";

export async function signUp(inputs: SignUpForm): Promise<string> {
    const foo = await appRestClient.post("/sign-up", inputs).catch(catchAxiosError);
    console.log(foo);
    return "hello sunshine my best friend";
}

function catchAxiosError(error: any) {
    if (error.response) {
        // The request was made and the server responded with a status code
        // that falls out of the range of 2xx
        console.log(error.response.data.message);
        console.log(error.response.status);
        console.log(error.response.headers);
    } else if (error.request) {
        // The request was made but no response was received
        // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
        // http.ClientRequest in node.js
        console.log(error.request);
    } else {
        // Something happened in setting up the request that triggered an Error
        console.log("Error", error.message);
    }
    console.log(error.config);
}
