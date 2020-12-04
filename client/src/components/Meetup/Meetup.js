import React, { useState, useEffect } from "react";
import PropTypes from "prop-types";
import {
  Avatar,
  Button,
  ButtonGroup,
  Card,
  CardContent,
  CardHeader,
  Chip,
  CircularProgress,
  Container,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  Grid,
  IconButton,
  Link,
  Paper,
  Tooltip,
  Typography,
} from "@material-ui/core";
import {
  Check,
  Clear,
  Edit,
  PersonAdd,
  PersonAddDisabled,
} from "@material-ui/icons";
import { useSelector } from "react-redux";
import { makeStyles } from "@material-ui/core/styles";
import { Link as RouterLink } from "react-router-dom";
import * as fetcher from "../../lib/fetch";
import { OWNER, ATTENDING, PENDING, REJECTED, NONE } from "../../constants";

const useStyles = makeStyles((theme) => ({
  // TODO: finalize styles, possibly refactor into separate file
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
    marginBottom: theme.spacing(1),
  },
  organizer: {
    marginTop: theme.spacing(2),
    width: 500,
  },
  organizerAvatar: {
    width: theme.spacing(8),
    height: theme.spacing(8),
  },
  organizerBio: {
    overflow: "hidden",
    textOverflow: "ellipsis",
    display: "-webkit-box",
    "-webkit-line-clamp": 3,
    "-webkit-box-orient": "vertical",
  },
  avatar: {
    width: theme.spacing(8),
    height: theme.spacing(8),
  },
  profileName: {
    textDecoration: "none",
    color: theme.palette.common.black,
  },
  spinner: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(5),
  },
  ownerInterests: {
    display: "inline-block",
    paddingLeft: 0,
    paddingRight: theme.spacing(1),
    listStyle: "none",
  },
  noInterests: {
    padding: 0,
    color: theme.palette.warning.main,
  },
}));

