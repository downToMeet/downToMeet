import React, { useEffect } from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import CreateMeetup from "./components/CreateMeetup/CreateMeetup";
import Login from "./components/Login/Login";
import Meetup from "./components/Meetup/Meetup";
import Profile from "./components/Profile/Profile";
import Search from "./components/Search/Search";
import Navbar from "./components/Navbar/Navbar";
import { getUserData } from "./lib/fetch";
import { clearUserData, updateUserData } from "./stores/user/actions";

function App() {
  const dispatch = useDispatch();
  const userID = useSelector((state) => state.id);

  useEffect(() => {
    (async () => {
      const { res, resJSON } = await getUserData();
      if (!res.ok) {
        dispatch(clearUserData());
        return;
      }
      dispatch(updateUserData(resJSON));
    })();
  }, [userID]);

  // TODO: use Paper/Cards for interface
  return (
    <Router>
      <Navbar />
      <Switch>
        <Route path="/create">
          <CreateMeetup />
        </Route>
        <Route path="/login">
          <Login />
        </Route>
        <Route
          path="/meetup/:id/edit"
          render={(input) => <CreateMeetup id={input.match.params.id} />}
        />
        <Route
          path="/meetup/:id"
          render={(input) => <Meetup id={input.match.params.id} />}
        />
        <Route
          path="/user/:id"
          render={(input) => <Profile id={input.match.params.id} />}
        />
        <Route path="/">
          <Search />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
