import { IN_PERSON, REMOTE, SERVER_URL } from "../constants";

export async function getUserData(id) {
  const getUserDataEndpoint = `${SERVER_URL}/user/${id}`;
  const res = await fetch(getUserDataEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export async function searchForMeetups({ lat, lon, radius, tags }) {
  const getMeetupsEndpoint = `${SERVER_URL}/meetup?lat=${lat}&lon=${lon}&radius=${radius}&tags=${tags}`;
  const res = await fetch(getMeetupsEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export async function searchForRemoteMeetups(tags) {
  const getMeetupsEndpoint = `${SERVER_URL}/meetup/remote?tags=${tags}`;
  const res = await fetch(getMeetupsEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export async function createMeetup({
  title,
  time,
  meetupType,
  meetupURL,
  meetupLocation,
  groupCount,
  description,
  tags,
}) {
  const createMeetupEndpoint = `${SERVER_URL}/meetup`;

  let location;
  if (meetupType === IN_PERSON) {
    location = {
      coordinates: {
        lat: meetupLocation.coords[0],
        lon: meetupLocation.coords[1],
      },
      name: meetupLocation.description,
    };
  }
  if (meetupType === REMOTE) {
    location = {
      url: meetupURL,
    };
  }

  const meetup = {
    description,
    location,
    minCapacity: groupCount[0],
    maxCapacity: groupCount[1],
    tags,
    time,
    title,
  };

  const res = await fetch(createMeetupEndpoint, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(meetup),
  });

  return { res, resJSON: await res.json() };
}
