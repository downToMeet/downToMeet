/* eslint-disable no-console */
/* eslint-disable no-unused-vars */
import React, { useState, useEffect } from "react";
import PropTypes from "prop-types";
import {
  Box,
  Typography,
  Chip,
  Card,
  CardHeader,
  CardContent,
  CardActions,
  Avatar,
  Button,
  Grid,
  CircularProgress,
  Dialog,
  DialogContent,
  DialogContentText,
  DialogActions,
} from "@material-ui/core";
import GroupAddIcon from "@material-ui/icons/GroupAdd";
import { makeStyles } from "@material-ui/core/styles";
import { Link as RouterLink } from "react-router-dom";
import Link from "@material-ui/core/Link";
import Paper from "@material-ui/core/Paper";
import * as fetcher from "../../lib/fetch";

const useStyles = makeStyles((theme) => ({
  // TODO: mobile scaling
  // TODO: finalize styles (font, color)
  root: {
    display: "flex",
    flexGrow: 1,
    flexDirection: "column",
    alignItems: "center",
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(4),
    padding: theme.spacing(5),
    width: "80%",
  },
  link: {
    textDecoration: "none",
    color: "grey",
  },
  tagList: {
    padding: 0,
    listStyle: "none",
  },
  tag: {
    display: "inline",
    margin: theme.spacing(0.25),
  },
  description: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  attendeeList: {
    marginTop: theme.spacing(2),
    padding: 0,
    listStyle: "none",
  },
  attendee: {
    marginBottom: theme.spacing(2),
  },
  organizer: {
    marginTop: theme.spacing(2),
    width: 500,
  },
  organizerBio: {
    overflow: "hidden",
    textOverflow: "ellipsis",
    display: "-webkit-box",
    "-webkit-line-clamp": 3,
    "-webkit-box-orient": "vertical",
  },
  spinner: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(5),
  },
  error: {
    color: theme.palette.error.main,
  },
}));
function Meetup({ id }) {
  const classes = useStyles();

  const [isLoading, setIsLoading] = useState(true);
  const [user, setUser] = useState(null);
  const [loginDialog, setLoginDialog] = useState(false);
  const [fetchMeetupError, setFetchMeetupError] = useState(null);
  const [userMeetupStatus, setUserMeetupStatus] = useState("");
  const [title, setTitle] = useState("");
  const [time, setTime] = useState(null);
  const [meetupLocation, setMeetupLocation] = useState(null);
  const [groupCount, setGroupCount] = useState([]);
  const [description, setDescription] = useState("");
  const [tags, setTags] = useState([]);
  const [organizer, setOrganizer] = useState("");
  const [attendees, setAttendees] = useState([]);
  const [pendingAttendees, setPendingAttendees] = useState([]);
  const profilePic =
    "http://web.cs.ucla.edu/~miryung/MiryungKimPhotoAugust2018.jpg";

  const locale = "en-US";
  const eventTimeOptions = {
    hour: "numeric",
    minute: "numeric",
    day: "numeric",
    month: "long",
    year: "numeric",
    timeZoneName: "short",
  };
  const userDateOptions = {
    day: "numeric",
    month: "long",
    year: "numeric",
  };

  const setData = (userData, meetupData) => {
    setTitle(meetupData.title);
    setTime(new Date(meetupData.time));
    setMeetupLocation(meetupData.location);
    setGroupCount([meetupData.minCapacity, meetupData.maxCapacity]);
    setDescription(meetupData.description);
    setTags(meetupData.tags);
    setOrganizer(meetupData.organizer);
    setAttendees(meetupData.attendees);
    setPendingAttendees(meetupData.pendingAttendees);
    if (meetupData.owner === userData.id) {
      setUserMeetupStatus("owner");
    } else if (
      meetupData.pendingAttendees &&
      meetupData.pendingAttendees.includes(userData.id)
    ) {
      setUserMeetupStatus("pending");
    } else if (
      meetupData.attendees &&
      meetupData.attendees.includes(userData.id)
    ) {
      setUserMeetupStatus("attending");
    } else if (meetupData.rejected) {
      setUserMeetupStatus("rejected");
    }
  };

  useEffect(async () => {
    const { res: userRes, resJSON: userJSON } = await fetcher.getUserData();
    if (!userRes.ok) {
      console.log("User not logged in");
    }
    const { res: meetupRes, resJSON: meetupJSON } = await fetcher.getMeetup(id);
    if (meetupRes.ok) {
      setData(userJSON, meetupJSON);
    } else {
      setFetchMeetupError(meetupRes.status);
    }
    setIsLoading(false);
  }, []);

  const Spinner = (
    <Grid
      container
      spacing={3}
      justify="center"
      alignItems="center"
      className={classes.spinner}
    >
      <CircularProgress />
    </Grid>
  );

  const renderError = (status) => {
    let errorMessage = "";

    switch (status) {
      case 404:
        errorMessage = "Error: Meetup not found.";
        break;
      case 400:
        errorMessage = "Error: Bad request.";
        break;
      case null:
        break;
      default:
        errorMessage = "Error: Unspecified error occured.";
    }

    return (
      <Grid container spacing={1}>
        <Typography className={classes.error}>{errorMessage}</Typography>
      </Grid>
    );
  };

  const renderTags = (tagList) => {
    // TODO: convert to search link w/ tags
    return (
      <>
        {tagList.map((tagText) => (
          <li key={tagText} className={classes.tag}>
            <Chip
              clickable
              size="small"
              label={tagText}
              // component={RouterLink}
              // to=""
            />
          </li>
        ))}
      </>
    );
  };

  const renderLocation = (location) => {
    let locationLink;
    if (location.url) {
      locationLink = (
        <Typography variant="body2">
          Location: Online (
          <Link href={location.url} rel="noreferrer" target="_blank">
            link
          </Link>
          )
        </Typography>
      );
    } else {
      const googleMapsLink = `https://www.google.com/maps/search/?api=1&query=${location.coordinates.lat},${location.coordinates.lon}`;
      locationLink = (
        <Typography variant="body2">
          Location:{" "}
          <Link href={googleMapsLink} rel="noreferrer" target="_blank">
            {location.name}
          </Link>
        </Typography>
      );
    }
    return (
      <Grid item>
        {/* TODO: Link to online event if joined? */}
        {locationLink}
      </Grid>
    );
  };

  const renderOrganizer = () => {
    return (
      <Grid item>
        <Typography variant="body2">Organizer:</Typography>
        <Card className={classes.organizer}>
          <CardHeader
            avatar={
              <Avatar className={classes.organizerAvatar} src={profilePic} />
            }
            title={organizer}
            subheader={`member since ${time.toLocaleString(
              locale,
              userDateOptions
            )}`}
          />
          <CardContent>
            {/* TODO: add bio (reach goal) or convert to list of user interests (tags) */}
            <Typography className={classes.organizerBio}>
              About Me: Lorem ipsum dolor sit amet, consectetur adipiscing elit,
              sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
              Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
              nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in
              reprehenderit in voluptate velit esse cillum dolore eu fugiat
              nulla pariatur. Excepteur sint occaecat cupidatat non proident,
              sunt in culpa qui officia deserunt mollit anim id est laborum.
            </Typography>
          </CardContent>
          <CardActions disableSpacing>
            {/* TODO: link to user profile, or create popup with contact info */}
            <Button>Contact</Button>
          </CardActions>
        </Card>
      </Grid>
    );
  };

  const renderAttendees = (attendeeList) => {
    // https://github.com/mui-org/material-ui/blob/master/packages/material-ui-lab/src/AvatarGroup/AvatarGroup.js
    // TODO: convert to modal expandable list when too many attendees

    let attendeeDisplay;
    if (attendeeList) {
      attendeeDisplay = (
        <Grid item container justify="center">
          {attendeeList.map((attendee) => (
            <Grid
              item
              container
              direction="column"
              xs={3}
              alignItems="center"
              className={classes.attendee}
            >
              <Avatar size="small" />
              {attendee}
            </Grid>
          ))}
        </Grid>
      );
    } else {
      attendeeDisplay = (
        <Grid item>
          <Typography>There are currently no attendees.</Typography>
        </Grid>
      );
    }

    return (
      <Grid
        item
        container
        direction="column"
        spacing={1}
        className={classes.attendeeList}
      >
        <Grid item>
          <Typography variant="body2">Attendees: </Typography>
        </Grid>
        {attendeeDisplay}
      </Grid>
    );
  };

  const handleJoinMeetup = async () => {
    // TODO: get user ID and submit join meetup to backend
    // eslint-disable-next-line no-console
    if (user) {
      fetcher.joinMeetup(id);
      console.log("Joining meetup");
    } else {
      setLoginDialog(true);
    }
  };

  const handleLeaveMeetup = async () => {
    console.log("Canceling join request");
    // patch with status = "none"
  };

  const renderMeetupAction = () => {
    let button;
    // TODO: Edit meetup redirect to pre-populated CreateMeetup page
    switch (userMeetupStatus) {
      case "owner":
        button = (
          <Button startIcon={<GroupAddIcon />} component={RouterLink}>
            Edit Meetup
          </Button>
        );
        break;
      case "attending":
        button = (
          <Button startIcon={<GroupAddIcon />} onClick={handleLeaveMeetup}>
            Leave Meetup
          </Button>
        );
        break;
      case "pending":
        button = (
          <Button startIcon={<GroupAddIcon />} onClick={handleLeaveMeetup}>
            Cancel Join Request
          </Button>
        );
        break;
      default:
        button = (
          <Button startIcon={<GroupAddIcon />} onClick={handleJoinMeetup}>
            Join Meetup
          </Button>
        );
    }
    return (
      <Grid item xs={2} container justify="center" alignItems="flex-start">
        {button}
      </Grid>
    );
  };

  const renderMeetup = (errorStatus) => {
    if (errorStatus) {
      return renderError(errorStatus);
    }

    return (
      <Grid item container spacing={2}>
        <Grid item container>
          <Grid item xs>
            <Typography variant="h3">{title}</Typography>
            {/* add share/link copy icon here? */}
          </Grid>
          {renderMeetupAction()}
        </Grid>
        <Grid
          item
          container
          direction="column"
          spacing={1}
          className={classes.eventDetails}
        >
          <Grid item>
            <Typography variant="body2">
              Time: {time.toLocaleString(locale, eventTimeOptions)}
            </Typography>
          </Grid>
          {renderLocation(meetupLocation)}
          <Grid item>
            <Typography variant="body2">
              # Attendees:{" "}
              {`${attendees ? attendees.length : 0} out of ${
                groupCount[1]
              } (min. ${groupCount[0]})`}
            </Typography>
          </Grid>
          <Grid item>
            <Typography
              component="ul"
              className={classes.tagList}
              variant="body2"
            >
              Tags: {renderTags(tags)}
            </Typography>
          </Grid>
          <Grid item>
            <Typography className={classes.description} variant="body1">
              {description}
            </Typography>
          </Grid>
          {renderOrganizer()}
          {renderAttendees(attendees)}
        </Grid>
      </Grid>
    );
  };

  return (
    <Box className={classes.root}>
      <Dialog open={loginDialog} onClose={() => setLoginDialog(false)}>
        <DialogContent>
          <DialogContentText>Please log in to join a meetup.</DialogContentText>
          <DialogActions>
            <Button onClick={() => setLoginDialog(false)}>OK</Button>
            <Button component={RouterLink} to="/login" color="primary">
              Sign in
            </Button>
          </DialogActions>
        </DialogContent>
      </Dialog>
      <Paper className={classes.paper}>
        <Grid container spacing={1}>
          <Grid item>
            <Typography
              className={classes.link}
              component={RouterLink}
              to="/"
              variant="subtitle1"
            >
              &lt; return home
            </Typography>
          </Grid>
          {isLoading ? Spinner : renderMeetup(fetchMeetupError)}
        </Grid>
      </Paper>
    </Box>
  );
}

Meetup.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Meetup;
