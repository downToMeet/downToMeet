import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import {
  AppBar,
  Avatar,
  Box,
  Button,
  CssBaseline,
  Menu,
  MenuItem,
  Toolbar,
  Typography,
} from "@material-ui/core";
import { AddCircle } from "@material-ui/icons";
import makeStyles from "@material-ui/styles/makeStyles";
import { Link } from "react-router-dom";
import { clearUserData } from "../../stores/user/actions";
import * as fetcher from "../../lib/fetch";

const useStyles = makeStyles(() => ({
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

const PROFILE_PATH = "/user";
const CREATE_PATH = "/create";
const LOGIN_PATH = "/login";

/**
 * Persistent navigation bar displayed throughout the app. At any time throughout the app,
 * the user can do the following:
 * - Clicking the DownToMeet title returns the user to the home (search) page.
 * - Clicking New Meetup will take the user to the [CreateMeetup](#createmeetup) page at `/create`,
 *   or the [Login](#login) page `/login` if they are not authenticated.
 * - If the user is authenticated, clicking the avatar will open a popup menu with a link
 *   to their [Profile](#profile) at `/user/me` and a log out button. If they are not authenticated,
 *   the button takes them to the [Login](#login) page.
 */
function Navbar() {
  const classes = useStyles();
  const [profileMenuAnchor, setProfileMenuAnchor] = useState(null);

  const dispatch = useDispatch();
  const user = useSelector((state) => state);
  const handleProfileMenuClick = (event) => {
    setProfileMenuAnchor(event.currentTarget);
  };

  const handleProfileMenuClose = () => {
    setProfileMenuAnchor(null);
  };

  const handleLogout = async () => {
    const res = await fetcher.logout();
    if (!res.ok) {
      return;
    }
    dispatch(clearUserData());
    handleProfileMenuClose();
  };

  const ProfileMenu = (
    <>
      <Button
        startIcon={<Avatar src={user.profilePic} className={classes.avatar} />}
        className={`${classes.button} ${classes.profileButton}`}
        onClick={handleProfileMenuClick}
      >
        {user.name}
      </Button>
      <Menu
        anchorEl={profileMenuAnchor}
        getContentAnchorEl={null}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "right",
        }}
        transformOrigin={{
          vertical: "top",
          horizontal: "right",
        }}
        keepMounted
        open={Boolean(profileMenuAnchor)}
        onClose={handleProfileMenuClose}
      >
        <MenuItem
          onClick={handleProfileMenuClose}
          component={Link}
          to={`${PROFILE_PATH}/me`}
        >
          Profile
        </MenuItem>
        <MenuItem onClick={handleLogout} component={Link} to="/login">
          Logout
        </MenuItem>
      </Menu>
    </>
  );

  const Login = (
    <Button
      startIcon={<Avatar className={classes.avatar} />}
      className={`${classes.button} ${classes.profileButton}`}
      component={Link}
      to={LOGIN_PATH}
    >
      Login
    </Button>
  );

  const authenticated = Boolean(user.id);

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
            startIcon={<AddCircle />}
            className={`${classes.button} ${classes.createButton}`}
            role="link"
            component={Link}
            // TODO: add redirect to current path
            to={authenticated ? CREATE_PATH : LOGIN_PATH}
          >
            New Meetup
          </Button>
          {authenticated ? ProfileMenu : Login}
        </Toolbar>
      </AppBar>
    </Box>
  );
}

export default Navbar;
