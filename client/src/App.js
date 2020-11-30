import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import CreateMeetup from "./components/CreateMeetup/CreateMeetup";
import Login from "./components/Login/Login";
import Meetup from "./components/Meetup/Meetup";
import Profile from "./components/Profile/Profile";
import Search from "./components/Search/Search";
import Navbar from "./components/Navbar/Navbar";

function App() {
  // use paper for entire thing

  // use card for meetup organizer info
  // use avatarbuttons for attendees, in a dynamic grid
  return (
    <Router>
      <Navbar />
      <div>
        <Link to="/search">Search</Link>
        <br />
        <Link to="meetup/123">Test meetup</Link>
      </div>
      <Switch>
        <Route path="/create">
          <CreateMeetup />
        </Route>
        <Route path="/login">
          <Login />
        </Route>
        <Route
          path="/meetup/:id"
          render={(input) => <Meetup id={input.match.params.id} />}
        />
        <Route
          path="/user/:id"
          render={(input) => <Profile id={input.match.params.id} />}
        />
        <Route path="/search">
          <Search />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
