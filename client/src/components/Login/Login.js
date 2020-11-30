import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import { Box, Button, Container, Typography } from "@material-ui/core";
import FacebookLogo from "./assets/FacebookLogo";
import GoogleLogo from "./assets/GoogleLogo";
import { SERVER_URL } from "../../constants";

const useStyles = makeStyles((theme) => ({
  loginButton: {
    margin: "1em",
    backgroundColor: "#EEEEEE",
    textTransform: "capitalize",
    fontFamily: theme.fontFamily,
    "&:hover": {
      backgroundColor: "#CCCCCC",
    },
  },
  loginBox: {
    height: "calc(90vh - 40px)",
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
  },
}));

function Login() {
  const classes = useStyles();
  return (
    <Container maxWidth="sm" className={classes.loginBox}>
      <Typography
        component="h1"
        variant="h2"
        align="center"
        style={{ marginBottom: "clamp(10px, 7vh, 44px)" }}
      >
        Sign in
      </Typography>
      <Box display="flex" justifyContent="center">
        <Button
          component="a"
          href={`${SERVER_URL}/user/facebook/auth`}
          className={classes.loginButton}
          startIcon={<FacebookLogo />}
        >
          Using Facebook
        </Button>
        <Button
          component="a"
          href={`${SERVER_URL}/user/google/auth`}
          className={classes.loginButton}
          startIcon={<GoogleLogo />}
        >
          Using Google
        </Button>
      </Box>
    </Container>
  );
}

export default Login;
