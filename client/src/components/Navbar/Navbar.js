import React, { useState } from "react";
import CssBaseline from "@material-ui/core/CssBaseline";
import Box from "@material-ui/core/Box";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
// import Grid from '@material-ui/core/Grid';
import Button from "@material-ui/core/Button";
// import IconButton from '@material-ui/core/IconButton';
import AddCircleIcon from "@material-ui/icons/AddCircle";
import Avatar from "@material-ui/core/Avatar";
import { makeStyles } from "@material-ui/core/styles";
import { Link } from "react-router-dom";

const useStyles = makeStyles(() => ({
  // TODO: mobile scaling
  // TODO: finalize styles (font, color)
  // TODO: add logo?
  root: {
    flexGrow: 1,
  },
  toolbar: {
    paddingLeft: 20,
    paddingRight: 15,
  },
  button: {
    minWidth: 100,
    justifyContent: "left",
    whiteSpace: "nowrap",
    overflow: "hidden",
    color: "white",
    textTransform: "none",
  },
  createButton: {
    margin: 5,
  },
  profileButton: {
    maxWidth: "12%", // to catch long profile name
  },
  avatar: {
    width: "1.5em",
    height: "1.5em",
  },
  title: {
    flexGrow: 1,
    color: "white",
    textDecoration: "none",
  },
}));

const PROFILE_PATH = "/profile";
const CREATE_PATH = "/create";
const LOGIN_PATH = "/login";

function Navbar() {
  const classes = useStyles();
  // TODO: connect authentication
  const [authenticated, setAuthenticated] = useState(true);
  // TODO: get profileID and avatar from user
  const profileID = 1234;
  const profilePic =
    "http://web.cs.ucla.edu/~miryung/MiryungKimPhotoAugust2018.jpg";

  return (
    <Box className={classes.root}>
      {/* CssBaseline clears default HTML styling (margins, etc.) */}
      <CssBaseline />
      <AppBar position="sticky">
        <Toolbar className={classes.toolbar}>
          <Typography
            variant="h5"
            className={classes.title}
            component={Link}
            to="/"
          >
            DownToMeet
          </Typography>
          <Button
            size="small"
            variant="outlined"
            color="secondary"
            onClick={() => {
              setAuthenticated(!authenticated);
            }}
            className={classes.button}
          >
            [DEBUG: toggle auth]
          </Button>
          <Button
            startIcon={<AddCircleIcon />}
            className={`${classes.button} ${classes.createButton}`}
            component={Link}
            to={CREATE_PATH}
          >
            New Meetup
          </Button>
          <Button
            startIcon={
              <Avatar
                src={authenticated ? profilePic : null}
                className={classes.avatar}
              />
            }
            className={`${classes.button} ${classes.profileButton}`}
            component={Link}
            to={authenticated ? `${PROFILE_PATH}/${profileID}` : LOGIN_PATH}
            width={100}
          >
            {authenticated ? "Test User" : "Login"}
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default Navbar;
