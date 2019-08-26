import { SignUpInputs } from "../../../pages/sign_up";
import { get, post } from "../../rest_client";
import { API_ROUTES } from "../../routes";

export async function signUp(inputs: SignUpInputs) {
  return await post(API_ROUTES.sign_up.create({}), new URLSearchParams(inputs));
}

export async function signUpConfirmation(inputs: { u: string, t: string }) {
  return await get(API_ROUTES.sign_up_confirmation.create({}), inputs);
}
