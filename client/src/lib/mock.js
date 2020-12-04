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

export function searchMockRemoteMeetups(meetups) {
  return async (tags) => {
    const goodMeetups = [];
    Object.values(meetups).forEach((meetup) => {
      if (meetup.location.url) {
        // eslint-disable-next-line no-restricted-syntax
        for (const tag in meetup.tags) {
          if (tags.length === 0 || tags.includes(tag)) {
            goodMeetups.push(meetup);
            break;
          }
        }
      }
    });
    const res = new Response(JSON.stringify(goodMeetups), {
      headers: { "Content-Type": "application/json" },
    });
    return { res, resJSON: await res.json() };
  };
}
