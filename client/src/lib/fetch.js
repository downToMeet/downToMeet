import { SERVER_URL } from "../constants";

export async function getUserData(id) {
  const getUserDataEndpoint = `${SERVER_URL}/user/${id}`;
  const res = await fetch(getUserDataEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export const doOtherThing = () => {
  return null;
};
