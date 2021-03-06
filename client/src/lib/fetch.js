import { IN_PERSON, REMOTE, SERVER_URL } from "../constants";

export async function getUserData(id = "me") {
  const getUserDataEndpoint = `${SERVER_URL}/user/${id}`;
  const res = await fetch(getUserDataEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export async function patchUserData(id, oldUser, newName, newContact) {
  const patchUserDataEndpoint = `${SERVER_URL}/user/${id}`;

  const newUser = { ...oldUser };
  newUser.name = newName;
  newUser.contactInfo = newContact;
  const res = await fetch(patchUserDataEndpoint, {
    method: "PATCH",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(newUser),
  });

  return { res, resJSON: await res.json() };
}

export async function searchForMeetups({ lat, lon, radius, tags }) {
  let getMeetupsEndpoint = `${SERVER_URL}/meetup?lat=${lat}&lon=${lon}&radius=${radius}`;
  if (tags.length !== 0) {
    getMeetupsEndpoint += `&tags=${tags}`;
  }
  const res = await fetch(getMeetupsEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export async function searchForRemoteMeetups(tags) {
  let getMeetupsEndpoint = `${SERVER_URL}/meetup/remote`;
  if (tags.length !== 0) {
    getMeetupsEndpoint += `?tags=${tags}`;
  }
  const res = await fetch(getMeetupsEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export async function cancelMeetup(id = "me") {
  const cancelMeetupEndpoint = `${SERVER_URL}/meetup/${id}`;
  const res = await fetch(cancelMeetupEndpoint, {
    method: "DELETE",
    credentials: "include",
  });

  return res;
}

export async function createOrEditMeetup(
  {
    id,
    title,
    time,
    meetupType,
    meetupURL,
    meetupLocation,
    groupCount,
    description,
    tags,
  },
  isEdit = false
) {
  const createMeetupEndpoint = `${SERVER_URL}/meetup`;
  const editMeetupEndpoint = `${SERVER_URL}/meetup/${id}`;
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

  const res = await fetch(isEdit ? editMeetupEndpoint : createMeetupEndpoint, {
    method: isEdit ? "PATCH" : "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(meetup),
  });

  return { res, resJSON: await res.json() };
}

export async function logout() {
  const getLogoutEndpoint = `${SERVER_URL}/user/logout`;
  const res = await fetch(getLogoutEndpoint, {
    credentials: "include",
  });
  return res;
}

export async function getMeetup(id) {
  const getMeetupEndpoint = `${SERVER_URL}/meetup/${id}`;

  const res = await fetch(getMeetupEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export async function joinMeetup(id) {
  const addAttendeeEndpoint = `${SERVER_URL}/meetup/${id}/attendee`;

  const res = await fetch(addAttendeeEndpoint, {
    method: "POST",
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}
export async function getMeetupAttendees(id) {
  const getMeetupEndpoint = `${SERVER_URL}/meetup/${id}/attendee`;

  const res = await fetch(getMeetupEndpoint, {
    credentials: "include",
  });

  return { res, resJSON: await res.json() };
}

export async function updateAttendeeStatus({ id, attendee, attendeeStatus }) {
  const updateAttendeeEndpoint = `${SERVER_URL}/meetup/${id}/attendee`;
  const attendeePatch = {
    attendee, // If empty, patches current user
    attendeeStatus,
  };

  const res = await fetch(updateAttendeeEndpoint, {
    method: "PATCH",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(attendeePatch),
  });

  return { res, resJSON: await res.json() };
}
