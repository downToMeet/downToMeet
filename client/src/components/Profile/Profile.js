/* eslint-disable no-nested-ternary */
/* eslint-disable react/jsx-props-no-spreading */
import React, { useState } from "react";
import PropTypes from "prop-types";
import {
  AppBar,
  Box,
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

  let ownedMeetupsEl = ownedMeetups.map((meetup) => (
    <MeetupCard
      key={meetup.id}
      title={meetup.title}
      time={meetup.time}
      location={meetup.location}
      id={meetup.id}
      owner={meetup.owner}
      tags={meetup.tags}
    />
  ));
  if (ownedMeetups.length === 0) {
    ownedMeetupsEl = (
      <Typography>
        You don’t own any meetups. <Link to="/create">Go make one!</Link>
      </Typography>
    );
  }

  let attendingMeetupsEl = attendingMeetups.map((meetup) => (
    <MeetupCard
      key={meetup.id}
      title={meetup.title}
      time={meetup.time}
      location={meetup.location}
      id={meetup.id}
      owner={meetup.owner}
      tags={meetup.tags}
    />
  ));
  if (attendingMeetups.length === 0) {
    attendingMeetupsEl = (
      <Typography>
        You are not attending any meetups.{" "}
        <Link to="/">Look for one to attend!</Link>
      </Typography>
    );
  }

  let pendingMeetupsEl = pendingMeetups.map((meetup) => (
    <MeetupCard
      key={meetup.id}
      title={meetup.title}
      time={meetup.time}
      location={meetup.location}
      id={meetup.id}
      owner={meetup.owner}
      tags={meetup.tags}
    />
  ));
  if (pendingMeetups.length === 0) {
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
  user: userType,
  editing: PropTypes.bool.isRequired,
  setEditing: PropTypes.func.isRequired,
  isMe: PropTypes.bool.isRequired,
  ownedMeetups: PropTypes.arrayOf(meetupType).isRequired,
  attendingMeetups: PropTypes.arrayOf(meetupType).isRequired,
  pendingMeetups: PropTypes.arrayOf(meetupType).isRequired,
  newName: PropTypes.string.isRequired,
  setNewName: PropTypes.func.isRequired,
  newContact: PropTypes.string.isRequired,
  setNewContact: PropTypes.string.isRequired,
  onSubmit: PropTypes.func.isRequired,
};

Profile.defaultProps = {
  user: null,
};

export default Profile;
