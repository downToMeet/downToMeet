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
} from "@material-ui/core";
import GroupAddIcon from "@material-ui/icons/GroupAdd";
import makeStyles from "@material-ui/styles/makeStyles";
import { Link } from "react-router-dom";
import Paper from "@material-ui/core/Paper";

const useStyles = makeStyles(() => ({
  // TODO: mobile scaling
  // TODO: finalize styles (font, color)
  root: {
    display: "flex",
    flexGrow: 1,
    flexDirection: "column",
    alignItems: "center",
  },
  paper: {
    padding: 25,
    width: "75%",
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
  },
  attendeeList: {
    padding: 0,
    listStyle: "none",
  },
  attendee: {
    display: "inline",
  },
  organizer: {
    width: "75%",
  },
}));
function Meetup({ id }) {
  const classes = useStyles();

  const [title, setTitle] = useState("");
  const [time, setTime] = useState(new Date());
  const [meetupType, setMeetupType] = useState("");
  const [meetupURL, setMeetupURL] = useState("");
  const [meetupLocation, setMeetupLocation] = useState(null);
  const [groupCount, setGroupCount] = useState([2, 10]);
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
    attendees: ["jim", "joe", "bob", "foo", "bar"],
    organizer: "kim",
  };

  useEffect(() => {
    setData(mockData);
    // split into organizer, event, attendee data
  }, []);

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
          <Avatar />
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
      <Typography className={classes.link} component={Link} to="/">
        Return home
      </Typography>
      <Paper className={classes.paper}>
        <Typography variant="h1">{title}</Typography>
        <Button startIcon={<GroupAddIcon />} onClick={handleJoinMeetup}>
          Join Meetup
        </Button>
        {/* add share/link copy icon here? */}
        <Typography>Time: {time.toString()}</Typography>
        <Typography>Location: {meetupLocation || "Online"}</Typography>
        <Typography>
          Attendees: {`${attendees.length} out of ${groupCount[1]}`}
        </Typography>
        <Typography component="ul" className={classes.tagList}>
          Tags: {renderTags(tags)}
        </Typography>
        <Card className={classes.organizer}>
          <CardHeader
            avatar={
              <Avatar className={classes.organizerAvatar} src={profilePic} />
            }
            title={organizer}
            subheader={`member since ${time.toString()}`}
          />
          <CardContent>
            <Typography>About me</Typography>
          </CardContent>
          <CardActions disableSpacing>
            <Button>Contact</Button>
          </CardActions>
        </Card>
        <Typography component="ul" className={classes.attendeeList}>
          Attendees: <br />
          {renderAttendees(attendees)}
        </Typography>
      </Paper>
    </Box>
  );
}

Meetup.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Meetup;
