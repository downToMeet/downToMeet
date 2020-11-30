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

export async function searchForMeetups({ lat, lon, radius, tags }) {
  const getMeetupsEndpoint = `${SERVER_URL}/meetup?lat=${lat}&lon=${lon}&radius=${radius}&tags=${tags}`;
  const res = await fetch(getMeetupsEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}
