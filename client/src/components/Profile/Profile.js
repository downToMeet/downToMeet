/* eslint-disable no-nested-ternary */
/* eslint-disable react/jsx-props-no-spreading */
import React, { useState } from "react";
import PropTypes from "prop-types";
import {
  AppBar,
  Box,
  Checkbox,
  FormControlLabel,
  Container,
  Tab,
  Typography,
  Button,
  TextField,
} from "@material-ui/core";
import AccountCircleIcon from "@material-ui/icons/AccountCircle";
import { TabList, TabContext, TabPanel } from "@material-ui/lab";
import { makeStyles, withStyles } from "@material-ui/styles";
import { Link } from "react-router-dom";
import MeetupCard from "../common/MeetupCard";

const useStyles = makeStyles(() => ({
  profilePic: {
    borderRadius: "100%",
    margin: "1.5rem",
    maxWidth: "150px",
  },
  meetupInfoTitle: {
    marginTop: "1rem",
    marginBottom: ".5rem",
  },
}));

const StyledTabList = withStyles({
  root: {
    backgroundColor: "white",
  },
  flexContainer: {
    justifyContent: "space-around",
  },
  indicator: {
    backgroundColor: "black",
  },
})((props) => <TabList {...props} />);

const StyledTab = withStyles({
  root: {
    color: "black",
    textTransform: "none",
  },
})((props) => <Tab disableRipple {...props} />);

const StyledTabPanel = withStyles({
  root: {
    backgroundColor: "white",
    width: "100%",
  },
})((props) => <TabPanel {...props} />);

function shouldDisplayMeetup(meetup, showCanceled, showPast) {
  return (
    (showCanceled || !meetup.canceled) &&
    (showPast || new Date(meetup.time) > new Date())
  );
}

/**
 * Displays the profile of the user with the corresponding `id`. If the profile is the
 * current user's own profile, meetups they own, attend, and have requested to attend
 * are shown as [MeetupCards](#meetupcard). They can also edit their display name and
 * contact information. If the profile is of another user, only the display name,
 * profile picture, and join date are shown for privacy reasons.
 */
function Profile({
  user,
  editing,
  setEditing,
  isMe,
  ownedMeetups,
  attendingMeetups,
  pendingMeetups,
  newName,
  setNewName,
  newContact,
  setNewContact,
  onSubmit,
}) {
  const classes = useStyles();

  if (user === null) {
    return <Typography>Specified user was not found.</Typography>;
  }

  const [tabValue, setTabValue] = useState("owned");
  const [showCanceled, setShowCanceled] = useState(false);
  const [showPast, setShowPast] = useState(false);

  let ownedMeetupsEl = ownedMeetups
    .filter((meetup) => shouldDisplayMeetup(meetup, showCanceled, showPast))
    .map((meetup) => (
      <MeetupCard
        key={meetup.id}
        title={meetup.title}
        time={meetup.time}
        location={meetup.location}
        id={meetup.id}
        owner={meetup.owner}
        tags={meetup.tags}
        canceled={meetup.canceled}
      />
    ));
  if (ownedMeetupsEl.length === 0) {
    ownedMeetupsEl = (
      <Typography>
        You don’t own any meetups. <Link to="/create">Go make one!</Link>
      </Typography>
    );
  }

  let attendingMeetupsEl = attendingMeetups
    .filter((meetup) => shouldDisplayMeetup(meetup, showCanceled, showPast))
    .map((meetup) => (
      <MeetupCard
        key={meetup.id}
        title={meetup.title}
        time={meetup.time}
        location={meetup.location}
        id={meetup.id}
        owner={meetup.owner}
        tags={meetup.tags}
        canceled={meetup.canceled}
      />
    ));
  if (attendingMeetupsEl.length === 0) {
    attendingMeetupsEl = (
      <Typography>
        You are not attending any meetups.{" "}
        <Link to="/">Look for one to attend!</Link>
      </Typography>
    );
  }

  let pendingMeetupsEl = pendingMeetups
    .filter((meetup) => shouldDisplayMeetup(meetup, showCanceled, showPast))
    .map((meetup) => (
      <MeetupCard
        key={meetup.id}
        title={meetup.title}
        time={meetup.time}
        location={meetup.location}
        id={meetup.id}
        owner={meetup.owner}
        tags={meetup.tags}
        canceled={meetup.canceled}
      />
    ));
  if (pendingMeetupsEl.length === 0) {
    pendingMeetupsEl = (
      <Typography>
        You don’t have any pending requests.{" "}
        <Link to="/">Look for a meetup to attend!</Link>
      </Typography>
    );
  }

  const personalInfo = isMe && (
    <>
      {editing ? (
        <Box display="flex" alignItems="center" mt={1}>
          <Typography style={{ marginRight: 15 }}>Contact Info:</Typography>
          <TextField
            value={newContact}
            size="small"
            variant="outlined"
            defaultValue={user.contactInfo}
            onChange={(event) => setNewContact(event.target.value)}
          />
        </Box>
      ) : user.contactInfo ? (
        <Typography>Contact Info: {user.contactInfo}</Typography>
      ) : null}
      {user.email && <Typography>Email: {user.email}</Typography>}
      <Typography
        variant="h4"
        component="h3"
        className={classes.meetupInfoTitle}
      >
        Your meetups:
      </Typography>
      <Box display="flex" flexDirection="row" justifyContent="space-between">
        <FormControlLabel
          control=<Checkbox
            onChange={({ target }) => setShowCanceled(target.checked)}
          />
          label="Show canceled meetups"
        />
        <FormControlLabel
          control=<Checkbox
            onChange={({ target }) => setShowPast(target.checked)}
          />
          label="Show past meetups"
        />
      </Box>
      <TabContext value={tabValue}>
        <AppBar position="static">
          <StyledTabList
            onChange={(_, newValue) => setTabValue(newValue)}
            variant="scrollable"
            scrollButtons="auto"
          >
            <StyledTab label="Owned Meetups" value="owned" />
            <StyledTab label="Attending Meetups" value="attending" />
            <StyledTab label="Pending Meetups" value="pending" />
          </StyledTabList>
        </AppBar>
        <StyledTabPanel value="owned">{ownedMeetupsEl}</StyledTabPanel>
        <StyledTabPanel value="attending">{attendingMeetupsEl}</StyledTabPanel>
        <StyledTabPanel value="pending">{pendingMeetupsEl}</StyledTabPanel>
      </TabContext>
    </>
  );

  return (
    <Container maxWidth="md">
      <Box display="flex" flexDirection="column" alignItems="center">
        {isMe ? (
          editing ? (
            <Button
              onClick={onSubmit}
              variant="outlined"
              style={{ alignSelf: "flex-end", marginTop: "30px" }}
            >
              Save Profile
            </Button>
          ) : (
            <Button
              onClick={() => setEditing(true)}
              variant="outlined"
              style={{ alignSelf: "flex-end", marginTop: "30px" }}
            >
              Edit Profile
            </Button>
          )
        ) : null}
        {user.profilePic && (
          <img
            src={user.profilePic}
            alt="profile"
            className={classes.profilePic}
          />
        )}
        {!user.profilePic && (
          <div
            style={{ position: "relative", width: "150px", height: "150px" }}
          >
            <AccountCircleIcon
              style={{
                position: "absolute",
                left: 0,
                top: 0,
                width: "100%",
                height: "100%",
              }}
            />
          </div>
        )}
        {editing ? (
          <Box display="flex" alignItems="center">
            <Typography style={{ marginRight: 15 }}>Name:</Typography>
            <TextField
              value={newName}
              size="small"
              variant="outlined"
              defaultValue={user.name}
              onChange={(event) => setNewName(event.target.value)}
            />
          </Box>
        ) : (
          <Typography component="h2" variant="h3">
            {user.name}
          </Typography>
        )}
        {personalInfo}
      </Box>
    </Container>
  );
}

