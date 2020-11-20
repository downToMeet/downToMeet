import React, { useState } from "react";
import CssBaseline from "@material-ui/core/CssBaseline";
import Box from "@material-ui/core/Box";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import AddCircleIcon from "@material-ui/icons/AddCircle";
import Avatar from "@material-ui/core/Avatar";
import Menu from "@material-ui/core/Menu";
import MenuItem from "@material-ui/core/MenuItem";
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
  const [profileMenuAnchor, setProfileMenuAnchor] = useState(null);
  // TODO: get profileID and avatar from user
  const profileID = 1234;
  const profileName = "Test User";
  const profilePic =
    "http://web.cs.ucla.edu/~miryung/MiryungKimPhotoAugust2018.jpg";
  const handleProfileMenuClick = (event) => {
    setProfileMenuAnchor(event.currentTarget);
  };

  const handleProfileMenuClose = () => {
    setProfileMenuAnchor(null);
  };

  const ProfileMenu = (
    <>
      <Button
        startIcon={<Avatar src={profilePic} className={classes.avatar} />}
        className={`${classes.button} ${classes.profileButton}`}
        onClick={handleProfileMenuClick}
      >
        {profileName}
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
          to={`${PROFILE_PATH}/${profileID}`}
        >
          Profile
        </MenuItem>
        <MenuItem
          onClick={() => {
            handleProfileMenuClose();
            setAuthenticated(false);
          }}
          component={Link}
          to="/"
        >
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

  const authToggle = (
    <Button
      size="small"
      variant="outlined"
      color="secondary"
      onClick={() => {
        setAuthenticated(!authenticated);
      }}
      className={classes.button}
      component={Link}
      to="/"
    >
      [DEBUG: toggle auth]
    </Button>
  );

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
          {authToggle}
          <Button
            startIcon={<AddCircleIcon />}
            className={`${classes.button} ${classes.createButton}`}
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