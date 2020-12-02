/* eslint-disable no-console */
/* eslint-disable no-unused-vars */
import React, { useState, useEffect } from "react";
import PropTypes from "prop-types";
import {
  Avatar,
  Box,
  Button,
  ButtonGroup,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Chip,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  Grid,
  IconButton,
  Tooltip,
  Typography,
} from "@material-ui/core";
import Link from "@material-ui/core/Link";
import Paper from "@material-ui/core/Paper";
import CheckIcon from "@material-ui/icons/Check";
import ClearIcon from "@material-ui/icons/Clear";
import EditIcon from "@material-ui/icons/Edit";
import PersonAddDisabledIcon from "@material-ui/icons/PersonAddDisabled";
import PersonAddIcon from "@material-ui/icons/PersonAdd";
import { makeStyles } from "@material-ui/core/styles";
import { Link as RouterLink } from "react-router-dom";
import * as fetcher from "../../lib/fetch";

const useStyles = makeStyles((theme) => ({
  // TODO: reorganize styles, possibly refactor into separate file
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
  error: {
    color: theme.palette.error.main,
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
  const [user, setUser] = useState(null);
  const [showLogin, setShowLogin] = useState(false);
  const [fetchMeetupError, setFetchMeetupError] = useState(null);
  const [userMeetupStatus, setUserMeetupStatus] = useState("");
  const [eventDetails, setEventDetails] = useState(null);
  const [attendees, setAttendees] = useState([]);
  const [pendingAttendees, setPendingAttendees] = useState([]);
  const profilePic =
    "http://web.cs.ucla.edu/~miryung/MiryungKimPhotoAugust2018.jpg";

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
    setAttendees(await fetcher.getDataForUsers(attendeeList));
    setPendingAttendees(await fetcher.getDataForUsers(pendingAttendeesList));
  }

  const setUserStatus = (userData, meetupData) => {
    if (userData && meetupData.owner === userData.id) {
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
      setIsLoading(false);
      return;
    }
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

    setUser(userJSON);
    setEventDetails({
      description: meetupJSON.description,
      location: meetupJSON.location,
      maxCapacity: meetupJSON.maxCapacity,
      minCapacity: meetupJSON.minCapacity,
      owner: ownerJSON,
      tags: meetupJSON.tags,
      time: new Date(meetupJSON.time),
      title: meetupJSON.title,
    });
    setAttendeeLists(meetupJSON.attendees, meetupJSON.pendingAttendees);
    setUserStatus(userJSON, meetupJSON);
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
        <Typography className={classes.error}>{errorMessage}</Typography>
      </Grid>
    );
  };

  const renderTags = (tags) => {
    // TODO: convert to search link w/ tags
    return (
      <>
        {tags.map((tagText) => (
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

  const renderLocation = () => {
    let locationLink;
    if (eventDetails.location.url) {
      const link = (
        <>
          (
          <Link
            href={eventDetails.location.url}
            rel="noreferrer"
            target="_blank"
          >
            link
          </Link>
          )
        </>
      );
      // Only show link if owner or attendee (zoombombing who?)
      locationLink = (
        <Typography variant="body2">
          Location: Online{" "}
          {(userMeetupStatus === "attending" || userMeetupStatus === "owner") &&
            link}
        </Typography>
      );
    } else {
      const googleMapsLink = `https://www.google.com/maps/search/?api=1&query=${eventDetails.coordinates.lat},${eventDetails.coordinates.lon}`;
      locationLink = (
        <Typography variant="body2">
          Location:{" "}
          <Link href={googleMapsLink} rel="noreferrer" target="_blank">
            {eventDetails.location.name}
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
                src={eventDetails.owner.profilePic}
                component={RouterLink}
                to={`/user/${
                  eventDetails.owner.id === user.id
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
                  eventDetails.owner.id === user.id
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
    // If status is "none", remove attendee
    // If attendee is null, update self
    if (!attendee && status === "none") {
      setUserMeetupStatus("");
    }
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
    // TODO: get user ID and submit join meetup to backend
    // eslint-disable-next-line no-console
    if (user) {
      setIsUpdating(true);
      await fetcher.joinMeetup(id);
      const { res, resJSON } = await fetcher.getMeetupAttendees(id);
      if (res.ok) {
        setAttendeeLists(resJSON.attending, resJSON.pending);
      }
      setIsUpdating(false);
      setUserMeetupStatus("pending");
    } else {
      setShowLogin(true);
    }
  };

  const renderAttendees = () => {
    // https://github.com/mui-org/material-ui/blob/master/packages/material-ui-lab/src/AvatarGroup/AvatarGroup.js
    // TODO: convert to modal expandable list when too many attendees

    let attendeeDisplay;
    if (attendees) {
      attendeeDisplay = (
        <>
          {attendees.map((attendee) => (
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
                  src={attendee.profilePic}
                  component={RouterLink}
                  to={`/user/${attendee.id === user.id ? "me" : attendee.id}`}
                />
              </Grid>
              <Grid item>
                <Typography
                  variant="body2"
                  className={classes.profileName}
                  component={RouterLink}
                  to={`/user/${attendee.id === user.id ? "me" : attendee.id}`}
                >
                  {attendee.name}
                </Typography>
              </Grid>
            </Grid>
          ))}
        </>
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
        spacing={3}
        className={classes.attendeeList}
      >
        <Grid item>
          <Typography variant="body2">Attendees: </Typography>
        </Grid>
        <Grid item>{attendeeDisplay}</Grid>
      </Grid>
    );
  };

  const renderPendingAttendees = () => {
    // https://github.com/mui-org/material-ui/blob/master/packages/material-ui-lab/src/AvatarGroup/AvatarGroup.js
    // TODO: convert to modal expandable list when too many attendees

    let pendingDisplay;
    if (pendingAttendees) {
      pendingDisplay = (
        <Grid item container justify="flex-start">
          {pendingAttendees.map((attendee) => (
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
                <Avatar className={classes.avatar} src={attendee.profilePic} />
              </Grid>
              <Grid item>
                <Typography variant="body2">{attendee.name}</Typography>
              </Grid>
              <Grid item>
                <ButtonGroup size="small">
                  <Tooltip title="Accept attendee">
                    <IconButton
                      size="small"
                      onClick={() =>
                        handleUpdateAttendee({
                          attendee: attendee.id,
                          status: "attending",
                        })
                      }
                      disabled={isUpdating}
                    >
                      <CheckIcon />
                    </IconButton>
                  </Tooltip>
                  <Tooltip title="Reject attendee">
                    <IconButton
                      size="small"
                      onClick={() =>
                        handleUpdateAttendee({
                          attendee: attendee.id,
                          status: "rejected",
                        })
                      }
                      disabled={isUpdating}
                    >
                      <ClearIcon />
                    </IconButton>
                  </Tooltip>
                </ButtonGroup>
              </Grid>
            </Grid>
          ))}
        </Grid>
      );
    } else {
      pendingDisplay = (
        <Grid item>
          <Typography>No pending attendees.</Typography>
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
          <Typography variant="body2">Pending Attendees: </Typography>
        </Grid>
        <Grid item>{pendingDisplay}</Grid>
      </Grid>
    );
  };

  const renderMeetupAction = () => {
    let button;
    // TODO: Edit meetup redirect to pre-populated CreateMeetup page
    switch (userMeetupStatus) {
      case "owner":
        button = (
          <Button startIcon={<EditIcon />} component={RouterLink}>
            Edit Meetup
          </Button>
        );
        break;
      case "attending":
        button = (
          <Button
            startIcon={<ClearIcon />}
            onClick={() => handleUpdateAttendee({ status: "none" })}
            disabled={isUpdating}
          >
            Leave Meetup
          </Button>
        );
        break;
      case "pending":
        button = (
          <Button
            startIcon={<PersonAddDisabledIcon />}
            onClick={() => handleUpdateAttendee({ status: "none" })}
            disabled={isUpdating}
          >
            Cancel Join Request
          </Button>
        );
        break;
      case "rejected":
        button = (
          <Grid container spacing={1} direction="column" align="center">
            <Grid item>
              <Button startIcon={<PersonAddIcon />} disabled variant="outlined">
                Join Meetup
              </Button>
            </Grid>
            <Grid item>
              <Typography variant="body2" className={classes.error}>
                you were rejected from the meetup.
              </Typography>
            </Grid>
          </Grid>
        );
        break;
      default:
        button = (
          <Button
            startIcon={<PersonAddIcon />}
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
            {/* add share/link copy icon here? */}
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
            {renderAttendees()}
            {user &&
              user.id === eventDetails.owner.id &&
              renderPendingAttendees(pendingAttendees)}
          </Grid>
        </Grid>
      </Grid>
    );
  };

  return (
    <Box className={classes.root}>
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
    </Box>
  );
}

Meetup.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Meetup;