const userType = PropTypes.shape({
  id: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  email: PropTypes.string.isRequired,
  connections: PropTypes.arrayOf(PropTypes.string.isRequired),
  contactInfo: PropTypes.string,
  profilePic: PropTypes.string,
});

const meetupType = PropTypes.shape({
  title: PropTypes.string.isRequired,
  time: PropTypes.string.isRequired,
  location: PropTypes.shape({
    coordinates: PropTypes.shape({
      lat: PropTypes.number,
      lon: PropTypes.number,
    }),
    name: PropTypes.string,
    url: PropTypes.string,
  }).isRequired,
  id: PropTypes.string.isRequired,
  owner: PropTypes.string.isRequired,
  tags: PropTypes.arrayOf(PropTypes.string).isRequired,
});

Profile.propTypes = {
  /**
   * userType: `{ id: <string>, 
                  name: <string>, 
                  email: <string>, 
                  connections: <string[]>, 
                  contactInfo: <string>, 
                  profilePic: <string> }`.

   *               
   * `connections` should contain `"Google"` and/or `"Facebook"` depending on what 
   * OAuth(s) were used to log in for the user with that `email`. `profilePic` is
   * a URL to the display picture retrieved from the `connrection` used to sign up.
   */
  user: userType,
  /**
   * Whether the user is editing their display name/contact info.
   */
  editing: PropTypes.bool,
  /** Setter for `editing`, generated from `useEffect()`. */
  setEditing: PropTypes.func,
  /** Whether the displayed profile belongs to the user. */
  isMe: PropTypes.bool.isRequired,
  /** Meetups owned by the user, of the form returned by `GET /meetup/:id`. See backend documentation for more details. */
  ownedMeetups: PropTypes.arrayOf(meetupType).isRequired,
  /** Meetups the user is attending, of the form returned by `GET /meetup/:id`. See backend documentation for more details. */
  attendingMeetups: PropTypes.arrayOf(meetupType).isRequired,
  /** Meetups the user has requested to attend attending, of the form returned by `GET /meetup/:id`. See backend documentation for more details. */
  pendingMeetups: PropTypes.arrayOf(meetupType).isRequired,
  /** Updated display name. */
  newName: PropTypes.string,
  /** Setter for `newName`, generated from `useEffect()`. */
  setNewName: PropTypes.func,
  /** Updated contact information. */
  newContact: PropTypes.string,
  /** Setter for `newContact`, generated from `useEffect()`. */
  setNewContact: PropTypes.func,
  /** Handler for submitting updated display name, contact info. */
  onSubmit: PropTypes.func,
};

Profile.defaultProps = {
  user: null,
  editing: false,
  setEditing: null,
  newName: null,
  setNewName: null,
  newContact: null,
  setNewContact: null,
  onSubmit: null,
};

export default Profile;