function Meetup({ id }) {
  const classes = useStyles();

  const [isLoading, setIsLoading] = useState(true);
  const [isUpdating, setIsUpdating] = useState(false);
  const [showLogin, setShowLogin] = useState(false);
  const [fetchMeetupError, setFetchMeetupError] = useState(null);
  const [eventDetails, setEventDetails] = useState(null);
  const [attendees, setAttendees] = useState([]);
  const [pendingAttendees, setPendingAttendees] = useState([]);

  const userID = useSelector((state) => state.id);

  const locale = "en-US";
  const meetupTimeOptions = {
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

  async function setAttendeeLists(attendeeList, pendingAttendeesList) {
    return Promise.all([
      (async () => {
        const res = await Promise.all(
          (attendeeList || []).map(fetcher.getUserData)
        );
        setAttendees(res.map(({ resJSON }) => resJSON));
      })(),
      (async () => {
        const res = await Promise.all(
          (pendingAttendeesList || []).map(fetcher.getUserData)
        );
        setPendingAttendees(res.map(({ resJSON }) => resJSON));
      })(),
    ]);
  }

  let userMeetupStatus = NONE;
  if (eventDetails) {
    if (userID && eventDetails.owner && eventDetails.owner.id === userID) {
      userMeetupStatus = OWNER;
    } else if (userID && pendingAttendees.some((user) => user.id === userID)) {
      userMeetupStatus = PENDING;
    } else if (userID && attendees.some((user) => user.id === userID)) {
      userMeetupStatus = ATTENDING;
    } else if (eventDetails.rejected) {
      userMeetupStatus = REJECTED;
    }
  }

  useEffect(async () => {
    const { res: meetupRes, resJSON: meetupJSON } = await fetcher.getMeetup(id);
    if (!meetupRes.ok) {
      setFetchMeetupError(meetupRes.status);
      setIsLoading(false);
      return;
    }
    const { res: ownerRes, resJSON: ownerJSON } = await fetcher.getUserData(
      meetupJSON.owner
    );
    if (!ownerRes.ok) {
      setIsLoading(false);
      return;
    }

    setEventDetails({
      description: meetupJSON.description,
      location: meetupJSON.location,
      maxCapacity: meetupJSON.maxCapacity,
      minCapacity: meetupJSON.minCapacity,
      owner: ownerJSON,
      tags: meetupJSON.tags,
      time: new Date(meetupJSON.time),
      title: meetupJSON.title,
      rejected: meetupJSON.rejected,
    });
    setAttendeeLists(meetupJSON.attendees, meetupJSON.pendingAttendees);
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
      default:
        errorMessage = "Error: Unspecified error occured.";
    }

    return (
      <Grid container spacing={1}>
        <Typography color="error">{errorMessage}</Typography>
      </Grid>
    );
  };

  const renderTags = (tags) => {
    // Reach goal: convert to search link w/ tags
    return (
      <>
        {tags.map((tagText) => (
          <li key={tagText} className={classes.tag}>
            <Chip
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

  const renderLocation = () => {
    let locationLink;
    if (eventDetails.location.url) {
      const link = (
        <Link href={eventDetails.location.url} rel="noreferrer" target="_blank">
          {eventDetails.location.url}
        </Link>
      );
      // Only show link if owner or attendee (zoombombing who?)
      locationLink = (
        <Typography variant="body2">
          Location: Online{" - "}
          {(userMeetupStatus === ATTENDING || userMeetupStatus === OWNER) &&
            link}
        </Typography>
      );
    } else {
      const googleMapsLink = `https://www.google.com/maps/search/?api=1&query=${eventDetails.location.coordinates.lat},${eventDetails.location.coordinates.lon}`;
      locationLink = (
        <Typography variant="body2">
          Location:{" "}
          <Link href={googleMapsLink} rel="noreferrer" target="_blank">
            {eventDetails.location.name || "Map"}
          </Link>
        </Typography>
      );
    }
    return <Grid item>{locationLink}</Grid>;
  };

  const renderOrganizer = () => {
    return (
      <Grid item>
        <Typography variant="body2">Organizer:</Typography>
        <Card className={classes.organizer}>
          <CardHeader
            avatar={
              <Avatar
                className={classes.organizerAvatar}
                alt={`${eventDetails.owner.name}'s profile pic`}
                src={eventDetails.owner.profilePic}
                component={RouterLink}
                to={`/user/${
                  userID && eventDetails.owner.id === userID
                    ? "me"
                    : eventDetails.owner.id
                }`}
              />
            }
            title={
              <Typography
                className={classes.profileName}
                component={RouterLink}
                to={`/user/${
                  userID && eventDetails.owner.id === userID
                    ? "me"
                    : eventDetails.owner.id
                }`}
              >
                {eventDetails.owner.name}
              </Typography>
            }
            subheader={`member since ${new Date(
              eventDetails.owner.joinDate
            ).toLocaleString(locale, userDateOptions)}`}
          />
          <CardContent>
            <Grid container direction="column" spacing={3}>
              <Grid item>
                <Typography variant="body2" className={classes.ownerInterests}>
                  {eventDetails.owner.interests
                    ? "Interests: "
                    : "Interests: none specified"}
                </Typography>
                {eventDetails.owner.interests && renderTags(eventDetails.tags)}
              </Grid>
              <Grid item>
                <Typography variant="body2" className={classes.ownerInterests}>
                  {`Contact info: ${
                    eventDetails.owner.contactInfo || "none specified"
                  }`}
                </Typography>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
      </Grid>
    );
  };

  const handleUpdateAttendee = async ({ attendee, status }) => {
    setIsUpdating(true);
    await fetcher.updateAttendeeStatus({
      id,
      attendee,
      attendeeStatus: status,
    });
    const { res, resJSON } = await fetcher.getMeetupAttendees(id);
    if (res.ok) {
      setAttendeeLists(resJSON.attending, resJSON.pending);
    }
    setIsUpdating(false);
  };

  const handleJoinMeetup = async () => {
    if (userID) {
      setIsUpdating(true);
      await fetcher.joinMeetup(id);
      const { res, resJSON } = await fetcher.getMeetupAttendees(id);
      if (res.ok) {
        setAttendeeLists(resJSON.attending, resJSON.pending);
      }
      setIsUpdating(false);
    } else {
      setShowLogin(true);
    }
  };

  const renderAttendees = (attendeeList, attendeeType) => {
    let attendeeDisplay;
    const attendeeActions = (attendee) => {
      return (
        <Grid item>
          <ButtonGroup size="small">
            <Tooltip title="Accept attendee">
              <IconButton
                size="small"
                onClick={() =>
                  handleUpdateAttendee({
                    attendee: attendee.id,
                    status: ATTENDING,
                  })
                }
                disabled={isUpdating}
              >
                <Check />
              </IconButton>
            </Tooltip>
            <Tooltip title="Reject attendee">
              <IconButton
                size="small"
                onClick={() =>
                  handleUpdateAttendee({
                    attendee: attendee.id,
                    status: REJECTED,
                  })
                }
                disabled={isUpdating}
              >
                <Clear />
              </IconButton>
            </Tooltip>
          </ButtonGroup>
        </Grid>
      );
    };

    if (attendeeList.length > 0) {
      attendeeDisplay = (
        <Grid item container justify="flex-start">
          {attendeeList.map((attendee) => (
            <Grid
              key={attendee.id}
              item
              container
              direction="column"
              xs={2}
              alignItems="center"
              className={classes.attendee}
              spacing={1}
            >
              <Grid item>
                <Avatar
                  className={classes.avatar}
                  alt={`${attendee.name}'s profile pic`}
                  src={attendee.profilePic}
                  component={RouterLink}
                  to={`/user/${
                    userID && attendee.id === userID ? "me" : attendee.id
                  }`}
                />
              </Grid>
              <Grid item>
                <Typography
                  variant="body2"
                  className={classes.profileName}
                  component={RouterLink}
                  to={`/user/${
                    userID && attendee.id === userID ? "me" : attendee.id
                  }`}
                >
                  {attendee.name}
                </Typography>
              </Grid>
              {attendeeType === PENDING && attendeeActions(attendee)}
            </Grid>
          ))}
        </Grid>
      );
    }

    return (
      <Grid
        item
        container
        direction="column"
        spacing={3}
        className={classes.attendeeList}
      >
        <Grid item>
          <Typography variant="body2">
            {`${
              attendeeType === ATTENDING ? "Attendees: " : "Pending attendees: "
            }`}
            {attendeeList.length === 0 &&
              `there are ${
                attendeeType === ATTENDING ? "currently no " : "no pending"
              } attendees.`}
          </Typography>
        </Grid>
        {attendeeDisplay}
      </Grid>
    );
  };

  const renderMeetupAction = () => {
    let button;
    // TODO: Edit meetup redirect to pre-populated CreateMeetup page
    switch (userMeetupStatus) {
      case OWNER:
        button = (
          <Button
            startIcon={<Edit />}
            variant="outlined"
            component={RouterLink}
          >
            Edit Meetup
          </Button>
        );
        break;
      case ATTENDING:
        button = (
          <Button
            startIcon={<Clear />}
            onClick={() => handleUpdateAttendee({ status: NONE })}
            disabled={isUpdating}
            variant="outlined"
          >
            Leave Meetup
          </Button>
        );
        break;
      case PENDING:
        button = (
          <Button
            startIcon={<PersonAddDisabled />}
            onClick={() => handleUpdateAttendee({ status: NONE })}
            disabled={isUpdating}
            variant="outlined"
          >
            Cancel Join Request
          </Button>
        );
        break;
      case REJECTED:
        button = (
          <Grid container spacing={1} direction="column" align="center">
            <Grid item>
              <Button startIcon={<PersonAdd />} disabled variant="outlined">
                Join Meetup
              </Button>
            </Grid>
            <Grid item>
              <Typography variant="body2" color="error">
                you were rejected from the meetup.
              </Typography>
            </Grid>
          </Grid>
        );
        break;
      default:
        button = (
          <Button
            startIcon={<PersonAdd />}
            onClick={handleJoinMeetup}
            disabled={isUpdating}
          >
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
    if (errorStatus || !eventDetails) {
      return renderError(errorStatus);
    }

    return (
      <Grid item container spacing={2}>
        <Grid item container>
          <Grid item xs>
            <Typography variant="h3">{eventDetails.title}</Typography>
            {/* TODO: add share/link copy icon here? */}
          </Grid>
          {renderMeetupAction()}
        </Grid>
        <Grid
          item
          container
          direction="column"
          spacing={1}
          className={classes.meetupDetails}
        >
          <Grid item>
            <Typography variant="body2">
              Time:{" "}
              {eventDetails.time.toLocaleString(locale, meetupTimeOptions)}
            </Typography>
          </Grid>
          {renderLocation()}
          <Grid item>
            <Typography variant="body2">
              # Attendees:{" "}
              {`${attendees ? attendees.length : 0} out of ${
                eventDetails.maxCapacity
              } (min. ${eventDetails.minCapacity})`}
            </Typography>
          </Grid>
          <Grid item>
            <Typography
              component="ul"
              className={classes.tagList}
              variant="body2"
            >
              Tags: {eventDetails.tags ? renderTags(eventDetails.tags) : "none"}
            </Typography>
          </Grid>
          <Grid item>
            <Typography className={classes.description} variant="body1">
              {eventDetails.description}
            </Typography>
          </Grid>
          {renderOrganizer()}
          <Grid item container justify="center" spacing={1}>
            {renderAttendees(attendees, ATTENDING)}
            {userID &&
              userID === eventDetails.owner.id &&
              renderAttendees(pendingAttendees, PENDING)}
          </Grid>
        </Grid>
      </Grid>
    );
  };

  return (
    <Container className={classes.root} maxWidth="lg">
      <Dialog open={showLogin} onClose={() => setShowLogin(false)}>
        <DialogContent>
          <DialogContentText>Please log in to join a meetup.</DialogContentText>
          <DialogActions>
            <Button onClick={() => setShowLogin(false)}>OK</Button>
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
    </Container>
  );
}

Meetup.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Meetup;
