/* eslint-env jest */
export function getMockUserData(users) {
  return async (id) => {
    let res;
    const user = users.get(id);
    if (user) {
      res = new Response(JSON.stringify(user), {
        headers: { "Content-Type": "application/json" },
      });
    } else {
      res = new Response(
        JSON.stringify({
          code: 404,
          message: `no user with id ${id} found`,
        }),
        {
          status: 404,
          headers: { "Content-Type": "application/json" },
        }
      );
    }
    return { res, resJSON: await res.json() };
  };
}

export function getMockMeetup(meetups) {
  return async (id) => {
    let res;
    const meetup = meetups.get(id);
    if (meetup) {
      res = new Response(JSON.stringify(meetup), {
        headers: { "Content-Type": "application/json" },
      });
    } else {
      res = new Response(
        JSON.stringify({
          code: 404,
          message: `no meetup with id ${id} found`,
        }),
        {
          status: 404,
          headers: { "Content-Type": "application/json" },
        }
      );
    }
    return { res, resJSON: await res.json() };
  };
}

export function getMockMeetupAttendees(meetups) {
  return async (id) => {
    let res;
    const meetup = meetups.get(id);
    if (meetup) {
      res = new Response(
        JSON.stringify({
          attending: meetup.attendees,
          pending: meetup.pendingAttendees,
        }),
        {
          headers: { "Content-Type": "application/json" },
        }
      );
    } else {
      res = new Response(
        JSON.stringify({
          code: 404,
          message: `no meetup with id ${id} found`,
        }),
        {
          status: 404,
          headers: { "Content-Type": "application/json" },
        }
      );
    }
    return { res, resJSON: await res.json() };
  };
}
