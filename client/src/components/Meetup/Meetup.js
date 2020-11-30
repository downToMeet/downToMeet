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
} from "@material-ui/core";
import GroupAddIcon from "@material-ui/icons/GroupAdd";
import { makeStyles } from "@material-ui/core/styles";
import { Link } from "react-router-dom";
import Paper from "@material-ui/core/Paper";

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
}));
function Meetup({ id }) {
  const classes = useStyles();

  const [title, setTitle] = useState("");
  const [time, setTime] = useState(new Date());
  const [meetupType, setMeetupType] = useState("");
  const [meetupURL, setMeetupURL] = useState("");
  const [meetupLocation, setMeetupLocation] = useState(null);
  const [groupCount, setGroupCount] = useState([2, 18]);
  const [description, setDescription] = useState("");
  const [tags, setTags] = useState([]);
  const [attendees, setAttendees] = useState([]);
  const [organizer, setOrganizer] = useState("");
  const profilePic =
    "http://web.cs.ucla.edu/~miryung/MiryungKimPhotoAugust2018.jpg";

  const setData = (data) => {
    setTitle(data.title);
    setTime(data.time);
    setMeetupType(data.type);
    setMeetupURL(data.url);
    setMeetupLocation(data.location);
    setGroupCount(data.groupCount);
    setDescription(data.description);
    setTags(data.tags);
    setAttendees(data.attendees);
    setOrganizer(data.organizer);
  };

  const mockData = {
    title: "Test meetup",
    time: new Date(),
    type: "Online",
    url: "/123",
    location: null,
    groupCount: [2, 10],
    description: "Test meetup, do not join",
    tags: ["send", "help", "pls"],
    attendees: [
      "jim",
      "joe",
      "bob",
      "foo",
      "bar",
      "baz",
      "idk",
      "any",
      "more",
      "names",
    ],
    organizer: "kim",
  };

  useEffect(() => {
    setData(mockData);
    // split into organizer, event, attendee data
  }, []);

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
              // component={Link}
              // to=""
            />
          </li>
        ))}
      </>
    );
  };

  const renderAttendees = (attendeeList) => {
    // https://github.com/mui-org/material-ui/blob/master/packages/material-ui-lab/src/AvatarGroup/AvatarGroup.js
    // TODO: convert to modal expandable list when too many attendees
    return (
      <>
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
      </>
    );
  };

  const handleJoinMeetup = () => {
    // TODO: get user ID and submit join meetup to backend
    // eslint-disable-next-line no-console
    console.log("Joining meetup");
  };

  return (
    <Box className={classes.root}>
      <Paper className={classes.paper}>
        <Grid container spacing={1}>
          <Grid item>
            <Typography className={classes.link} component={Link} to="/">
              &lt; return home
            </Typography>
          </Grid>
          <Grid item container>
            <Grid item xs>
              <Typography variant="h3">{title}</Typography>
              {/* add share/link copy icon here? */}
            </Grid>
            <Grid
              item
              xs={2}
              container
              justify="center"
              alignItems="flex-start"
            >
              <Button startIcon={<GroupAddIcon />} onClick={handleJoinMeetup}>
                Join Meetup
              </Button>
            </Grid>
          </Grid>
          <Grid item container direction="column" spacing={1}>
            <Grid item>
              <Typography>
                Time: {time.toLocaleString(locale, eventTimeOptions)}
              </Typography>
            </Grid>
            <Grid item>
              {/* TODO: Link to online event if joined? */}
              <Typography>Location: {meetupLocation || "Online"}</Typography>
            </Grid>
            <Grid item>
              <Typography>
                Attendees:{" "}
                {`${attendees.length} out of ${groupCount[1]} (min. ${groupCount[0]})`}
              </Typography>
            </Grid>
            <Grid item>
              <Typography component="ul" className={classes.tagList}>
                Tags: {renderTags(tags)}
              </Typography>
            </Grid>
            <Grid item>
              <Typography>Organizer:</Typography>
              <Card className={classes.organizer}>
                <CardHeader
                  avatar={
                    <Avatar
                      className={classes.organizerAvatar}
                      src={profilePic}
                    />
                  }
                  title={organizer}
                  subheader={`member since ${time.toLocaleString(
                    locale,
                    userDateOptions
                  )}`}
                />
                <CardContent>
                  <Typography className={classes.organizerBio}>
                    About Me: Lorem ipsum dolor sit amet, consectetur adipiscing
                    elit, sed do eiusmod tempor incididunt ut labore et dolore
                    magna aliqua. Ut enim ad minim veniam, quis nostrud
                    exercitation ullamco laboris nisi ut aliquip ex ea commodo
                    consequat. Duis aute irure dolor in reprehenderit in
                    voluptate velit esse cillum dolore eu fugiat nulla pariatur.
                    Excepteur sint occaecat cupidatat non proident, sunt in
                    culpa qui officia deserunt mollit anim id est laborum.
                  </Typography>
                </CardContent>
                <CardActions disableSpacing>
                  <Button>Contact</Button>
                </CardActions>
              </Card>
            </Grid>
            <Grid item container direction="column" spacing={1}>
              <Grid item>
                <Typography>Attendees: </Typography>
              </Grid>
              <Grid
                item
                container
                justify="center"
                className={classes.attendeeList}
              >
                {renderAttendees(attendees)}
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Paper>
    </Box>
  );
}

Meetup.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Meetup;
